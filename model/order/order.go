package order

import (
	"encoding/json"
	"github.com/labstack/echo"
	"tranning_golang/model/connectdb"
	"net/http"
	"tranning_golang/model/messageapi"
	"tranning_golang/model/writelog"
	"time"
	"strconv"
)

type T_order struct {
	ID int `json:"id"`
	Product_id string `json:"code"`
	Product_topping_id string `json:"product_topping_id"`
	User_id string `json:"user_id"`
	Payment_id string `json:"payment"`
	Quantity int `json:"quantity"`
	Amount float64 `json:"amount"`
	Note string `json:"note"`
	Size int `json:"size"`
	Shipment_date time.Time `json:"shipment_date"`
	Status int `json:"status"`
	Date_created time.Time `json:"date_created"`
	Dte_update time.Time `json:"date_updated"`
	Is_delete int `json:"is_delete"`
}

//kiểm tra đơn hang tồn tại 
func checkorder(id int) int{
	
	var result []T_order

	connectdb.DB.Raw("SELECT  * FROM t_order where id = " +strconv.Itoa(id) ).Scan(&result)
	
	
	return  len(result)
}
type Jsonbody struct {
	ID int 
}
// o       Upate thông tin huỷ đơn hàng
func Updateordercancel(c echo.Context) error {
	if (connectdb.Connnectdb()){
		tx := connectdb.DB.Begin()
			var id Jsonbody
			json.NewDecoder(c.Request().Body).Decode(&id)
			
			if(checkorder(id.ID) ==0){
				errorupdate := messageapi.Objectapi{
					Status:500,
					Message:"order no exist",
				}
				writelog.Writelog(errorupdate)
				return  c.JSON(http.StatusBadRequest, errorupdate)
			}else{
				err := tx.Table("t_order").Where(" id = ?", id.ID).Update("status", 0).Error

				if err != nil {
					tx.Rollback()
					errorupdate := messageapi.Objectapi{
						Status:500,
						Message:"update error",
					}
					writelog.Writelog(errorupdate)
					return  c.JSON(http.StatusBadRequest, errorupdate)
				} else {
					tx.Commit()
					successupdate := messageapi.Objectapi{
						Status:200,
						Message:"update success",
						}
					return  c.JSON(http.StatusOK, successupdate)
				}
			}
		

	}else {
		errorconnnet := messageapi.Objectapi{
			Status:500,
			Message:"database disconnect",
		}
		writelog.Writelog(errorconnnet)
		return  c.JSON(http.StatusBadRequest, errorconnnet)
	}

}