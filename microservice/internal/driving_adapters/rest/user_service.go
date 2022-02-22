package drivingAdapters

import (
	"log"
	common "microservice/internal"
	drivingPorts "microservice/internal/driving_ports"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserServiceRestAdapter struct {
	service drivingPorts.UserServicePort
}

func (adapter *UserServiceRestAdapter) IsLoginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sessionId, err := adapter.service.IsLoggedIn(vars["userID"])
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, err.Error())
	} else {
		common.TryWriteResponse(w, "User logged in with ID "+strconv.FormatInt(int64(sessionId), 10))
	}
}

func (adapter *UserServiceRestAdapter) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusBadRequest)
	requestBody, _ := common.UnmarshalRequestBody(r)

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

	sessionID, err := adapter.service.Login(username.(string), password.(string))
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, err.Error())
	}

	common.TryWriteResponse(w, "User logged in with session ID "+strconv.FormatInt(int64(sessionID), 10))
}

func (adapter *UserServiceRestAdapter) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, _ := strconv.ParseInt(vars["userID"], 10, 32)

	err := adapter.service.Logout(int(userID))
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, err.Error())
	}

	common.TryWriteResponse(w, "User logged out")
}

func (adapter *UserServiceRestAdapter) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusBadRequest)
	requestBody, _ := common.UnmarshalRequestBody(r)

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
	adapter.service.Register(username.(string), password.(string))
}

func (adapter *UserServiceRestAdapter) Run() {
	mux := mux.NewRouter()

	mux.Handle("/", http.FileServer(http.Dir("../client/build/")))

	// NOTE: The api gateay should contain a prefix user/, which is stripped before forwarding
	mux.HandleFunc("/login", adapter.LoginHandler).Methods("POST")
	mux.HandleFunc("/register", adapter.RegisterHandler).Methods("POST")
	mux.HandleFunc("/isLogin/{userId:[0-9]+}", adapter.IsLoginHandler).Methods("GET")
	mux.HandleFunc("/logout/{userId:[0-9]+}", adapter.LogoutHandler).Methods("DELETE")

	log.Println("Starting server on Port 8080")
	log.Fatalf("server failed to start: %v", http.ListenAndServe(":8080", mux))
}
