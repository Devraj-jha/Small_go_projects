package main

import (
	"fmt"
	"net/http" // for http thing

)
func main(){
	
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello_you visted this server %s", r.URL.Path)

	})
	fmt.Println("Server is running at http://localhost:8080")
    http.ListenAndServe(":999", nil)

}
