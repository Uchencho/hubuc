package app_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Uchencho/hubuc/internal"
	"github.com/Uchencho/hubuc/internal/app"
	"github.com/Uchencho/hubuc/pkg"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type row struct {
	name       string
	in         interface{}
	out        interface{}
	mocks      app.Option
	customTest func(resp *http.Response, name string)
}

func TestCreateUser(t *testing.T) {

	var (
		getUserInvoked, insertUserInvoked bool
	)

	table := []row{
		{
			name: "successfully create user",
			in:   pkg.CreateUserRequest{Email: "uche@gmail.com", Name: "uche", Password: "testPass"},
			out:  http.StatusNoContent,
			mocks: func(oa *app.OptionalArgs) {
				oa.GetUser = func(email string) (internal.User, error) {
					getUserInvoked = true
					return internal.User{}, errors.Errorf("user does not exist")
				}

				oa.HashPassword = func(password string) string {
					return "hashed-password"
				}

				oa.InsertUser = func(u internal.User) error {
					insertUserInvoked = true
					return nil
				}
			},
			customTest: func(resp *http.Response, name string) {
				t.Run(fmt.Sprintf("%s-http response is as expected", name), func(t *testing.T) {
					assert.Equal(t, http.StatusNoContent, resp.StatusCode)
				})
				t.Run(fmt.Sprintf("%s-get user is invoked", name), func(t *testing.T) {
					assert.True(t, getUserInvoked)
				})
				t.Run(fmt.Sprintf("%s-insert user is invoked", name), func(t *testing.T) {
					assert.True(t, insertUserInvoked)
				})
			},
		},
		{
			name: "fails to create user due to user already existing",
			in:   pkg.CreateUserRequest{Email: "uche@gmail.com", Name: "uche", Password: "testPass"},
			out:  http.StatusBadRequest,
			mocks: func(oa *app.OptionalArgs) {
				oa.GetUser = func(email string) (internal.User, error) {
					getUserInvoked = true
					return internal.User{}, nil
				}

				oa.HashPassword = func(password string) string {
					return "hashed-password"
				}

				oa.InsertUser = func(u internal.User) error {
					insertUserInvoked = true
					return nil
				}
			},
			customTest: func(resp *http.Response, name string) {
				t.Run(fmt.Sprintf("%s-http response is as expected", name), func(t *testing.T) {
					assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
				})
				t.Run(fmt.Sprintf("%s-get user is invoked", name), func(t *testing.T) {
					assert.True(t, getUserInvoked)
				})
				t.Run(fmt.Sprintf("%s-insert user is NOT invoked", name), func(t *testing.T) {
					assert.False(t, insertUserInvoked)
				})
			},
		},
		{
			name: "fails to create user due to invalid payload",
			in:   pkg.CreateUserRequest{Email: "uche@gmail.com"},
			out:  http.StatusBadRequest,
			mocks: func(oa *app.OptionalArgs) {
				oa.GetUser = func(email string) (internal.User, error) {
					getUserInvoked = true
					return internal.User{}, nil
				}

				oa.HashPassword = func(password string) string {
					return "hashed-password"
				}

				oa.InsertUser = func(u internal.User) error {
					insertUserInvoked = true
					return nil
				}
			},
			customTest: func(resp *http.Response, name string) {
				t.Run(fmt.Sprintf("%s-http response is as expected", name), func(t *testing.T) {
					assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
				})
				t.Run(fmt.Sprintf("%s-get user is NOT invoked", name), func(t *testing.T) {
					assert.False(t, getUserInvoked)
				})
				t.Run(fmt.Sprintf("%s-insert user is NOT invoked", name), func(t *testing.T) {
					assert.False(t, insertUserInvoked)
				})
			},
		},
	}

	for _, entry := range table {
		a := app.New(entry.mocks)
		ht := httptest.NewServer(a.Handler())
		defer ht.Close()

		log.Printf("\n\n\nRunning Test : %s \n\n\n", entry.name)

		bb, _ := json.Marshal(entry.in)
		rr := bytes.NewReader(bb)
		req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/user", ht.URL), rr)
		resp, _ := http.DefaultClient.Do(req)

		entry.customTest(resp, entry.name)

		getUserInvoked, insertUserInvoked = false, false
	}
}

