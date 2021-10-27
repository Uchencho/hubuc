package db

import "github.com/Uchencho/hubuc/internal"

type GetUserByEmailFunc func(email string) (internal.User, error)

type InsertUserFunc func(u internal.User) error
