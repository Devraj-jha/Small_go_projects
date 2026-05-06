package main

import (
	"fmt"
	"net/http" // for http thing

)
func main(){
	
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello_you visited this server %s\n", r.URL.Path)
    fmt.Fprintf(w, "this is a backend server\n")
   fmt.Fprintf(w, "this is made in Go 🌟\n")


	})
	fmt.Println("Server is running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)

}
