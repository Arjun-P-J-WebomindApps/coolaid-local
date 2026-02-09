package resolvers

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/company"
	models "github.com/webomindapps-dev/coolaid-backend/internal/domain/master/model"
)

func mapAuthError(err error) error {
	switch err {
	case auth.ErrInvalidCredentials:
		return gqlerror.Errorf("invalid credentials")
	case auth.ErrInvalidOTP:
		return gqlerror.Errorf("invalid otp")
	case auth.ErrExpiredOTP:
		return gqlerror.Errorf("otp expired")
	case auth.ErrUnauthorized:
		return gqlerror.Errorf("unauthorized")
	default:
		return gqlerror.Errorf("internal error")
	}
}

func mapCompanyError(err error) error {
	switch err {
	case company.ErrCompanyNotFound:
		return gqlerror.Errorf("company not found")

	case company.ErrCompanyExists:
		return gqlerror.Errorf("company already exists")

	case company.ErrInvalidInput:
		return gqlerror.Errorf("invalid input")

	default:
		return gqlerror.Errorf("internal error")
	}
}

func mapModelError(err error) error {
	switch err {

	case models.ErrModelNotFound:
		return gqlerror.Errorf("model not found")

	case models.ErrModelExists:
		return gqlerror.Errorf("model already exists")

	case models.ErrCompanyNotFound:
		return gqlerror.Errorf("company not found")

	case models.ErrInvalidInput:
		return gqlerror.Errorf("invalid input")

	default:
		return gqlerror.Errorf("internal error")
	}
}
