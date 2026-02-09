package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/internal/shared/sqlnull"
)

func (q *authQueries) GetUserByEmail(
	ctx context.Context,
	email string,
) (*auth.UserRow, error) {

	row, err := q.q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return mapUserRow(row), nil
}

func (q *authQueries) GetUserByUsername(
	ctx context.Context,
	username string,
) (*auth.UserRow, error) {

	row, err := q.q.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return mapUserRow(row), nil
}

func (q *authQueries) GetUserById(
	ctx context.Context,
	id auth.ID,
) (*auth.UserRow, error) {

	uid, err := id.ToUUID(ctx)
	if err != nil {
		return nil, err
	}

	row, err := q.q.GetUserById(ctx, uid)
	if err != nil {
		return nil, err
	}

	return mapUserRow(row), nil
}

func (q *authQueries) CreateUser(
	ctx context.Context,
	p auth.CreateUserParams,
) (*auth.UserRow, error) {

	row, err := q.q.CreateUser(ctx, sqlc.CreateUserParams{
		ID:        uuid.MustParse(string(p.ID)),
		Name:      p.Name,
		Username:  p.Username,
		Email:     p.Email,
		Password:  p.Password,
		Mobile:    p.Mobile,
		Role:      p.Role,
		IsActive:  p.IsActive,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}

	return mapUserRow(row), nil
}

func (q *authQueries) UpdateUserPassword(
	ctx context.Context,
	userID auth.ID,
	hashedPassword string,
) error {

	uid, err := userID.ToUUID(ctx)
	if err != nil {
		return err
	}

	return q.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:       uid,
		Password: sqlnull.String(&hashedPassword),
	})
}
