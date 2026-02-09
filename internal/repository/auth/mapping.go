package auth

import (
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func mapUserRow(row sqlc.User) *auth.UserRow {
	return &auth.UserRow{
		ID:       auth.ID(row.ID.String()),
		Name:     row.Name,
		Username: row.Username,
		Email:    row.Email,
		Password: row.Password,
		Mobile:   row.Mobile,
		Role:     row.Role,
		IsActive: row.IsActive,
	}
}

func mapSessionRow(row sqlc.UserSession) *auth.SessionRow {
	return &auth.SessionRow{
		SessionID: auth.ID(row.SessionID.String()),
		UserID:    auth.ID(row.UserID.String()),
		ExpiresAt: row.ExpiresAt,
		RevokedAt: sqlnull.TimePtr(row.RevokedAt),
	}
}
