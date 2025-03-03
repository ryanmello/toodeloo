package handlers

import (
	"toodeloo/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// the 'H' tells the complier that the function can be exported
// a lowercase letter tells the complier that the function is private and can't be exported
func Handler(r * chi.Mux){
	// removes trailing slashes from the url
	// trailing slahes lead to a 404 error
	r.Use(chimiddle.StripSlashes)
	
	r.Route("/account", func(router chi.Router){
		router.Use(middleware.Authorization)
		
		// endpoint at '/account/coins'
		router.Get("/coins", GetCoinBalance)
	})
}