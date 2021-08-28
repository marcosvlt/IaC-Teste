package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		points := rand.Intn(40)
		var m string
		if points > 1 {
			m = "points"
		} else {
			m = "point"
		}
		fmt.Fprintf(w, "Hey! You win "+strconv.Itoa(points)+" "+m)
	})
// alterei para 0.0.0.0 se não o container só aceita conexão dele mesmo se 
// utilizar 127.0.0.1
	http.ListenAndServe("0.0.0.0:8080", nil)
}