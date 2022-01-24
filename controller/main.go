package main
import (
	"github.com/labstack/echo"
	"tranning_golang/model/product"
	"tranning_golang/model/topping"
	"tranning_golang/model/coupon"
	"tranning_golang/model/order"
	
)

func main() {
	e := echo.New()
	e.GET("/toppingforproduct", topping.ToppingForProduct)
	e.GET("/coupon",coupon.GetCoupon)
	e.GET("/detailproduct", product.DetailProduct)
	e.GET("/sizeproduct", product.SizeProduct)
	e.GET("/cancelorder", order.Updateordercancel)
	e.Logger.Fatal(e.Start(":1323"))
}

