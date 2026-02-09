package models

import (
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/master/company"
)

type Service struct {
	DB             DB
	CompanyService *company.Service
}

func NewService(db DB, companyService *company.Service) *Service {
	return &Service{DB: db, CompanyService: companyService}

}
