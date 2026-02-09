package auth

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
)

func (q *authQueries) CreateOTP(
	ctx context.Context,
	p auth.CreateOTPParams,
) error {

	uid, err := p.UserID.ToUUID(ctx)
	if err != nil {
		return err
	}

	_, errOTP := q.q.CreateOTP(ctx, sqlc.CreateOTPParams{
		UserID:    uid,
		OtpCode:   p.OtpCode,
		ExpiresAt: p.ExpiresAt,
	})

	return errOTP
}

func (q *authQueries) GetLatestOTPFromUser(
	ctx context.Context,
	userID auth.ID,
) (*auth.OTPRow, error) {

	uid, err := userID.ToUUID(ctx)
	if err != nil {
		return nil, err
	}

	row, err := q.q.GetLatestOTPFromUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	return &auth.OTPRow{
		UserID:    userID,
		OtpCode:   row.OtpCode,
		ExpiresAt: row.ExpiresAt,
		IsUsed:    row.IsUsed,
	}, nil
}

func (q *authQueries) DeleteUserOTPByUserId(
	ctx context.Context,
	userID auth.ID,
) error {

	uid, err := userID.ToUUID(ctx)
	if err != nil {
		return err
	}

	return q.q.DeleteUserOTPByUserId(ctx, uid)
}

func (q *authQueries) MarkOTPAsUsed(
	ctx context.Context,
	userID auth.ID,
) error {

	uid, err := userID.ToUUID(ctx)
	if err != nil {
		return err
	}

	return q.q.MarkOTPAsUsed(ctx, uid)
}
