package auth

import "context"

func (s *Service) CheckUserStatus(
	ctx context.Context,
	writer CookieWriter,
) (*AuthResult, error) {

	claims, err := s.Crypto.ValidateJWTFromContext(ctx)
	if err != nil {
		writer.ClearAuth()
		return nil, nil
	}

	session, err := s.DB.Queries().GetUserSessionById(ctx, ID(claims.SessionID))
	if err != nil {
		writer.ClearAuth()
		return nil, nil
	}

	user, err := s.DB.Queries().GetUserById(ctx, session.UserID)
	if err != nil || !user.IsActive {
		writer.ClearAuth()
		return nil, nil
	}

	return &AuthResult{
		UserID:   user.ID,
		Username: user.Username,
	}, nil
}

func (s *Service) Logout(ctx context.Context, writer CookieWriter) error {
	claims, err := s.Crypto.ValidateJWTFromContext(ctx)
	if err != nil {
		writer.ClearAuth()
		return nil
	}

	_ = s.DB.Queries().DeleteRefreshTokensBySession(ctx, ID(claims.SessionID))
	_ = s.DB.Queries().DeleteUserSession(ctx, ID(claims.SessionID))

	writer.ClearAuth()
	return nil
}
