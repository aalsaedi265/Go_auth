
package main

import( 
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	dsn :=  "server=localhost;database=Go_auth;integrated security=SSPI;"
	_, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error:", err)
		panic("could not connect to the database")
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	app.Listen(":8000")
}