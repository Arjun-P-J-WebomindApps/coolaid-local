package resolvers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
)

type GinCookieWriter struct {
	ctx *gin.Context
}

func NewGinCookieWriter(ctx *gin.Context) auth.CookieWriter {
	return &GinCookieWriter{ctx: ctx}
}

func (w *GinCookieWriter) SetAuth(access, refresh string) {
	http.SetCookie(w.ctx.Writer, &http.Cookie{
		Name:     "access_token",
		Value:    access,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(w.ctx.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    refresh,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
}

func (w *GinCookieWriter) ClearAuth() {
	http.SetCookie(w.ctx.Writer, &http.Cookie{
		Name:   "access_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	http.SetCookie(w.ctx.Writer, &http.Cookie{
		Name:   "refresh_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
