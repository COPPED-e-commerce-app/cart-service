package rest

import (
	"github.com/COPPED/cart-service/internal/cart/service"
	"github.com/gofiber/fiber/v2"
)

func FetchCart(c *fiber.Ctx) error {
	id := c.Params("id")
	cart := service.FetchCart(id)
	return c.JSON(cart)
}

func FetchCartList(c *fiber.Ctx) error {
	cartList := service.FetchCartList()
	return c.JSON(cartList)
}
