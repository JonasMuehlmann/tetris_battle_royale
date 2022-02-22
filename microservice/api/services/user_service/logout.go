package userService

import (
	"log"
	"microservice/api/common"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func logout(w http.ResponseWriter, sessionID int) {

	err := common.DeleteSession(sessionID)
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "Failed to end session")
	}

	common.TryWriteResponse(w, "Successfully ended session")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, _ := strconv.ParseInt(vars["userID"], 10, 32)
	logout(w, int(userID))
}
