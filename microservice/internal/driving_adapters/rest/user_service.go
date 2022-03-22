package drivingAdapters

import (
	"fmt"
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
		adapter.Logger.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	user, err := adapter.Service.UserRepository.GetUserFromName(vars["username"])

	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	common.TryWriteResponse(w, fmt.Sprintf(`{"sessionID": "%v", "userID": "%v", "username": "%v"}`, sessionID, user.ID, user.Username))
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
		adapter.Logger.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	user, err := adapter.Service.UserRepository.GetUserFromName(username.(string))

	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	common.TryWriteResponse(w, fmt.Sprintf(`{"sessionID": "%v", "userID": "%v", "username": "%v"}`, sessionID, user.ID, user.Username))
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
		adapter.Logger.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	common.TryWriteResponse(w, `{"message": "User logged out"}`)
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
		adapter.Logger.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError("Failed to register"))

		return
	}

	sessionID, err := adapter.Service.Login(username.(string), password.(string))
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	common.TryWriteResponse(w, fmt.Sprintf(`{"sessionID": "%v", "userID": "%v", "username": "%v"}`, sessionID, userID, username.(string)))

}

func (adapter UserServiceRestAdapter) Run() {
	mux := mux.NewRouter()

	mux.Handle("/", http.FileServer(http.Dir("../client/build/")))

	mux.HandleFunc("/login", adapter.LoginHandler).Methods("POST")
	mux.HandleFunc("/register", adapter.RegisterHandler).Methods("POST")
	mux.HandleFunc("/isLogin/{username:[a-zA-Z0-9]+}", adapter.IsLoginHandler).Methods("GET")
	// TODO: The id should be part of the URL, not the request body
	mux.HandleFunc("/logout", adapter.LogoutHandler).Methods("DELETE")

	adapter.Logger.Println("Starting server on Port 8080")
	adapter.Logger.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
}
