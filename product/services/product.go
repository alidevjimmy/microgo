package services

import (
	"errors"

	"github.com/alidevjimmy/microgo/product/models"
	"github.com/alidevjimmy/microgo/product/repository"
) 	

type ProductService interface {
	Create(product *models.Product) error
	FindAll() ([]models.Product, error)
}

type productService struct {
	ProductRepo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		ProductRepo: repo,
	}
}

func (p *productService) Create(product *models.Product) error {
	if product.Quantity < 5 {
		return errors.New("quantity must be greater than 5")
	}
	return p.ProductRepo.Create(product)
}

func (p *productService) FindAll() ([]models.Product, error) {
	return p.ProductRepo.FindAll()
}
