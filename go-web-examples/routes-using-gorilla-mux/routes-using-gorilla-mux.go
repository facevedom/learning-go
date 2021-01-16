/* Go's `net/http` package provides a lot of functionalities for the HTTP 
   protocol. One thing it doesn't do very well is complex request routing like
   segmenting a request URL into single parameters. Fortunately there is a very
   popular package for this, which is well known for the good code quality in
   the Go community. In this example you will see how to use the `gorilla/mux`
   package to create routes with named parameters, GET/POST handlers and domain
   restrictions. 
   
   `Gorilla/mux` is a package which adapts to Go's default HTTP router. It comes
   with a lot of features to increase the productivity when writing web 
   applications. It is also compliant to Go's default request handler signature
   `func (w http.ResponseWriter, r *http.Request)`, so the package can be mixed
   and matched with other HTTP libraries like middleware or existing applications.
   Use the `go get` command to install the package from GitHub like so:

go get -u github.com/gorilla/mux */
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	/* First create a new request router. The router is the main router for your
	   web application and will later be passed as parameter to the server. It 
	   will receive all HTTP connections and pass it on to the request handlers
	   you will register on it */
	r := mux.NewRouter()

	/* Register handlers as usual. Instead of calling `http.HandleFunc(...)`
	   you call `r.HandleFunc(...)`.

	   The biggest strenght of the `gorilla/mux` router is the ability to extract
	   segments from the request URL. In this case, two segments of the URL are 
	   dynamic: The book title and the page.*/
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		/* The last thing is to get the data from these segments. The package comes
		with the function `mux.Vars(r)` which takes the `http.Requests` as 
		parameter and returns a map of the segments. */
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	/* In previous examples, the second parameter of ListenAndServe was `nil`.
	   It is `r` now. It is the main router of the HTTP server. `nil` means to 
	   use the default router of the `net/http` package. */
	http.ListenAndServe(":8080", r)
}

/* Features of `gorilla/mux` 
   -Methods: Restrict the request handler to specific HTTP methods.
      r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
      r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
      r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
      r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
   -Hostnames & subdomains: Restrict the request handler to specific hostnames or
						   subdomains.
	  r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")
   -Schemes: Restrict the request handler to http/https.
	  r.HandleFunc("/secure", SecureHandler).Schemes("https")
	  r.HandleFunc("/insecure", InsecureHandler).Schemes("http")
   -Path prefixes & Subroutes: Restrict the request handler to specific path
							   prefixes
	  bookrouter := r.PathPrefix("/books").Subrouter()
	  bookrouter.HandleFunc("/", AllBooks)
	  bookrouter.HandleFunc("/{title}", GetBook)
*/