func TestCreateUserV2(t *testing.T) {

	var (
		getUserInvoked, insertUserInvoked bool
	)

	table := []row{
		{
			name: "successfully create user",
			in:   pkg.CreateUserRequest{Email: "uche@gmail.com", Name: "uche", Password: "testPass"},
			out:  http.StatusNoContent,
			mocks: func(oa *app.OptionalArgs) {
				oa.GetUser = func(email string) (internal.User, error) {
					getUserInvoked = true
					return internal.User{}, errors.Errorf("user does not exist")
				}

				oa.HashPassword = func(password string) string {
					return "hashed-password"
				}

				oa.InsertUser = func(u internal.User) error {
					insertUserInvoked = true
					return nil
				}
			},
			customTest: func(resp *http.Response, name string) {
				t.Run(fmt.Sprintf("%s-http response is as expected", name), func(t *testing.T) {
					assert.Equal(t, http.StatusNoContent, resp.StatusCode)
				})
				t.Run(fmt.Sprintf("%s-get user is invoked", name), func(t *testing.T) {
					assert.True(t, getUserInvoked)
				})
				t.Run(fmt.Sprintf("%s-insert user is invoked", name), func(t *testing.T) {
					assert.True(t, insertUserInvoked)
				})
			},
		},
		{
			name: "fails to create user due to user already existing",
			in:   pkg.CreateUserRequest{Email: "uche@gmail.com", Name: "uche", Password: "testPass"},
			out:  http.StatusBadRequest,
			mocks: func(oa *app.OptionalArgs) {
				oa.GetUser = func(email string) (internal.User, error) {
					getUserInvoked = true
					return internal.User{}, nil
				}

				oa.HashPassword = func(password string) string {
					return "hashed-password"
				}

				oa.InsertUser = func(u internal.User) error {
					insertUserInvoked = true
					return nil
				}
			},
			customTest: func(resp *http.Response, name string) {
				t.Run(fmt.Sprintf("%s-http response is as expected", name), func(t *testing.T) {
					assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
				})
				t.Run(fmt.Sprintf("%s-get user is invoked", name), func(t *testing.T) {
					assert.True(t, getUserInvoked)
				})
				t.Run(fmt.Sprintf("%s-insert user is NOT invoked", name), func(t *testing.T) {
					assert.False(t, insertUserInvoked)
				})
			},
		},
		{
			name: "fails to create user due to invalid payload",
			in:   pkg.CreateUserRequest{Email: "uche@gmail.com"},
			out:  http.StatusBadRequest,
			mocks: func(oa *app.OptionalArgs) {
				oa.GetUser = func(email string) (internal.User, error) {
					getUserInvoked = true
					return internal.User{}, nil
				}

				oa.HashPassword = func(password string) string {
					return "hashed-password"
				}

				oa.InsertUser = func(u internal.User) error {
					insertUserInvoked = true
					return nil
				}
			},
			customTest: func(resp *http.Response, name string) {
				t.Run(fmt.Sprintf("%s-http response is as expected", name), func(t *testing.T) {
					assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
				})
				t.Run(fmt.Sprintf("%s-get user is NOT invoked", name), func(t *testing.T) {
					assert.False(t, getUserInvoked)
				})
				t.Run(fmt.Sprintf("%s-insert user is NOT invoked", name), func(t *testing.T) {
					assert.False(t, insertUserInvoked)
				})
			},
		},
	}

	for _, entry := range table {
		a := app.New(entry.mocks)
		ht := httptest.NewServer(a.Handler())
		defer ht.Close()

		log.Printf("\n\n\nRunning Test : %s \n\n\n", entry.name)

		bb, _ := json.Marshal(entry.in)
		rr := bytes.NewReader(bb)
		req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/v2/user", ht.URL), rr)
		resp, _ := http.DefaultClient.Do(req)

		entry.customTest(resp, entry.name)

		getUserInvoked, insertUserInvoked = false, false
	}
}
