package main

import (
	"fmt"
	"log"

	"github.com/alidevjimmy/microgo/product/models"
	"github.com/alidevjimmy/microgo/product/repository"
	"github.com/alidevjimmy/microgo/product/services"
)

func main() {
	productRepo := repository.NewProductRepository()
	productService := services.NewProductService(productRepo)

	p1 := models.Product{
		Title:    "Product 1",
		Price:    10.00,
		Quantity: 4,
	}

	p2 := models.Product{
		Title:    "Product 2",
		Price:    20.00,
		Quantity: 20,
	}

	if err := productService.Create(&p1); err != nil {
		log.Println(err)
	}

	if err := productService.Create(&p2); err != nil {
		log.Println(err)
	}

	products, err := productService.FindAll()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(products)

}
