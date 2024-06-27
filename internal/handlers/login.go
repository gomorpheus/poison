package handlers

import (
	"fmt"
	"net/http"
)

func (s *State) Login(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("login attempted")
	s.dumpRequest(r)

	// don't care about credentials passed, just issue jwt, and status 200
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "dummyjwt")
}
