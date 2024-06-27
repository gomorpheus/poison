package handlers

import (
	"fmt"
	"github.com/spoonboy-io/koan"
	"io"
	"net/http"
)

type State struct {
	Logger *koan.Logger
	Host   string
	Port   int
}

func (s *State) Dump(w http.ResponseWriter, r *http.Request) {
	s.Logger.Warn("Unhandled Request, dumping")
	s.dumpRequest(r)
	w.WriteHeader(http.StatusOK)
}

func (s *State) Home(w http.ResponseWriter, r *http.Request) {
	s.Logger.Info("home page accessed")
	w.WriteHeader(http.StatusOK)
	out := `<html>
<head></head>
<body>Poison server is running...</body>
</html>
`
	_, _ = fmt.Fprint(w, out)
}

func (s *State) dumpRequest(r *http.Request) {
	header := ""

	for k, v := range r.Header {
		header += fmt.Sprintf("%s:%s\n", k, v)
	}
	fmt.Printf("\nRoute:\n------\n%s\n\n", r.URL)
	fmt.Printf("Method:\n-------\n%s\n", r.Method)
	fmt.Printf("Request Headers:\n----------------\n%s\n", header)

	rb, err := io.ReadAll(r.Body)

	if err == nil {
		fmt.Printf("Request Body:\n-------------\n%s\n", string(rb))
	}
}
