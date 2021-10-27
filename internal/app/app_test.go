package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Uchencho/hubuc/internal"
	"github.com/Uchencho/hubuc/internal/app"
	"github.com/Uchencho/hubuc/pkg"
)

type row struct {
	name  string
	in    interface{}
	out   interface{}
	mocks app.Option
}

func TestCreateUser(t *testing.T) {

	table := []row{
		{
			name: "successfully create user",
			in:   pkg.CreateUserRequest{Email: "uche@gmail.com"},
			out:  http.StatusOK,
		},
	}

	var (
		getUserInvoked, insertUserInvoked bool
	)

	mockOptions := func(oa *app.OptionalArgs) {
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
	}

	for _, entry := range table {
		a := app.New(mockOptions)
		ht := httptest.NewServer(a.Handler())
		defer ht.Close()

		req := pkg.CreateUserRequest{Email: "uche@gmail.com"}
		// bytes.NewReader()
		// http.NewRequest(http.MethodPost, fmt.Sprintf("%s/user", ht.URL), )
	}

}
