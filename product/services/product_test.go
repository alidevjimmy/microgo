package services

import (
	"testing"

	"github.com/alidevjimmy/microgo/product/models"
	"github.com/alidevjimmy/microgo/product/repository"
)

func TestProductMustBeGreaterThanFive(t *testing.T) {
	fakeRepo := repository.NewProductRepository()
	productService := NewProductService(fakeRepo)

	product := models.Product{
		Title:    "Fake Product 1",
		Price:    10.00,
		Quantity: 4,
	}

	err := productService.Create(&product)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
