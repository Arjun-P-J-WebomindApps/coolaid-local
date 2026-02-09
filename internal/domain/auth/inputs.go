package auth

// CreateUserInput is the input required to register a new user.
// This is a domain-level DTO, not a DB model.
type CreateUserInput struct {
	Name     string
	Username string
	Email    string
	Password string
	Mobile   string
	Role     string
}
