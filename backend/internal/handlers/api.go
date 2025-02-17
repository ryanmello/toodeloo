package handlers

import (
	"toodeloo/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r * chi.Mux){
	r.Use(chimiddle.StripSlashes)
	
	r.Route("/account", func(router chi.Router){
		router.Use(middleware.Authorization)
		
		// endpoint at '/account/coins'
		router.Get("/coins", GetCoinBalance)
		router.Get("/amount", GetCoinBalance)
	})

	r.Route("/acc", func(router chi.Router){
		router.Use(middleware.Authorization)
		
		// endpoint at '/acc/coins'
		router.Get("/coins", GetCoinBalance)
		router.Get("/amount", GetCoinBalance)
	})
}