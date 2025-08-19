package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("[LOG] %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		fmt.Printf("[DONE] in %v\n", time.Since(start))
	})
}
