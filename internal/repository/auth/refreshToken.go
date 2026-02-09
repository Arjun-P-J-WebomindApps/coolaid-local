package auth

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (q *authQueries) CreateRefreshToken(
	ctx context.Context,
	p auth.CreateRefreshTokenParams,
) error {

	uid, err := p.UserID.ToUUID(ctx)
	if err != nil {
		return err
	}

	sid, err := p.SessionID.ToUUID(ctx)
	if err != nil {
		return err
	}

	_, errToken := q.q.CreateRefreshToken(ctx, sqlc.CreateRefreshTokenParams{
		UserID:    uid,
		SessionID: sid,
		TokenHash: p.TokenHash,
		ExpiresAt: p.ExpiresAt,
	})

	return errToken
}

func (q *authQueries) DeleteRefreshTokensBySession(
	ctx context.Context,
	sessionID auth.ID,
) error {

	sid, err := sessionID.ToUUID(ctx)
	if err != nil {
		return err
	}

	return q.q.DeleteRefreshTokensBySession(ctx, sid)
}

func (q *authQueries) GetRefreshTokenByHash(
	ctx context.Context,
	hash string,
) (*auth.RefreshTokenRow, error) {

	row, err := q.q.GetRefreshTokenByHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	return &auth.RefreshTokenRow{
		ID:        auth.ID(row.ID.String()),
		UserID:    auth.ID(row.UserID.String()),
		SessionID: auth.ID(row.SessionID.String()),
		TokenHash: row.TokenHash,
		ExpiresAt: row.ExpiresAt,
		RevokedAt: sqlnull.TimePtr(row.RevokedAt),
		CreatedAt: row.CreatedAt,
	}, nil
}
