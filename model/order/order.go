package order

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"tranning_golang/model/connectdb"
	"tranning_golang/model/messageapi"
	"tranning_golang/model/writelog"

	"github.com/labstack/echo"
)

type T_order struct {
	ID                 int       `json:"id"`
	Product_id         string    `json:"product_id"`
	Product_topping_id string    `json:"product_topping_id"`
	User_id            string    `json:"user_id"`
	Payment_id         string    `json:"payment_id"`
	Quantity           int       `json:"quantity"`
	Amount             float64   `json:"amount"`
	Note               string    `json:"note"`
	Size               int       `json:"size"`
	Shipment_date      time.Time `json:"shipment_date"`
	Status             int       `json:"status"`
	Date_created       time.Time `json:"date_created"`
	Date_update        time.Time `json:"date_update"`
	Is_delete          int       `json:"is_delete"`
}
type T_payment struct {
	ID                int       `json:"id"`
	User_id           string    `json:"user_id"`
	Credit_card_id    string    `json:"credit_card_id"`
	Coupon_id         string    `json:"coupon_id"`
	Total_amount      string    `json:"total_amount"`
	Payment_result_id int       `json:"payment_result_id"`
	Pament_date       time.Time `json:"pament_date"`
	Refund_date       time.Time `json:"refund_date"`
	Is_refund         int       `json:"is_refund"`
	Date_created      time.Time `json:"date_created"`
	Date_update       time.Time `json:"date_update"`
}

//kiểm tra đơn hang tồn tại
func checkorder(id int) int {

	var result []T_order

	connectdb.DB.Raw("SELECT  * FROM t_order where id = " + strconv.Itoa(id)).Scan(&result)

	return len(result)
}

type Jsonbody struct {
	ID int
}

func Order(c echo.Context) error {
	if connectdb.Connnectdb() {
		tx := connectdb.DB.Begin()
		var order []T_order
		json.NewDecoder(c.Request().Body).Decode(&order)
		err := tx.Table("t_order").Create(&order).Error
		if err != nil {
			tx.Rollback()
			errorapi := messageapi.Objectapi{
				Status:  500,
				Message: "create error",
			}
			writelog.Writelog(errorapi)
			return c.JSON(http.StatusCreated, errorapi)
		} else {
			tx.Commit()
			successapi := messageapi.Objectapi{
				Status:  200,
				Message: "create success",
			}
			writelog.Writelog(successapi)
			return c.JSON(http.StatusOK, successapi)
		}

	} else {
		errorconnnet := messageapi.Objectapi{
			Status:  500,
			Message: "database disconnect",
		}
		writelog.Writelog(errorconnnet)
		return c.JSON(http.StatusBadRequest, errorconnnet)
	}
}

// o       Upate thông tin huỷ đơn hàng
func Updateordercancel(c echo.Context) error {
	if connectdb.Connnectdb() {
		tx := connectdb.DB.Begin()

		var id Jsonbody
		json.NewDecoder(c.Request().Body).Decode(&id)

		if checkorder(id.ID) == 0 {
			errorupdate := messageapi.Objectapi{
				Status:  500,
				Message: "order no exist",
			}
			writelog.Writelog(errorupdate)
			return c.JSON(http.StatusBadRequest, errorupdate)
		} else {
			err := tx.Table("t_order").Where(" id = ?", id.ID).Update("status", 0).Error

			if err != nil {
				tx.Rollback()
				errorupdate := messageapi.Objectapi{
					Status:  500,
					Message: "update error",
				}
				writelog.Writelog(errorupdate)
				return c.JSON(http.StatusBadRequest, errorupdate)
			} else {
				tx.Commit()
				successupdate := messageapi.Objectapi{
					Status:  200,
					Message: "update success",
				}
				return c.JSON(http.StatusOK, successupdate)
			}
		}

	} else {
		errorconnnet := messageapi.Objectapi{
			Status:  500,
			Message: "database disconnect",
		}
		writelog.Writelog(errorconnnet)
		return c.JSON(http.StatusBadRequest, errorconnnet)
	}

}
