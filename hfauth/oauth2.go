package hfauth

type AuthHandler interface {
	SetToken(token string)
}
