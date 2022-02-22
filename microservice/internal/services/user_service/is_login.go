package userService

import (
	"log"
	"microservice/api/common"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func isLoggedIn(w http.ResponseWriter, r *http.Request, username string) {
	user, err := common.GetUserFromName(username)

	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "User does not exist")
		return
	}

	session, err := common.GetSession(user.ID)

	// TODO: Check if session expired

	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "User not logged in")
	} else {
		common.TryWriteResponse(w, "User logged in with ID "+strconv.FormatInt(int64(session.ID), 10))
	}
}

func IsLoginHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	isLoggedIn(w, r, vars["userID"])
}
