package ticketservice

import "github.com/webomindapps-dev/coolaid-backend/db"

type Service struct {
	DB *db.DBContext
}

func NewService(dbCtx *db.DBContext) *Service {
	return &Service{
		DB: dbCtx,
	}
}
