package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"order_svc/pkg/client"
	"order_svc/pkg/config"
	"order_svc/pkg/db"
	"order_svc/pkg/pb"
	"order_svc/pkg/services"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("failed to listing:", err)
	}

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	if err != nil {
		log.Fatalln("failed to listing:", err)
	}

	fmt.Println("order Svc on", c.Port)

	s := services.Server{
		ProductSvc: productSvc,
		H:          h,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
