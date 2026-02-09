package auth

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
)

func (q *authQueries) GetUserTotpByUserID(
	ctx context.Context,
	userID auth.ID,
) (*auth.UserTotpRow, error) {

	uid, err := userID.ToUUID(ctx)
	if err != nil {
		return nil, err
	}

	row, err := q.q.GetUserTotpByUserID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return &auth.UserTotpRow{
		UserID: userID,
		Secret: row.Secret,
	}, nil
}

func (q *authQueries) CreateUserTotp(
	ctx context.Context,
	p auth.CreateUserTotpParams,
) error {

	uid, err := p.UserID.ToUUID(ctx)
	if err != nil {
		return err
	}

	_, errTOTP := q.q.CreateUserTotp(ctx, sqlc.CreateUserTotpParams{
		UserID: uid,
		Secret: p.Secret,
	})

	return errTOTP
}
