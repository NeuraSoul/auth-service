package main

import (
	"log"
	"net"

	"auth-service/internal/domain/services"
	"auth-service/internal/repositories"
	pd "auth-service/pkg/grpc"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main () {
	port := ":9000"
	listner ,err := net.Listen("tcp",port)
	if err != nil {
		log.Panic("fail to serve in %v",port)
	}

	s := grpc.NewServer()

	go func () {
		log.Printf("start gRPC server port  %v",port)
		s.Serve(listner)

	}()
	

	rep := repositories.NewRepositories(&gorm.DB{})

	pd.RegisterAuthServiceServer(s,services.NewAuthServer(rep))

}