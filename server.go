package main

import (
	"fmt"
	"net/http"
	"urlShortner/dbmanager"
	pack "urlShortner/urlProcessor"

	"github.com/gorilla/mux"
)

func main() {
	// TestSomeLinks()
	dbmanager.Init()
	router := mux.NewRouter()
	router.HandleFunc("/{path}", handelShorty)
	router.HandleFunc("/", greet).Methods("GET")
	fmt.Println("Server is running on port http://127.0.0.1:8080")
	http.ListenAndServe(":8080", router)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Paste your URL here</h1>"))
}
func handelShorty(writer http.ResponseWriter, reader *http.Request) {
	params := mux.Vars(reader) // placeholder {path in this case} will be the key
	if len(params) > 0 {
		url := pack.RetriveURL(params["path"])
		if url == "" {
			writer.WriteHeader(404)
			writer.Write([]byte("Page not foud"))
		} else {
			fmt.Println(url)
			http.Redirect(writer, reader, url, http.StatusPermanentRedirect)
		}
	}
}
