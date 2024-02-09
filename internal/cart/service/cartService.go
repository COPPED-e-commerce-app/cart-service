package service

import "github.com/COPPED/cart-service/pkg/models"

func FetchCart(id string) *models.Cart {
	return &models.Cart{}
}

func FetchCartList() *[]models.Cart {
	return &[]models.Cart{
		models.Cart{},
	}
}
