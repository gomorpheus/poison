package handlers

import (
	"fmt"
	"net/http"
)

func (s *State) EntryGET(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("entry (GET) attempted")
	s.dumpRequest(r)

	out := `{
	"entries": [
		{
			"values":{
				"Request ID": "p02"
			}
		}	
	]
}`

	fmt.Printf("results out: %+v", string(out))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, out)
}
