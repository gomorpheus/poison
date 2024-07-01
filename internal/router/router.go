package router

import (
	"github.com/gorilla/mux"
	"github.com/spoonboy-io/poison/internal/handlers"
	"net/http"
)

func RegisterRoutes(hs *handlers.State, r *mux.Router) {

	r.HandleFunc("/api/jwt/login", hs.Login).Methods("POST")
	r.HandleFunc("/api/arsys/v1/entry/COM:Company", hs.Company).Methods("GET")
	r.HandleFunc("/api/arsys/v1/entry/Group", hs.Group).Methods("GET")
	r.HandleFunc("/api/arsys/v1/entry/CTM:People", hs.People).Methods("GET")
	r.HandleFunc("/api/arsys/v1/entry/User", hs.User).Methods("GET")
	r.HandleFunc("/api/cmdb/v1.0/instances", hs.InstancePOST).Methods("POST")
	r.HandleFunc("/api/cmdb/v1.0/instances", hs.InstancePATCH).Methods("PATCH")
	r.HandleFunc("/api/cmdb/v1.0/instances", hs.InstanceGET).Methods("GET")
	r.HandleFunc("/api/arsys/v1/entry/AST:ComputerSystem", hs.EntryGET).Methods("GET")

	r.HandleFunc("/", hs.Home).Methods("GET")

	// anything we don't recognise we will dump
	r.NotFoundHandler = http.HandlerFunc(hs.Dump)
}
