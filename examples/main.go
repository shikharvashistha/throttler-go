package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/shikharvashistha/throttler-go/pkg/middleware"
	keyvalue "github.com/shikharvashistha/throttler-go/pkg/store"
	"github.com/shikharvashistha/throttler-go/pkg/utils"
)

func main() {

	godotenv.Load()

	utils.RedisConnect(
		os.Getenv("REDIS_ADDR"),
		os.Getenv("REDIS_PASSWORD"),
		"default",
		0,
	)

	kvs := keyvalue.NewKVStore()

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		a := middleware.GetAnonymousThrottle(10, time.Minute, "test", 5, kvs)
		c, _ := a.AllowRequest(r)
		wait, _ := a.Wait()
		fmt.Fprintln(w, "Limit exceeded: ", !c)
		fmt.Fprintln(w, "Recomened Wait: ", wait)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	log.Printf("Listenning on localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe failed: " + err.Error())
	}

}
