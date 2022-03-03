package drivingAdapters

import (
	"log"
	common "microservice/internal"
	drivingPorts "microservice/internal/core/driving_ports"
	"net/http"

	"github.com/gorilla/mux"
)

type UserServiceRestAdapter struct {
	Service drivingPorts.UserServicePort
	Logger  *log.Logger
}

func (adapter UserServiceRestAdapter) IsLoginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sessionID, err := adapter.Service.IsLoggedIn(vars["username"])
	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))
	} else {
		common.TryWriteResponse(w, "{sessionID: "+sessionID+"}")
	}
}

func (adapter UserServiceRestAdapter) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusBadRequest)
	requestBody, _ := common.UnmarshalRequestBody(r)

	username, okUsername := requestBody["username"]
	if !okUsername {
		w.WriteHeader(http.StatusBadRequest)
		common.TryWriteResponse(w, common.MakeJsonError("Missing username"))

		return
	}

	password, okPassword := requestBody["password"]
	if !okPassword {
		w.WriteHeader(http.StatusBadRequest)
		common.TryWriteResponse(w, common.MakeJsonError("Missing password"))

		return
	}

	sessionID, err := adapter.Service.Login(username.(string), password.(string))
	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))
	}

	common.TryWriteResponse(w, "{sessionID: "+sessionID+"}")
}

func (adapter UserServiceRestAdapter) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := common.UnmarshalRequestBody(r)

	sessionID, okSessionID := requestBody["sessionId"]
	if !okSessionID {
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError("Missing username"))

		return
	}

	err := adapter.Service.Logout(sessionID.(string))
	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))
	}

	common.TryWriteResponse(w, "{message: \"User logged out\"}")
}

func (adapter UserServiceRestAdapter) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusBadRequest)
	requestBody, _ := common.UnmarshalRequestBody(r)

	username, okUsername := requestBody["username"]
	if !okUsername {
		w.WriteHeader(http.StatusBadRequest)
		common.TryWriteResponse(w, common.MakeJsonError("Missing username"))

		return
	}

	password, okPassword := requestBody["password"]
	if !okPassword {
		w.WriteHeader(http.StatusBadRequest)
		common.TryWriteResponse(w, common.MakeJsonError("Missing username"))

		return
	}

	adapter.Logger.Println("Received registration request")
	userID, err := adapter.Service.Register(username.(string), password.(string))

	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError("Failed to register"))
	}

	common.TryWriteResponse(w, "{message: \""+userID+"\"}")
}

func (adapter UserServiceRestAdapter) Run() {
	mux := mux.NewRouter()

	mux.Handle("/", http.FileServer(http.Dir("../client/build/")))

	// NOTE: The api gateay should contain a prefix user/, which is stripped before forwarding
	mux.HandleFunc("/login", adapter.LoginHandler).Methods("POST")
	mux.HandleFunc("/register", adapter.RegisterHandler).Methods("POST")
	mux.HandleFunc("/isLogin/{username:[a-zA-Z0-9]+}", adapter.IsLoginHandler).Methods("GET")
	mux.HandleFunc("/logout", adapter.LogoutHandler).Methods("DELETE")

	adapter.Logger.Println("Starting server on Port 8080")
	log.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
}
