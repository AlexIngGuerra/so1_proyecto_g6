package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

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

type Informacion struct {
	team1 string
	team2 string
	score string
	phase string
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/getinfo", guardaInfo)
	mux.HandleFunc("/guardainfo", guardaInfo)

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

func guardaInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	} else {
		cuerpo := string(body)

		if (len(cuerpo)) > 0 {

			ctx := context.Background()
			go produce(ctx, cuerpo)
			log.Printf("data: %s", (cuerpo))
			fmt.Fprint(w, cuerpo)
		}

	}

	log.Printf("data: %s", "ingreso en data nueva")

}

func produce(ctx context.Context, data string) {
	// initialize a counter
	i := 0

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	//for {
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa(i)),
		// create an arbitrary message payload for the value
		Value: []byte(data),
	})
	if err != nil {
		panic("no guardo mensaje" + err.Error())
	}

	// log a confirmation once the message is written
	fmt.Println("writes:", i)
	i++
	// sleep for a second
	time.Sleep(time.Second)
	//}
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
