/* Go is a battery included programming language and has a webserver already
   built in. The `net/http` package from the standard library contains all
   functionalities about the HTTP protocol. This includes (among many other things)
   an HTTP client and an HTTP server. In this example you will figure out how
   simple it is, to create a webserver that you can view in your browser. */
package main

import (
	"fmt"
	"log"
	"net/http"
)

/* First, create a Handler which receives all incomming HTTP connections from
   browsers, HTTP clients or API requests. 
   
   The function receives two parameters:
   - An `http.ResponseWriter` which is where you write your text/html response to.
   - An `http.Request` which contains all information about this HTTP request
	 including things like the URL or header fields. */
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler) // Register the request handler

	/* The request handler alone can not accept any HTTP connections from the 
	outside. An HTTP server has to listen on a port to pass connections on to the 
	request handler. */
	log.Fatal(http.ListenAndServe(":8080", nil))
}
