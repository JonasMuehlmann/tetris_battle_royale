package userService

import (
	"net/http"
)

func MakeUserServiceMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", LoginHandler)
	mux.HandleFunc("/isLogin", IsLoginHandler)
	mux.HandleFunc("/logout", LogoutHandler)

	return mux
}
