package auth

import (
	"context"
	"time"

	"github.com/pquerna/otp/totp"
)

func (s *Service) Register(ctx context.Context, in CreateUserInput) (*RegisterResult, error) {
	tx, qtx, err := s.DB.BeginTx(ctx)
	if err != nil {
		return nil, ErrInternal
	}
	defer tx.Rollback()

	hashed, err := s.Crypto.HashPassword(in.Password)
	if err != nil {
		return nil, ErrInternal
	}

	user, err := qtx.CreateUser(ctx, CreateUserParams{
		Name:      in.Name,
		Username:  in.Username,
		Email:     in.Email,
		Password:  hashed,
		Mobile:    in.Mobile,
		Role:      in.Role,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, ErrUserExists
	}

	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "CoolAid",
		AccountName: in.Email,
	})

	errTOTP := qtx.CreateUserTotp(ctx, CreateUserTotpParams{
		UserID: user.ID,
		Secret: key.Secret(),
	})

	if errTOTP != nil {
		return nil, ErrInternal
	}

	tx.Commit()

	return &RegisterResult{
		UserID:     ID(user.ID),
		Username:   user.Username,
		Email:      user.Email,
		OtpSecret:  key.Secret(),
		OtpAuthURL: key.URL(),
	}, nil
}
