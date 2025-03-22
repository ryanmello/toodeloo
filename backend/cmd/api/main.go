package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Users...."))
}

func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("User created...."))
}

func main() {
	api := &api{ addr: ":8080" }
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr: api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
