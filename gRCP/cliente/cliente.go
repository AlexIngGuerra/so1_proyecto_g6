package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grcp/proyecto/protoc"
	"log"
	"time"
)

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

func main() {
	IngresarDatos("Guatemala", "Belgica", "1-2", "1")
	IngresarDatos("Alemania", "Argentina", "2-3", "2")
	IngresarDatos("Brasil", "Portugal", "5-3", "3")
	IngresarDatos("Inglaterra", "EstadosUnidos", "1-7", "4")
}
