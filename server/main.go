package main

import (
	"log"
	"net"
	service "omdb-server/service"
	"os"

	pb "omdb-server/pb-file"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		log.Println("using default environment variable")
	}

	port := os.Getenv("API_PORT")

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln("Failed to listen")
	}

	s := service.Server{}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pb.RegisterOMDBServiceServer(grpcServer, &s)

	log.Print("gRPC server started at ", port)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal("Failed to serve")
	}

}
