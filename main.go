package main

import (
	"encoding/json"
	"github.com/rodrigoavalente/gopokemon"
	"net/http"
	"strings"
)

func Hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("<h1>Hello, world!</h1>"))
}

func main() {
	http.HandleFunc("/pokemon/", func(writer http.ResponseWriter, request *http.Request) {
		pokemon_id := strings.SplitN(request.URL.Path, "/", 3)[2]
		pokemon, err := gopokemon.QueryPokemon(pokemon_id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(writer).Encode(pokemon)
	})
	http.ListenAndServe(":8080", nil)
}
