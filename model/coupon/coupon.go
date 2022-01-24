package coupon
import (
	"github.com/labstack/echo"
	"tranning_golang/model/connectdb"
	"net/http"
	"tranning_golang/model/messageapi"
	"tranning_golang/model/writelog"
	"time"
)

type Coupon struct {
	ID int `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Discount float64 `json:"discount"`
	Date_expired time.Time `json:"date_expired"`
	Note string `json:"note"`
	Type int `json:"type"`
	Date_created time.Time `json:"date_created"`
	Dte_update time.Time `json:"date_updated"`
	Is_delete int `json:"is_delete"`


}
// o       Thực hiện giảm giá theo mã coupon
func GetCoupon(c echo.Context) error{
	if(connectdb.Connnectdb()){
		var result []Coupon
		code := c.QueryParam("code")
		if(len(code) == 0){
			errorconnnet := messageapi.Objectapi{
				Status:500,
				Message:"requied code",
			}
			writelog.Writelog(errorconnnet)
			return  c.JSON(http.StatusBadRequest, errorconnnet)
		}else{
			connectdb.DB.Raw("select * from t_coupon where is_delete = 0 and  date_expiry > '"+time.Now().Format("2006-01-02 15:04:05")+"' and  code = '"+code+"'").Scan(&result)
			return c.JSON(http.StatusOK, result)
		}

	}else{
		errorconnnet := messageapi.Objectapi{
			Status:500,
			Message:"database disconnect",
		}
		writelog.Writelog(errorconnnet)
		return  c.JSON(http.StatusBadRequest, errorconnnet)
	}

}