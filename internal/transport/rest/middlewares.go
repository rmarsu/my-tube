package rest

import (
     "net/http"
)

func CorsMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
          w.Header().Set("Access-Control-Allow-Methods", "*")
          w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")

          if r.Method != "OPTIONS" {
			next.ServeHTTP(w, r)
          } else {
			w.WriteHeader(http.StatusOK)
		}
	})
}

