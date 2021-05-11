package main

import (
	"fmt"
	"net/http"
)

//like (req, res) in express js
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World this is Web")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}

//ctrl + F5 will start the server
func main() {
	//like router
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server starting")
	//like app.listen(PORT, () => null)
	http.ListenAndServe(":3000", nil)
}
