package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"order_svc/pkg/pb"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("could't connect: %w\n", err)
	}

	return ProductServiceClient{Client: pb.NewProductServiceClient(cc)}
}

func (c ProductServiceClient) FindOne(ctx context.Context, productID int64) (*pb.FindOneResponse, error) {
	return c.Client.FindOne(ctx, &pb.FindOneRequest{Id: productID})
}

func (c ProductServiceClient) DecreaseStock(ctx context.Context, productID, orderID int64) (*pb.DecreaseStockResponse, error) {
	return c.Client.DecreaseStock(ctx, &pb.DecreaseStockRequest{
		Id:      productID,
		OrderId: orderID,
	})
}
