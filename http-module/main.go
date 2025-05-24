// executable point
package main

import (
	"fmt"
	"net/http" //provide http ser
)

// This is a function you’re defining.
// hello is a handler function, which responds to HTTP requests.
// w http.ResponseWriter: this is how you write your response.
// req *http.Request: this holds details of the incoming request (like URL, headers, etc.).
func hello(a http.ResponseWriter, req *http.Request){
	fmt.Fprint(a, "hello wassup from go")
}

func main(){
	// registers the route handler 
	http.HandleFunc("/hello", hello)

	// starts the http server on the port and 
	// nill means use go default Http request multiplexer(router)
	http.ListenAndServe(":3000", nil)
}

// flow :
//  main -> register the route  -> starts listening on the server