package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "admin" {
		return "$1$GIFBxVQH$hNhX59hRo7cwcXUFDAw8/0"
	}

	return "secret"
}
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: <http-dir> <port>")
		os.Exit(1)
	}

	httpDir := os.Args[1]
	port := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("fileserver.com", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))
	fmt.Printf("Listening on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
