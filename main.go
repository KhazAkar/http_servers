package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Redirecting to /app/")
	fmt.Println(r.URL.Path)
	fmt.Println("Location Header:", w.Header().Get("Location"))
	http.Redirect(w, r, "/app/", http.StatusMovedPermanently)
	fmt.Println("Location Header After Redirect:", w.Header().Get("Location"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/app", redirectHandler)
	handler := http.FileServer(http.Dir("./pages"))
	handler = http.StripPrefix("/app", handler)
	mux.Handle("/app/", handler)
	mux.HandleFunc("/healthz", healthHandler)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
