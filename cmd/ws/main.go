package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handler := corsMiddleware(http.DefaultServeMux)
	webServer := &http.Server{Addr: ":52018"}

	http.Handle("/", http.FileServer(http.Dir("/home/dr/src/hundred-marbels/wwwRoot/")))

	// go func() {
	if err := webServer.ListenAndServe(); err != nil {
		fmt.Printf("err: %s\n", err)
	}
	// }()
}
