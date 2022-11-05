package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

type Info struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Score string `json:"score"`
	Phase string `json:"phase"`
}

func Hola(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("Hola Mundo")
}

func write(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var info Info
	json.NewDecoder((request.Body)).Decode(&info)

	topic := "topic_test"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "34.135.161.214:9092", topic, partition)
	if err != nil {
		json.NewEncoder(response).Encode("failed to dial leader")
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	_, err = conn.WriteMessages(kafka.Message{Value: []byte("{" +
		"\"team1\":\"" + info.Team1 + "\"," +
		"\"team2\":\"" + info.Team2 + "\"," +
		"\"score\":\"" + info.Score + "\"," +
		"\"phase\":\"" + info.Phase + "\"" +
		"}\n")})

	if err != nil {
		json.NewEncoder(response).Encode("failed to write messages")
	}

	if err := conn.Close(); err != nil {
		json.NewEncoder(response).Encode("failed to close writer")
	}

	json.NewEncoder(response).Encode("Enviado")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/input", write).Methods("POST")
	router.HandleFunc("/", Hola).Methods("GET")

	fmt.Println("Server on port", 5010)
	err := http.ListenAndServe(":5010", router)
	if err != nil {
		fmt.Println(err)
	}
}
