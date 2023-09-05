package product

import (
	"api-gateway/pkg/config"
	"api-gateway/pkg/product/pb"
	"fmt"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil
	}

	return pb.NewProductServiceClient(cc)
}
