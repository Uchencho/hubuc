package db

import "github.com/Uchencho/hubuc/internal"

func GetUserByEmail() GetUserByEmailFunc {
	return func(email string) (internal.User, error) {
		return internal.User{Email: "uche@hubuc.com"}, nil
	}
}

func InsertUser() InsertUserFunc {
	return func(u internal.User) error {
		return nil
	}
}
