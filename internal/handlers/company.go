package handlers

import (
	"fmt"
	"net/http"
)

func (s *State) Company(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("company attempted")
	s.dumpRequest(r)
	
	out := `{
	"entries": [
		{
			"values":{
				"Company":"Poison Company"
			}
		}	
	]
}`

	fmt.Printf("results out: %+v", string(out))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, out)
}
