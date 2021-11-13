package main

import (
	"TarlaIO/router"
	"log"
	"net/http"
)

func main() {
	// Start Server at :8080 port and neglect to router
	log.Fatal(http.ListenAndServe(":8080",router.Routes()))
}
