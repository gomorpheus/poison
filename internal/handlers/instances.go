package handlers

import (
	"fmt"
	"net/http"
)

func (s *State) InstancePOST(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("instance (POST) attempted")
	s.dumpRequest(r)

	out := `{
	"instance_ids": [ 
		"poisonInstance"
	]
}`

	fmt.Printf("results out: %+v", string(out))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, out)

}

func (s *State) InstancePATCH(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("instance (PATCH) attempted")
	s.dumpRequest(r)

	out := `{
	"entries": [
		{
			"values":{
				"Group ID": "Poison Group"
			}
		}	
	]
}`

	fmt.Printf("results out: %+v", string(out))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, out)

}

func (s *State) InstanceGET(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("instance (GET) attempted")
	s.dumpRequest(r)

	out := `{
	"entries": [
		{
			"values":{
				"Group ID": "Poison Group"
			}
		}	
	]
}`

	fmt.Printf("results out: %+v", string(out))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, out)

}