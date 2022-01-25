package main

import (
	"tranning_golang/model/coupon"
	"tranning_golang/model/order"
	"tranning_golang/model/product"
	"tranning_golang/model/topping"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/toppingforproduct", topping.ToppingForProduct)
	e.GET("/coupon", coupon.GetCoupon)
	e.GET("/detailproduct", product.DetailProduct)
	e.GET("/sizeproduct", product.SizeProduct)
	e.GET("/cancelorder", order.Updateordercancel)
	e.POST("/order", order.Order)
	e.GET("/products", product.GetProduct)
	e.GET("/couponbyid", coupon.GetCouponByUserLogin)
	e.GET("/getsuggest", product.GetProductSuggest)
	e.Logger.Fatal(e.Start(":1323"))
}
