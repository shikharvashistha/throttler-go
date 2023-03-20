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

	anonymous_throttle := middleware.GetAnonymousThrottle(10, time.Minute, "test_anonymuous", 5, kvs)
	http.HandleFunc("/testAnonymuous", func(w http.ResponseWriter, r *http.Request) {

		c, _ := anonymous_throttle.AllowRequest(r)
		wait, _ := anonymous_throttle.Wait()
		fmt.Fprintln(w, "Limit exceeded: ", !c)
		fmt.Fprintln(w, "Recomened Wait: ", wait)
	})

	custom_throttle := middleware.GetCustomThrottle(10, time.Minute, "test_custom", kvs, func(r *http.Request, scope string) (string, error) {
		return r.Header.Get("user_id") + scope, nil
	})
	http.HandleFunc("/testCustom", func(w http.ResponseWriter, r *http.Request) {
		c, _ := custom_throttle.AllowRequest(r)
		wait, _ := custom_throttle.Wait()
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
