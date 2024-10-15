package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Mariscal6/be-fe-template/hello-brambles-be/dbclient/rickmortyapi"
)

func main() {
	mux := http.NewServeMux()
	dbClient := rickmortyapi.RickMortyAPI{}

	mux.HandleFunc("/characters", func(w http.ResponseWriter, r *http.Request) {
		pageParam := r.URL.Query().Get("page")
		page := 0
		if pageParam != "" {
			i, err := strconv.Atoi(pageParam)
			if err == nil {
				page = i
			}
		}
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		characters, err := dbClient.GetCharacters(r.Context(), page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(characters)
	})
	fmt.Println("Server is running on port 8081")
	http.ListenAndServe(":8081", mux)
}
