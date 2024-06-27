package handlers

import (
	"fmt"
	"net/http"
)

func (s *State) Group(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("group attempted")
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
