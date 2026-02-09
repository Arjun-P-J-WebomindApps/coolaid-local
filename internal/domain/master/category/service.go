package category

type Service struct {
	DB DB
}

func NewService(db DB) *Service {
	return &Service{DB: db}
}
