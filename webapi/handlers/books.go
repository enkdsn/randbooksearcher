package handlers

import (
	"fmt"
	"net/http"
)

// HTTP handler accessing data from the request context.
func MyRequestHandler(w http.ResponseWriter, r *http.Request) {
	// here we read from the request context and fetch out `"user"` key set in
	// the MyMiddleware example above.
	user := r.Context().Value("user").(string)

	// respond to the client
	w.Write([]byte(fmt.Sprintf("hi %s", user)))
}
