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
	e.PUT("/updateorder", order.UpdateOrder)
	e.POST("/order", order.Order)
	e.GET("/products", product.GetProduct)
	e.GET("/couponbyid", coupon.GetCouponByUserLogin)
	e.GET("/getsuggest", product.GetProductSuggest)

	e.Logger.Fatal(e.Start(":1323"))
	// c := cron.New()

	// c.AddFunc("@every 0h0m5s", CronTab)

	// c.Start()
}

// kh moi type=2
// kh tn type =1
// khách hàng type = 0 là bình thường
func CronTab() {
	if connectdb.Connnectdb() {
		var result []User
		tx := connectdb.DB.Begin()
		connectdb.DB.Raw("select * from t_user").Scan(&result)
		fmt.Println(len(result))
		for i := 0; i < len(result); i++ {
			var s = strconv.Itoa(result[i].ID)
			if getCountdate(s) <= 7 {
				err := tx.Table("t_user").Where("id = ?", result[i].ID).Update("type", 2).Error
				if err != nil {
					tx.Rollback()
					errorapi := messageapi.Objectapi{
						Status:  500,
						Message: "update error",
					}
					writelog.Writelog(errorapi)
				} else {
					tx.Commit()
					successapi := messageapi.Objectapi{
						Status:  200,
						Message: "update success",
					}
					writelog.Writelog(successapi)
				}
			} else if getCountorder(s) > 10 {
				err := tx.Table("t_user").Where("id = ?", result[i].ID).Update("type", 1).Error
				if err != nil {
					tx.Rollback()
					errorapi := messageapi.Objectapi{
						Status:  500,
						Message: "update error",
					}
					writelog.Writelog(errorapi)
				} else {
					tx.Commit()
					successapi := messageapi.Objectapi{
						Status:  200,
						Message: "update success",
					}
					writelog.Writelog(successapi)
				}

			} else {
				err := tx.Table("t_user").Where("id = ?", result[i].ID).Update("type", 0).Error
				if err != nil {
					tx.Rollback()
					errorapi := messageapi.Objectapi{
						Status:  500,
						Message: "update error",
					}
					writelog.Writelog(errorapi)
				} else {
					tx.Commit()
					successapi := messageapi.Objectapi{
						Status:  200,
						Message: "update success",
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

func getCountorder(id string) int {
	var result []order.T_order
	connectdb.DB.Select("*").Table("t_order").Where("user_id = ?", id).Scan(&result)

	return len(result)
}

type GetDateuser struct {
	Date int
}

func getCountdate(id string) int {
	var result GetDateuser
	connectdb.DB.Select("DATEDIFF(CURRENT_TIMESTAMP ,date_created) as date").Table("t_user").Where("id = ?", id).Scan(&result)
	fmt.Println(result)
	return result.Date
}
