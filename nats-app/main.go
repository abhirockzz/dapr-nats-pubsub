package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var port string
var natsSubject string
var host string

func init() {
	port = os.Getenv("APP_PORT")
	if port == "" {
		log.Fatalf("missing env var %s", "APP_PORT")
	}

	natsSubject = os.Getenv("NATS_SUBJECT")
	if natsSubject == "" {
		log.Fatalf("missing env var %s", "NATS_SUBJECT")
	}

	host = os.Getenv("HOSTNAME")
}

func main() {
	http.HandleFunc("/dapr/subscribe", func(w http.ResponseWriter, r *http.Request) {
		response := `["` + natsSubject + `"]`
		fmt.Println("subscribed to NATS subject", natsSubject)

		w.Write([]byte(response))
	})

	http.HandleFunc("/"+natsSubject, func(w http.ResponseWriter, r *http.Request) {
		respBytes, _ := ioutil.ReadAll(r.Body)
		fmt.Println("Recieved Message from NATS - " + string(respBytes) + "\nHost: " + host)
	})

	fmt.Println("starting HTTP server....")
	http.ListenAndServe(":"+port, nil)
}
