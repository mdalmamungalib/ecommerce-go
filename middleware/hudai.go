package middleware

import (
	"log"
	"net/http"
)

func Hudai(next http.Handler) http.Handler{
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Panicln("ami hudai middleware")
		next.ServeHTTP(w, r)
	})
}