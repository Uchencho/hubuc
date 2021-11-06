package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Uchencho/hubuc/internal/db"
	"github.com/Uchencho/hubuc/internal/middleware"
	"github.com/Uchencho/hubuc/internal/workflow"
	"github.com/Uchencho/hubuc/pkg"
)

func CreateUserHandler(getUser db.GetUserByEmailFunc,
	hashPassword middleware.HashPasswordFunc,
	insertUser db.InsertUserFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload pkg.CreateUserRequest

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			serveError(w, err, http.StatusBadRequest)
			return
		}

		if err := validateCreateUserRequest(payload); err != nil {
			log.Println(err)
			serveError(w, err, http.StatusBadRequest)
			return
		}

		wf := workflow.CreateUser(getUser, hashPassword, insertUser)
		if err := wf(payload); err != nil {
			log.Println(err)
			serveError(w, err, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func CreateUserV2Handler(flow workflow.CreateUserFlow) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload pkg.CreateUserRequest

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			serveError(w, err, http.StatusBadRequest)
			return
		}

		if err := validateCreateUserRequest(payload); err != nil {
			log.Println(err)
			serveError(w, err, http.StatusBadRequest)
			return
		}

		wf := workflow.CreateUserV2(flow)
		if err := wf(payload); err != nil {
			log.Println(err)
			serveError(w, err, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func serveError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}
