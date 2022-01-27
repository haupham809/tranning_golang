package main

import (
	"fmt"
	"strconv"
	"time"
	"tranning_golang/model/connectdb"
	"tranning_golang/model/coupon"
	"tranning_golang/model/messageapi"
	"tranning_golang/model/order"
	"tranning_golang/model/payment"
	"tranning_golang/model/product"
	"tranning_golang/model/topping"
	"tranning_golang/model/writelog"

	"github.com/labstack/echo"
	"gopkg.in/robfig/cron.v2"
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
	c.AddFunc("@every 1h", CronTab)
	c.AddFunc("TZ=Asia/Bangkok 30 04 * * * *", func() { fmt.Println("Runs at 04:30 Bangkok time every day") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 0h0m1s", func() { fmt.Println("Every second") })
	c.Start()

	// Funcs are invoked in their own goroutine, asynchronously.
	go CronTab()
	// Funcs may also be added to a running Cron
	c.AddFunc("@daily", func() { fmt.Println("Every day") })

	// Added time to see output
	time.Sleep(10 * time.Second)

	c.Stop() // Stop the scheduler (does not stop any jobs already running).
}

func CronTab() {
	if connectdb.Connnectdb() {
		var result []User
		tx := connectdb.DB.Begin()
		connectdb.DB.Raw("select * from t_user").Scan(&result)
		for i := 0; i < len(result); i++ {
			var s = strconv.Itoa(result[i].ID)
			if getCountPaymentID(s) > 7 {
				err := tx.Table("t_user").Where("id = ?", result[i].ID).Update("type", 2).Error
				if err != nil {
					tx.Rollback()
					errorapi := messageapi.Objectapi{
						Status:  500,
						Message: "create error",
					}
					writelog.Writelog(errorapi)
				} else {
					tx.Commit()
					successapi := messageapi.Objectapi{
						Status:  200,
						Message: "create success",
					}
					writelog.Writelog(successapi)
				}
			}

		}
	} else {
		errorconnnet := messageapi.Objectapi{
			Status:  500,
			Message: "database disconnect",
		}
		writelog.Writelog(errorconnnet)
	}
}

func Getpayment(id string) User {
	var result []User

	connectdb.DB.Raw("SELECT  * FROM user where id = " + id).Scan(&result)
	return result[0]
}

func getCountPaymentID(id string) int {
	var countn int64
	var result []Payment
	connectdb.DB.Raw("select * from t_payment where user_id = " + id).Scan(&result)
	s := strconv.FormatInt(countn, 10) // s == "97" (decimal))
	fmt.Println("long kute :" + s)
	return len(result)
}
