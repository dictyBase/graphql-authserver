package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func AuthorizeMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		hdr := r.Header
		if hdr.Get("X-Scheme") != "https" {
			http.Error(
				w,
				fmt.Sprintf("scheme is %s not https", hdr.Get("X-Schema")),
				http.StatusBadRequest,
			)
			return
		}
		if hdr.Get("X-Original-Method") == "OPTIONS" || hdr.Get("X-Original-Method") == "GET" {
			_, err := w.Write([]byte("passthrough for non-POST method"))
			if err != nil {
				log.Println("write error")
			}
			return
		}
		if hdr.Get("X-GraphQL-Method") == "Query" {
			_, err := w.Write([]byte("passthrough for GraphQL query"))
			if err != nil {
				log.Println("write error")
			}
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
