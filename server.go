package main

import (
	"fmt"
	"log"
	"net"
	"proyek/controllers"
	"proyek/pb/cities"
	"proyek/pkg/database"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	db, err := database.OpenDatabase()
	if err != nil {
		log.Fatalf("error: connecting to db: %s", err)
	}
	defer db.Close()

	grpcServer := grpc.NewServer()

	cityServer := controllers.City{DB: db}
	cities.RegisterCitiesServiceServer(grpcServer, &cityServer)

	fmt.Println("running server grpc")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %s\n", err)
		return
	}
}
