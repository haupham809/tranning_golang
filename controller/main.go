package main
import (
	"github.com/labstack/echo"
	"tranning_golang/model/product"
	
)

func main() {
	e := echo.New()
	e.GET("/check", product.DetailProduct)

	e.Logger.Fatal(e.Start(":1323"))
}

