package order

import (
	"api-gateway/pkg/config"
	"api-gateway/pkg/order/pb"
	"fmt"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil
	}

	return pb.NewOrderServiceClient(cc)
}
