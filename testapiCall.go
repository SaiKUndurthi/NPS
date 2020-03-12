package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NPS/apiKeys"
)

func index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, apiKeys.API_KEY)
	fmt.Fprintf(w, "<h1>Hello World</h1>\n")
	response, err := http.Get("https://developer.nps.gov/api/v1/parks?parkCode=acad&api_key=" + apiKeys.API_KEY)
	if err != nil {
		fmt.Fprintf(w, "The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Fprintf(w, string(data))
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
