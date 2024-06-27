package handlers

import (
	"fmt"
	"net/http"
)

func (s *State) People(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("people attempted")
	s.dumpRequest(r)

	out := `{
	"entries": [
		{
			"values":{
				"Request ID": "p01",
				"Full Name": "Poison User",
				"Group List": "Poison Group"
			}
		}	
	]
}`

	fmt.Printf("results out: %+v", string(out))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, out)
}
