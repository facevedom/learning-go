/* A basic HTTP server has a few key jobs to take care of.
   - Process dynamic requests: Process incoming requests from users who browse
	 the website, log into their accounts or post images.
   - Serve static assets: Serve JavaScript, CSS and images to browsers to create
	 a dynamic experience for the user.
   - Accept connections: The HTTP server must listen on a specific port to be able
     to accept connections from the internet.	 
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	/* The `http.Request` contains all information about the request and it's 
	   parameters. You can read GET parameters with `r.URL.Query().Get("token")`
	   or POST parameters with `r.FormValue("email")`. */
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	/* To serve static assets like JavaScript, CSS and images we use the builtin
	   `http.FileServer` and point it to a URL path. For the file server to work
	   properly it nees to know where to serve files from. We can do this like so:*/
	fs := http.FileServer(http.Dir("static/"))
	/* Once our file server is in place, we just need to point a URL path at it,
	   just like we did with the dynamic requests. One thing to note: in order to
	   serve files correctly, we need to strip away a part of the URL path. 
	   Usually this is the name of the directory our files live in. */
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}