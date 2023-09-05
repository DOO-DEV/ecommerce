package auth

import (
	"api-gateway/pkg/auth/pb"
	"api-gateway/pkg/config"
	"google.golang.org/grpc"
	"log"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// TODO - add retry function
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		// TODO - add logger
		log.Println("cant connect to auth grpc server", err)
		return nil
	}

	return pb.NewAuthServiceClient(cc)
}
