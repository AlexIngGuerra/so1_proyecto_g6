package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	pb "sopes/protoc"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var total int

//bases de datos
var client *mongo.Client

//INFORMACION A INGRESAR
type Info struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Score string `json:"score"`
	Phase string `json:"phase"`
}

/* ***************************
LEER ARCHIVO DE KAFKA
******************************/
func read() {
	fmt.Println("inicando lectura..")
	// to consume messages
	topic := "topic_test"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "34.135.161.214:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	count := 0
	for {
		count += 1
		n, err := batch.Read(b)
		if err != nil {
			break
		}

		if count > total {
			fmt.Print(count)
			fmt.Print(" - ")
			enviarInfo(string(b[:n]))
			total += 1
		}
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

func enviarInfo(info string) {
	var inf Info
	json.Unmarshal([]byte(info), &inf)
	fmt.Println(inf)

	//INGRESAR INFO A MONGO
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://172.17.0.2:27017"))
	ingresarLog(inf)
	obtenerLogs()

	defer client.Disconnect(ctx)
}

/* ***************************
INSERTAR EN MONGO
******************************/

func ingresarLog(info Info) {
	fmt.Println("Enviamos info a mongo")

	colletion := client.Database("goDB").Collection("Log")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := colletion.InsertOne(ctx, info)

	IngresarDatos(info.Team1, info.Team2, info.Score, info.Phase)

	fmt.Println(result)
}

func obtenerLogs() {

	colletion := client.Database("goDB").Collection("Log")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := colletion.Find(ctx, bson.M{})

	var lista []Info
	if err != nil {
		fmt.Println(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var auto Info
		cursor.Decode(&auto)
		lista = append(lista, auto)
	}

	fmt.Println(lista)
}

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "20.120.51.0:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func IngresarDatos(team1 string, team2 string, score string, phase string) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.IngresoDatos(ctx, &pb.IngresoSolicitud{Team1: team1, Team2: team2, Score: score, Phase: phase})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("\nCodigo: %s\nMensaje: %s", r.GetCodigo(), r.GetMensaje())
}

/* ***************************
FUNCION MAIN
******************************/
func main() {
	total = 0
	//CONFIGURAR DB MONGO

	for true {

		fmt.Print("Total: ")
		fmt.Println(total)
		read()
		time.Sleep(time.Second)
	}

}
