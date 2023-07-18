package quic_server

import (
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

func setupGenericHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%#v\n", r)
		bytes,_ := json.Marshal("{}")
		w.Write(bytes)
	})
	return nil
}