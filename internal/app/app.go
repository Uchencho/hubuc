package app

import (
	"net/http"

	"github.com/Uchencho/hubuc/internal"
	"github.com/Uchencho/hubuc/internal/db"
	"github.com/Uchencho/hubuc/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

type App struct {
	CreateUser   http.HandlerFunc
	CreateUserV2 http.HandlerFunc
}

type OptionalArgs struct {
	GetUser      db.GetUserByEmailFunc
	HashPassword middleware.HashPasswordFunc
	InsertUser   db.InsertUserFunc
}

type wfs struct {
	arg OptionalArgs
}

func (w wfs) GetUser(email string) (internal.User, error) {
	return w.arg.GetUser(email)
}

func (w wfs) HashPassword(password string) string {
	return w.arg.HashPassword(password)
}

func (w wfs) InsertUser(u internal.User) error {
	return w.arg.InsertUser(u)
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
	flow := wfs{arg: o}

	createUser := CreateUserHandler(o.GetUser, o.HashPassword, o.InsertUser)
	createUSerV2 := CreateUserV2Handler(flow)
	return App{CreateUser: createUser, CreateUserV2: createUSerV2}
}

func (a *App) Handler() http.HandlerFunc {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/user", a.CreateUser)
	router.HandlerFunc(http.MethodPost, "/v2/user", a.CreateUser)

	h := http.HandlerFunc(router.ServeHTTP)
	return h
}
