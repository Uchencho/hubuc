package workflow

import (
	"fmt"

	"github.com/Uchencho/hubuc/internal"
	"github.com/Uchencho/hubuc/internal/db"
	"github.com/Uchencho/hubuc/internal/middleware"
	"github.com/Uchencho/hubuc/pkg"
	"github.com/pkg/errors"
)

//
type CreateUserFlow interface {
	GetUser(email string) (internal.User, error)
	InsertUser(u internal.User) error
	HashPassword(password string) string
}

func CreateUser(getUser db.GetUserByEmailFunc,
	hashPassword middleware.HashPasswordFunc,
	insertUser db.InsertUserFunc) CreateUserFunc {
	return func(req pkg.CreateUserRequest) error {

		_, err := getUser(req.Email)
		if err == nil {
			return fmt.Errorf("user alreeady exists, please login")
		}

		hashedPassword := hashPassword(req.Password)

		u := internal.User{
			Name:           req.Name,
			Email:          req.Email,
			HashedPassword: hashedPassword,
		}

		if err := insertUser(u); err != nil {
			return errors.Wrapf(err, "CreateUser - unable to insert user")
		}

		return nil
	}
}

func CreateUserV2(flow CreateUserFlow) CreateUserFunc {
	return func(req pkg.CreateUserRequest) error {
		_, err := flow.GetUser(req.Email)
		if err == nil {
			return fmt.Errorf("user alreeady exists, please login")
		}

		hashedPassword := flow.HashPassword(req.Password)

		u := internal.User{
			Name:           req.Name,
			Email:          req.Email,
			HashedPassword: hashedPassword,
		}

		if err := flow.InsertUser(u); err != nil {
			return errors.Wrapf(err, "CreateUser - unable to insert user")
		}

		return nil
	}
}
