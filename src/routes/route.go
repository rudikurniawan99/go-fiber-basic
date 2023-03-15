package route

import (
	"github.com/gofiber/fiber/v2"
	product_route "github.com/rudikurniawan99/go-fiber-basic/src/routes/products"
	user_route "github.com/rudikurniawan99/go-fiber-basic/src/routes/user"
)

func Init(e *fiber.App) {
	users := e.Group("users")
	user_route.Mount(users)
	product := e.Group("product")
	product_route.Mount(product)
}
