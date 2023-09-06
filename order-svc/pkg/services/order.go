package services

import (
	"context"
	"net/http"
	"order_svc/pkg/client"
	"order_svc/pkg/db"
	"order_svc/pkg/models"
	"order_svc/pkg/pb"
)

type Server struct {
	ProductSvc client.ProductServiceClient
	pb.UnimplementedOrderServiceServer
	H db.Handler
}

func (s Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	product, err := s.ProductSvc.FindOne(ctx, req.ProductId)
	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusNotFound, Error: err.Error()}, nil
	} else if product.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{Status: product.Status, Error: product.Error}, nil
	} else if product.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: product.Error}, nil
	}

	order := models.Order{
		Price:     product.Data.Price,
		ProductId: req.ProductId,
		UserId:    req.UserId,
		Quantity:  req.Quantity,
	}

	s.H.DB.Create(&order)

	res, err := s.ProductSvc.DecreaseStock(ctx, req.ProductId, order.Id)
	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if res.Status == http.StatusConflict {
		s.H.DB.Delete(&models.Order{}, order.Id)

		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: res.Error}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
