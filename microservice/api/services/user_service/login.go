package userService

import (
	"bytes"
	"encoding/binary"
	"log"
	"microservice/api/common"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request, username string, password string) {
	w.WriteHeader(http.StatusOK)
	var passwordHash []byte
	var salt []byte

	user, err := common.GetUserFromName(username)

	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "User does not exist")
		return
	}

	salt = []byte(user.Salt)
	passwordHash = []byte(user.PwHash)

	inputHash := hashPw([]byte(password), salt)

	if bytes.Compare(inputHash, passwordHash) != 0 {
		common.TryWriteResponse(w, "InvalID username or password")
		return
	}

	sessionID, err := common.CreateSession(user.ID)

	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "User already logged in")
		return
	}

	// Not needed, but better be explicit!
	w.Header().Set("Content-Type", "text/plain; charset")

	var sessionIDEnc []byte

	binary.LittleEndian.PutUint32(sessionIDEnc, uint32(sessionID))

	common.TryWriteResponse(w, string(sessionIDEnc))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusBadRequest)
	requestBody, err := common.UnmarshalRequestBody(r)

	username, okUsername := requestBody["username"]
	if !okUsername {
		common.TryWriteResponse(w, "Missing username")
		return
	}

	password, okPassword := requestBody["password"]
	if !okPassword {
		common.TryWriteResponse(w, "Missing password")
		return
	}

	log.Println("Received registration request")
	login(w, r, username.(string), password.(string))
}
