package resolvers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/webomindapps-dev/coolaid-backend/config"
	"github.com/webomindapps-dev/coolaid-backend/db"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
	"github.com/webomindapps-dev/coolaid-backend/utils"

	sql_models "github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
)

type AuthClaims struct {
	UserID    string    `json:"sub"`
	SessionID string    `json:"sid"`
	ExpiresAt time.Time `json:"exp"`
}

func AuthContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqCtx := c.Request.Context()

		// ------------------------------------------------
		// 1) CHECK ACCESS TOKEN COOKIE
		// ------------------------------------------------
		accessCookie, err := c.Request.Cookie("access_token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				oplog.Info(reqCtx, "unauthenticated request: no access_token cookie")
			} else {
				oplog.Error(reqCtx, "failed to read access_token cookie:", err)
			}

			logRefreshCookie(reqCtx, c, "unauthenticated")
			c.Request = c.Request.WithContext(reqCtx)
			c.Next()
			return
		}

		oplog.Info(reqCtx, "access_token cookie present")
		logRefreshCookie(reqCtx, c, "access_ok")

		// ------------------------------------------------
		// 2) VALIDATE ACCESS TOKEN
		// ------------------------------------------------
		claims, err := utils.ValidateJWT(accessCookie.Value)
		if err != nil {
			oplog.Warn(reqCtx, "invalid or expired access token:", err)
			logRefreshCookie(reqCtx, c, "access_invalid")

			c.Request = c.Request.WithContext(reqCtx)
			c.Next()
			return
		}

		// ------------------------------------------------
		// 3) VALIDATE UUIDs FROM TOKEN
		// ------------------------------------------------
		userID, err := uuid.Parse(claims.UserID)
		if err != nil {
			oplog.Error(reqCtx, "invalid user UUID in JWT:", claims.UserID, err)
			c.Request = c.Request.WithContext(reqCtx)
			c.Next()
			return
		}

		sessionID, err := uuid.Parse(claims.SessionID)
		if err != nil {
			oplog.Error(reqCtx, "invalid session UUID in JWT:", claims.SessionID, err)
			c.Request = c.Request.WithContext(reqCtx)
			c.Next()
			return
		}

		// ------------------------------------------------
		// 4) LOAD SESSION FROM DB
		// ------------------------------------------------
		q := db.DB.Queries
		session, err := q.GetUserSessionById(reqCtx, sessionID)
		if err != nil {
			oplog.Error(reqCtx, "failed to load session:", sessionID, err)
			c.Request = c.Request.WithContext(reqCtx)
			c.Next()
			return
		}

		if session.RevokedAt.Valid {
			oplog.Warn(reqCtx, "session revoked:", sessionID)
			c.Request = c.Request.WithContext(reqCtx)
			c.Next()
			return
		}

		if session.ExpiresAt.Before(time.Now()) {
			oplog.Warn(reqCtx, "session expired:", sessionID, session.ExpiresAt)
			c.Request = c.Request.WithContext(reqCtx)
			c.Next()
			return
		}

		// ------------------------------------------------
		// 5) LOAD USER
		// ------------------------------------------------
		user, err := q.GetUserById(reqCtx, userID)
		if err != nil || !user.IsActive {
			oplog.Warn(reqCtx, "failed to load active user:", userID, err)
			c.Request = c.Request.WithContext(reqCtx)
			c.Next()
			return
		}

		// ------------------------------------------------
		// 6) ATTACH USER + SESSION TO CONTEXT
		// ------------------------------------------------
		reqCtx = context.WithValue(reqCtx, config.UserContextKey, &user)
		reqCtx = context.WithValue(reqCtx, config.SessContextKey, &session)

		// Attach user_id for oplog
		reqCtx = oplog.AttachUserID(reqCtx, user.ID.String())

		oplog.Info(
			reqCtx,
			"auth context attached",
			"user_id:", user.ID,
			"session_id:", sessionID,
			"role:", user.Role,
		)

		c.Request = c.Request.WithContext(reqCtx)
		c.Next()
	}
}

// ------------------------------------------------
// Helpers
// ------------------------------------------------

func logRefreshCookie(ctx context.Context, c *gin.Context, phase string) {

	_, err := c.Request.Cookie("refresh_token")

	if errors.Is(err, http.ErrNoCookie) {
		oplog.Info(ctx, "refresh_token cookie missing:", phase)
		return
	}

	if err != nil {
		oplog.Error(ctx, "failed to read refresh_token cookie:", phase, err)
		return
	}

	oplog.Info(ctx, "refresh_token cookie present:", phase)
}

func UserFromCtx(ctx context.Context) (*sql_models.User, bool) {
	u, ok := ctx.Value(config.UserContextKey).(*sql_models.User)
	return u, ok
}
