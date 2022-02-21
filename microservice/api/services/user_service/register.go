package userService

import (
	"log"
	"microservice/api/common"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func register(db *sqlx.DB, w http.ResponseWriter, r *http.Request, username string, password string) {
	salt := generateSalt(saltLength)

	passwordHash := hashPw([]byte(password), salt)

	_, err := getUserFromName(db, username)

	if err == nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "Username is already in use")
		return
	}

	log.Println("Created new password salt")

	var userId int
	err = db.QueryRow("INSERT INTO users(username, pw_hash, salt) VALUES($1, $2, $3) RETURNING id", username, string(passwordHash), string(salt)).Scan(&userId)

	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "Failed to create account")
		return
	}

	log.Printf("Created new user")

	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "Failed to create account")
		return
	}

	session, err := createSession(db, userId)

	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "Failed to create account")
		return
	}

	// Not needed, but better be explicit!
	w.Header().Set("Content-Type", "text/plain; charset")

	common.TryWriteResponse(w, strconv.Itoa(session.ID))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
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
	register(db, w, r, username.(string), password.(string))
}
