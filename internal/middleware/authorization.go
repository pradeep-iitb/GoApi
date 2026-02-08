package middleware

import (
	"errors"
	"net/http"

	"github.com/pradeep-iitb/GoApi/api"
	"github.com/pradeep-iitb/GoApi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Invalid username or password. ")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
	var username string = r.URL.Query().Get("username")
	var token = r.Header.Get("Authorization")

	if token == "" || username == "" {
		log.Error("Unauthorized access attempt. Missing token or username.")
		api.RequestErrorHandler(w, UnauthorizedError)
		return
	}

	var database *tools.DatabaseInterface 
	database , err := tools.NewDatabase()
	if err != nil {
		api.InternalServerErrorHandler(w,err)
		return 
	}

	var loginDetails *tools.LoginDetails
	loginDetails = (*database).GetUserLoginDetails(username)

	if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
		api.RequestErrorHandler(w, UnauthorizedError)
		return
	}

	next.ServeHTTP(w,r)

  })
}