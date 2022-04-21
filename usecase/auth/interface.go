package auth

type AuthUseCaseInterface interface {
	Login(email string, password string) (string, error)
}
