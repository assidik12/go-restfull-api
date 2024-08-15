package domain

type Account struct {
	ID       int
	Username string
	Email    string
	Password string
}

type AuthToken struct {
	Token string
}
