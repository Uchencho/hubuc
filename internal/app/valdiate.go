package app

import (
	"encoding/json"

	"github.com/Uchencho/hubuc/pkg"
	"github.com/pkg/errors"
)

func validateCreateUserRequest(payload pkg.CreateUserRequest) error {
	m := make(map[string]string)

	if payload.Email == "" {
		m["email"] = "email is required"
	}
	if payload.Name == "" {
		m["name"] = "name is required"
	}
	if payload.Password == "" {
		m["password"] = "password is required"
	}
	if payload.Password != "" && len(payload.Password) < 6 {
		m["password"] = "password must be at least 6 characters"
	}
	if len(m) != 0 {
		bb, _ := json.Marshal(m)
		return errors.Errorf(string(bb))
	}
	return nil
}
