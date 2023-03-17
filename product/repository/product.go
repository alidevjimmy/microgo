package repository

import (
	"github.com/alidevjimmy/microgo/product/models"
)

type ProductRepository interface {
	Create(product *models.Product) error
	FindAll() ([]models.Product, error)
}

type productRepository struct {
	Products []models.Product
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		Products: []models.Product{},
	}
}

func (p *productRepository) Create(product *models.Product) error {
	p.Products = append(p.Products, *product)
	return nil
}

func (p *productRepository) FindAll() ([]models.Product, error) {
	return p.Products, nil
}
