package app

import (
	"net/http"

	"github.com/Uchencho/hubuc/internal/db"
	"github.com/Uchencho/hubuc/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

type App struct {
	CreateUser http.HandlerFunc
}

type OptionalArgs struct {
	GetUser      db.GetUserByEmailFunc
	HashPassword middleware.HashPasswordFunc
	InsertUser   db.InsertUserFunc
}

type Option func(oa *OptionalArgs)

func New(opts ...Option) App {
	o := OptionalArgs{
		GetUser:      db.GetUserByEmail(),
		HashPassword: middleware.HashPassword(),
		InsertUser:   db.InsertUser(),
	}

	for _, opt := range opts {
		opt(&o)
	}

	createUser := CreateUserHandler(o.GetUser, o.HashPassword, o.InsertUser)
	return App{CreateUser: createUser}
}

func (a *App) Handler() http.HandlerFunc {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/user", a.CreateUser)

	h := http.HandlerFunc(router.ServeHTTP)
	return h
}
