package services

import (
	"context"
	"net/http"
	"product-svc/pkg/db"
	"product-svc/pkg/models"
	"product-svc/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedProductServiceServer
}

func (s Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := models.Product{
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	}

	if result := s.H.DB.WithContext(ctx).Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{Status: http.StatusConflict, Error: result.Error.Error()}, nil
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.ID,
	}, nil
}

func (s Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	product := models.Product{}
	if result := s.H.DB.WithContext(ctx).First(&product, req.Id); result.Error != nil {
		return &pb.FindOneResponse{Status: http.StatusNotFound, Error: result.Error.Error()}, nil
	}

	data := &pb.FindOneData{
		Id:    product.ID,
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product

	if result := s.H.DB.WithContext(ctx).First(&product, req.Id); result.Error != nil {
		return &pb.DecreaseStockResponse{Status: http.StatusNotFound, Error: result.Error.Error()}, nil
	}

	if product.Stock <= 0 {
		return &pb.DecreaseStockResponse{Status: http.StatusConflict, Error: "stock too low"}, nil
	}

	var log models.StockDecreaseLog
	if result := s.H.DB.WithContext(ctx).Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&log); result.Error == nil {
		return &pb.DecreaseStockResponse{Status: http.StatusConflict, Error: "Stock already decreased"}, nil
	}

	product.Stock -= 1
	s.H.DB.Save(&product)

	log.OrderId = req.OrderId
	log.ProductRefer = product.ID

	s.H.DB.Create(&log)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}
