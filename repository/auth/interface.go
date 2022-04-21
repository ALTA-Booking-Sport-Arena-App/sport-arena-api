package auth

type AuthRepositoryInterface interface {
	Login(email string, password string) (string, error)
}
