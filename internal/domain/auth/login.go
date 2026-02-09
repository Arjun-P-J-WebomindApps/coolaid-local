package auth

import "context"

func (s *Service) Login(ctx context.Context, email, password string) (*AuthResult, error) {
	user, err := s.DB.Queries().GetUserByEmail(ctx, email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := s.Crypto.ComparePassword(password, user.Password); err != nil {
		return nil, ErrInvalidCredentials
	}

	return &AuthResult{
		UserID:   user.ID,
		Username: user.Username,
	}, nil
}
