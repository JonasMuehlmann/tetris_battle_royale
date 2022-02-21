package userService

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"microservice/api/common"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	dbname   = "prod"
)

var connectionString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, username, dbname)

func login(db *sqlx.DB, w http.ResponseWriter, r *http.Request, username string, password string) {
	w.WriteHeader(http.StatusOK)
	var passwordHash []byte
	var salt []byte

	user := User{}

	err := db.Get(&user, "SELECT * FROM users WHERE username = $1", username)

	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "User does not exist")
		return
	}

	salt = []byte(user.Salt)
	passwordHash = []byte(user.PwHash)

	inputHash := hashPw([]byte(password), salt)

	if bytes.Compare(inputHash, passwordHash) != 0 {
		common.TryWriteResponse(w, "Invalid username or password")
		return
	}

	session, err := createSession(db, user.ID)

	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "User already logged in")
		return
	}

	// Not needed, but better be explicit!
	w.Header().Set("Content-Type", "text/plain; charset")

	var sessionIDEnc []byte

	binary.LittleEndian.PutUint32(sessionIDEnc, uint32(session.ID))

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

	db, err := sqlx.Open("postgres", connectionString)

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Printf("Failed to open db: %v", err)

		return
	}

	log.Println("Received registration request")
	login(db, w, r, username.(string), password.(string))
}
