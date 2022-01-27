package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"tranning_golang/model/connectdb"
	"tranning_golang/model/messageapi"
	"tranning_golang/model/writelog"

	"github.com/asaskevich/govalidator"
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
type UpdateOrder_Model struct {
	ID         []int `valid:"required"`
	Payment_id int   `valid:"required"`
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

func ConvertStrtoInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}
func UpdateOrder(c echo.Context) error {
	if connectdb.Connnectdb() {
		tx := connectdb.DB.Begin()
		var param UpdateOrder_Model
		json.NewDecoder(c.Request().Body).Decode(&param)
		if _, err := govalidator.ValidateStruct(param); err != nil {
			errorconnnets := messageapi.Objectapi{
				Status:  500,
				Message: "Value",
			}
			return c.JSON(http.StatusCreated, errorconnnets)
		} else {
			for i := 0; i < len(param.ID); i++ {
				tx.Table("t_order").Where("id = ?", param.ID[i]).Update("payment_id", param.Payment_id)
			}
			tx.Commit()
			errorconnnet := messageapi.Objectapi{
				Status:  200,
				Message: "Successs",
			}
			writelog.Writelog(errorconnnet)
			return c.JSON(http.StatusBadRequest, errorconnnet)
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
func Order(c echo.Context) error {
	if connectdb.Connnectdb() {
		tx := connectdb.DB.Begin()
		var orders []T_order
		json.NewDecoder(c.Request().Body).Decode(&orders)
		err := tx.Table("t_order").Create(&orders).Error
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
			return c.JSON(http.StatusOK, orders)
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
