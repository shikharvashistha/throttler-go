package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/shikharvashistha/throttler-go/pkg/utils"
)

func main() {
	logger := utils.NewLogger("main")

	logger.Info("Attempting to connect to key value store")

	utils.RedisConnect()

	logger.Info("Successfully connected to key value store")

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
