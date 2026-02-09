package auth

import "context"

func (s *Service) RefreshAuth(
	ctx context.Context,
	writer CookieWriter,
	rawRefresh string,
) error {

	hash, _ := s.Crypto.HashToken(rawRefresh)
	rt, err := s.DB.Queries().GetRefreshTokenByHash(ctx, hash)
	if err != nil {
		return ErrUnauthorized
	}

	session, err := s.DB.Queries().GetUserSessionById(ctx, rt.SessionID)
	if err != nil {
		return ErrUnauthorized
	}

	jwt, _ := s.Crypto.GenerateJWT(ctx, session.SessionID, rt.UserID)
	writer.SetAuth(jwt, rawRefresh)

	return nil
}
