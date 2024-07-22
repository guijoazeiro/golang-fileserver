package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: <http-dir> <port>")
		os.Exit(1)
	}

	httpDir := os.Args[1]
	port := os.Args[2]
	fs := http.FileServer(http.Dir(httpDir))
	fmt.Printf("Listening on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, fs))
}
