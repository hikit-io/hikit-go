package hkauth

type AuthHandler interface {
	SetToken(token string)
}
