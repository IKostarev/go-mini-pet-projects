package middleware

import (
	"log"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler work")
	w.WriteHeader(http.StatusOK)
}

func myMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware work")
		next(w, r)
	}
}

func main() {
	http.HandleFunc("/", myHandler)

	http.ListenAndServe(":8080", myMiddleware(myHandler))
}
