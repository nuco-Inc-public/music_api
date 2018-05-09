

package main

import (
	"net/http"
	"io/ioutil"
	"os"
	"fmt"
)

func main() {

	http.Handle("/", BaseHandlerFunc(HomeIndex))
	http.Handle("/music_list", BaseHandlerFunc(MusicListIndex))

	err := http.ListenAndServe(":3000", nil)    
	if err != nil {
		fmt.Printf("Failed to connect port %s", err)
	}
}

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("NOT FOUND"))
	} else {
		w.Write([]byte("HOME"))
	}
}

func MusicListIndex(w http.ResponseWriter, r *http.Request) {
	file, e := ioutil.ReadFile("./assets/json/music_list.json")
	if e != nil {
		// fmt.Printf("file error: %v\n", e)
		os.Exit(1)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(file)
}

// http.HandlerFunc(handler)→Cast Index to Hander
func BaseHandlerFunc(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return BaseHandler(http.HandlerFunc(handler))
}

func BaseHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		handler.ServeHTTP(w, r)
	})
}