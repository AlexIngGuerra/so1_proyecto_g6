package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/segmentio/kafka-go"
)

var (
	// flagPort is the open port the application listens on
	flagPort = flag.String("port", "9000", "Port to listen on")
)

const (
	topic         = "prueba9"
	brokerAddress = "localhost:9092"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/getinfo", getDataFiltro)

	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := json.Marshal("ok")
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}

func getDataFiltro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	//go produce(ctx)
	var info = consume(ctx)

	fmt.Fprint(w, string(info))

}

func consume(ctx context.Context) string {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		// assign the logger to the reader
		//Logger: l,
	})

	msg, err := r.ReadMessage(ctx)
	if err != nil {
		panic("could not read message " + err.Error())
	}

	//fmt.Println("received: ", string(msg.Value))

	return string(msg.Value)
}
