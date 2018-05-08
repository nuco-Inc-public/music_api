// package main

// import (
//     "net/http"
// )

// func main() {
//     http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
// 			w.Header().Add("Access-Control-Allow-Origin", "*")
// 			//w.Header().Add("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
// 			w.Write([]byte("Hello"))
//     })
//     http.ListenAndServe(":3000", nil)
// }

package main

import (
	"net/http"
	// "github.com/gorilla/mux"
)

func main() {
	
	
	// r := mux.NewRouter()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Write([]byte("HOME"))
	})

	http.HandleFunc("/music_list", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Write([]byte("MUSIC_LIST"))
	})
	// http.Handle("/", r)
  http.ListenAndServe(":3000", nil)
    
}