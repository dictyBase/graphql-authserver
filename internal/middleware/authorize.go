package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Params struct {
	OperationName string                 `json:"operationName"`
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
}

func AuthorizeMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		allowedMutations := []string{"CreateOrder", "Login"}
		var params Params
		// decode request body into struct, return 400 status code if error
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// let slice of allowed mutations pass through
		for _, m := range allowedMutations {
			if strings.Contains(params.OperationName, m) {
				fmt.Printf("got %s mutation, no token necessary \n", params.OperationName)
				w.Write([]byte("passthrough for allowed mutation"))
				return
			}
		}
		// verify request is not a mutation
		if !strings.Contains(params.Query, "mutation") {
			w.Write([]byte("passthrough for non-mutation request"))
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
