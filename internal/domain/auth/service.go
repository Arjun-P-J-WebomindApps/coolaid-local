package auth

type Service struct {
	DB     DB
	Crypto Crypto
	Mailer Mailer
}

func NewService(db DB, crypto Crypto, mailer Mailer) *Service {
	return &Service{db, crypto, mailer}
}
