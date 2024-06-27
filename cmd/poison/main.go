package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spoonboy-io/koan"
	"github.com/spoonboy-io/poison/internal/handlers"
	"github.com/spoonboy-io/poison/internal/router"
	"log"
	"net/http"
)

func main() {
	var logger *koan.Logger

	// state
	state := &handlers.State{
		Logger: logger,
		Host:   "localhost",
		Port:   9001,
	}

	// router
	rtr := mux.NewRouter()
	router.RegisterRoutes(state, rtr)

	logger.Info(fmt.Sprintf("starting Poison Server. Listening at 'http://%s:%d'", state.Host, state.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", state.Host, state.Port), rtr))
}
