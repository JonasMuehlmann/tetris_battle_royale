package userService

import (
	"net/http"
)

func MakeUserServiceMux() http.Handler {
	mux := http.NewServeMux()

	// TODO: The routers can be simplified with gorilla/mux
	mux.HandleFunc("/login", LoginHandler)
	mux.HandleFunc("/isLogin", IsLoginHandler)
	mux.HandleFunc("/logout", LogoutHandler)

	return mux
}
