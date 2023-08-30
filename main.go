
package main

import( 
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
    "github.com/denisenkom/go-mssqldb"
)

func main(){
	dbString := "sqlserver://ADANOS\\aalsa@Adanos\\SQLEXPRESS?database=Go_auth"
	_, err := gorm.Open("mssql", dbString)
	if err != nil {
        panic("could not connect to database")
	}


	app := fiber.New()
	app.Get("/", func(c*fiber.Ctx) error{
		return c.SendString("the world")
	})

	app.Listen("8000")
}