package main

import (
	"fmt"
	"net/http"
	"time"
	"tranning_golang/model/connectdb"
	"tranning_golang/model/coupon"
	"tranning_golang/model/order"
	"tranning_golang/model/payment"
	"tranning_golang/model/product"
	"tranning_golang/model/topping"

	"github.com/labstack/echo"
	"gopkg.in/robfig/cron.v2"
	"gorm.io/gorm"
)

type User struct {
	ID           int
	Name         string `valid:"required"`
	Email        string `valid:"required"`
	Password     string `valid:"required"`
	Address      string `valid:"required"`
	Phone        string `valid:"required"`
	Type         int    `valid:"required"`
	Date_created time.Time
	Date_update  time.Time
}
type Payment struct {
	ID                int
	UserID            int     `valid:"required"`
	Credit_card_id    int     `valid:"required"`
	Coupon_id         int     `valid:"required"`
	Total_amount      float64 `valid:"required"`
	Payment_result_id string  `valid:"required"`
	Payment_date      time.Time
	Refund_date       time.Time
	Is_refund         int `valid:"required"`
	Date_created      time.Time
	Date_update       time.Time
}

func main() {
	e := echo.New()
	e.GET("/toppingforproduct", topping.ToppingForProduct)
	e.GET("/coupon", coupon.GetCoupon)
	e.GET("/detailproduct", product.DetailProduct)
	e.GET("/sizeproduct", product.SizeProduct)
	e.GET("/cancelorder", order.Updateordercancel)
	e.POST("/savepayment", payment.SavePayment)
	e.POST("/savecreditcard", payment.SaveCreditCard)
	e.POST("/updatepaymentcancel", payment.Updatepaymentcancel)
	e.POST("/refundmomo", payment.Refundmomo)
	// e.PUT("/updateorder", order.UpdateOrder)
	e.POST("/order", order.Order)
	e.GET("/products", product.GetProduct)
	e.GET("/couponbyid", coupon.GetCouponByUserLogin)
	e.GET("/getsuggest", product.GetProductSuggest)
	e.Logger.Fatal(e.Start(":1323"))

	c := cron.New()
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("TZ=Asia/Bangkok 30 04 * * * *", func() { fmt.Println("Runs at 04:30 Bangkok time every day") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 0h0m1s", func() { fmt.Println("Every second") })
	c.Start()

	// Funcs are invoked in their own goroutine, asynchronously.

	// Funcs may also be added to a running Cron
	c.AddFunc("@daily", func() { fmt.Println("Every day") })

	// Added time to see output
	time.Sleep(10 * time.Second)

	c.Stop() // Stop the scheduler (does not stop any jobs already running).
}

// func Cron() error {
// 	if connectdb.Connnectdb() {
// 		tx := connectdb.DB.Begin()

// 	} else {
// 		errorconnnet := messageapi.Objectapi{
// 			Status:  500,
// 			Message: "database disconnect",
// 		}
// 		writelog.Writelog(errorconnnet)
// 	}
// }

func getUserOr404(db *gorm.DB, id string) (*User, *echo.HTTPError) {
	s := &User{}
	if err := db.First(&s, id).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return s, nil
}

func getCountPaymentID(id string) {
	var count int64
	var result []Payment
	connectdb.DB.Raw("select * from t_size_product where is_delete = 0 and product_id = " + id).Scan(&result).Count(&count)
}
