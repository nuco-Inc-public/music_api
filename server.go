

package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Write([]byte("HOME"))
	})

	http.HandleFunc("/music_list", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Write([]byte("MUSIC_LIST"))
	})
  http.ListenAndServe(":3000", nil)
    
}