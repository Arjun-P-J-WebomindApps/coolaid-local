package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/webomindapps-dev/coolaid-backend/internal/assets/mail"
)

func (s *Service) ForgotPassword(ctx context.Context, email string) error {
	user, err := s.DB.Queries().GetUserByEmail(ctx, email)
	if err != nil {
		return nil
	}

	_ = s.DB.Queries().DeleteUserOTPByUserId(ctx, user.ID)

	otp, _ := s.Crypto.GenerateNumericOTP(6)
	hash, _ := s.Crypto.HashToken(fmt.Sprintf("%s", otp))

	_ = s.DB.Queries().CreateOTP(ctx, CreateOTPParams{
		UserID:    user.ID,
		OtpCode:   hash,
		ExpiresAt: time.Now().Add(15 * time.Minute),
	})

	optHTML := mail.ForgotPasswordOTPTemplate(mail.ForgotPasswordOTPEmailData{
		CommonEmailData: mail.CommonEmailData{
			Year: time.Now().Year(),
		},
		OTP: otp,
	})

	s.Mailer.Send(user.Email, "Reset Password", optHTML)
	return nil
}

func (s *Service) ResetPassword(ctx context.Context, email, token, newPass string) error {
	tx, qtx, err := s.DB.BeginTx(ctx)
	if err != nil {
		return ErrInternal
	}
	defer tx.Rollback()

	user, err := qtx.GetUserByEmail(ctx, email)
	if err != nil {
		return ErrInvalidCredentials
	}

	otp, err := qtx.GetLatestOTPFromUser(ctx, user.ID)
	if err != nil || otp.IsUsed || time.Now().After(otp.ExpiresAt) {
		return ErrInvalidOTP
	}

	if err := s.Crypto.CompareToken(otp.OtpCode, token); err != nil {
		return ErrInvalidOTP
	}

	hashed, _ := s.Crypto.HashPassword(newPass)

	_ = qtx.UpdateUserPassword(ctx, user.ID, hashed)
	_ = qtx.MarkOTPAsUsed(ctx, user.ID)
	_ = qtx.DeleteUserOTPByUserId(ctx, user.ID)

	tx.Commit()
	return nil
}
