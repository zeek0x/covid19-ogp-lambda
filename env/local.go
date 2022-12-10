//go:build !release

package env

import (
	"io"
	"log"
	"net/http"
)

func Main(handler func() (string, error)) {
	h := func(w http.ResponseWriter, _ *http.Request) {
		s, _ := handler()
		io.WriteString(w, s)
	}
	http.HandleFunc("/", h)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
