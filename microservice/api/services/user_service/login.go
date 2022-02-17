package userService

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"microservice/api/common"
	"net/http"
	"strconv"

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

func login(db *sqlx.DB, w http.ResponseWriter, r *http.Request, username string, password string) {
	w.WriteHeader(http.StatusOK)
	var passwordHash []byte
	var salt []byte

	user := User{}

	err := db.Get(&user, "SELECT * FROM users WHERE username = $1", username)

	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "Unknown username")
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
		common.TryWriteResponse(w, "Failed to login")
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
	usernameParam, okUsername := r.URL.Query()["username"]

	if !okUsername {
		common.TryWriteResponse(w, "Missing username or password")
		return
	}

	username := usernameParam[0]

	passwordParam, okPassword := r.URL.Query()["password"]

	var wantsToRegister bool

	wantsToRegisterParam, okWantsToRegister := r.URL.Query()["wantsToRegister"]

	if !okWantsToRegister {
		wantsToRegister = false
	} else {
		var err error

		wantsToRegister, err = strconv.ParseBool(wantsToRegisterParam[0])

		if err != nil {
			log.Printf("Error: %v", err)
			common.TryWriteResponse(w, "Invalid value for parameter 'wantsToRegister', needs to be 'true' or 'false'")
			return
		}
	}

	db, err := sqlx.Open("postgres", connectionString)

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Printf("Failed to open db: %v", err)

		return
	}

	if okUsername && !okPassword && !wantsToRegister {
		log.Println("Received login check request")

		isLoggedIn(db, w, r, username)
	} else if okUsername && okPassword && !wantsToRegister {
		log.Println("Received login request")

		password := passwordParam[0]
		login(db, w, r, username, password)
	} else if okUsername && okPassword && wantsToRegister {
		log.Println("Received registration request")

		password := passwordParam[0]
		register(db, w, r, username, password)
	} else {
		log.Println("Received invalid request")

		_, err := w.Write([]byte("Invalid request"))

		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}
