package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func loginLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Login!")
}
func registerLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Register!")
}

func main() {
	fmt.Printf("hello, లక్ష్మణ \n")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/requestLogin", loginLink)
	router.HandleFunc("/api/registerMe", registerLink)

	s := &http.Server{
		Addr:           ":4500",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Server failed: ", err.Error())
	} else {
		fmt.Printf("Server Started")
	}
	// http.ListenAndServe(":6000", nil)

}