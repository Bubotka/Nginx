package service

type Auther interface {
	Register(in RegisterIn) RegisterOut
	Login(in LoginIn) LoginOut
}

type RegisterIn struct {
	Username string
	Password string
}

type RegisterOut struct {
	Error error
}

type LoginIn struct {
	Username string
	Password string
}
type LoginOut struct {
	Token string
	Error error
}
