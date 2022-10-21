package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grcp/proyecto/protoc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) IngresoDatos(ctx context.Context, in *pb.IngresoSolicitud) (*pb.Respuesta, error) {
	log.Printf("Se registra información. \nteam1: %v\nteam2: %v\nscore: %v\nphase: %v", in.Team1, in.Team2, in.Score, in.Phase)
	return &pb.Respuesta{Codigo: "200", Mensaje: "Se registró información"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
