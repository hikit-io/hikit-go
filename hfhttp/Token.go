package hfhttp

type TokenSource interface {
	Token() (*Token, error)
}

type Token struct {
}
