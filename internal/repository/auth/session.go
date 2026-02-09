package auth

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
)

func (q *authQueries) GetUserSessionByUserId(
	ctx context.Context,
	userID auth.ID,
) (*auth.SessionRow, error) {

	uid, err := userID.ToUUID(ctx)
	if err != nil {
		return nil, err
	}

	row, err := q.q.GetUserSessionByUserId(ctx, uid)
	if err != nil {
		return nil, err
	}

	return mapSessionRow(row), nil
}

func (q *authQueries) GetUserSessionById(
	ctx context.Context,
	sessionID auth.ID,
) (*auth.SessionRow, error) {

	sid, err := sessionID.ToUUID(ctx)
	if err != nil {
		return nil, err
	}

	row, err := q.q.GetUserSessionById(ctx, sid)
	if err != nil {
		return nil, err
	}

	return mapSessionRow(row), nil
}

func (q *authQueries) CreateUserSession(
	ctx context.Context,
	p auth.CreateSessionParams,
) (*auth.SessionRow, error) {

	uid, err := p.UserID.ToUUID(ctx)
	if err != nil {
		return nil, err
	}

	sid, err := p.SessionID.ToUUID(ctx)
	if err != nil {
		return nil, err
	}

	row, err := q.q.CreateUserSession(ctx, sqlc.CreateUserSessionParams{
		UserID:    uid,
		SessionID: sid,
		IpAddress: p.IP,
		UserAgent: p.UserAgent,
		ExpiresAt: p.ExpiresAt,
	})
	if err != nil {
		return nil, err
	}

	return mapSessionRow(row), nil
}

func (q *authQueries) DeleteUserSession(
	ctx context.Context,
	sessionID auth.ID,
) error {

	sid, err := sessionID.ToUUID(ctx)
	if err != nil {
		return err
	}

	return q.q.DeleteUserSession(ctx, sid)
}
