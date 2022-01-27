package coupon

import (
	"net/http"
	"time"
	"tranning_golang/model/connectdb"
	"tranning_golang/model/messageapi"
	"tranning_golang/model/writelog"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"tranning_golang/model/validation"
)

type Coupon struct {
	ID           int       `json:"id"`
	Code         string    `json:"code"`
	Name         string    `json:"name"`
	Discount     float64   `json:"discount"`
	Date_expired time.Time `json:"date_expired"`
	Note         string    `json:"note"`
	Type         int       `json:"type"`
	Date_created time.Time `json:"date_created"`
	Dte_update   time.Time `json:"date_updated"`
	Is_delete    int       `json:"is_delete"`
}
type Pramcode struct {
	Code string  `valid:"required,SqlInjection"`
}
// o       Thực hiện giảm giá theo mã coupon
func GetCoupon(c echo.Context) error {
	if connectdb.Connnectdb() {
		var result []Coupon
		var code Pramcode
		code.Code = c.QueryParam("code")
		checkvalidation.SqlInjection()
		if _, err := govalidator.ValidateStruct(code); err != nil {
			
			return  c.JSON(http.StatusBadRequest, err)
		}else{
			if (len(code.Code) == 0) {
				errorconnnet := messageapi.Objectapi{
					Status:  500,
					Message: "requied code",
				}
				writelog.Writelog(errorconnnet)
				return c.JSON(http.StatusBadRequest, errorconnnet)
			} else {
				connectdb.DB.Select("* ").Table("t_coupon").Where("is_delete = ?",0).Where("date_expiry > ?",time.Now().Format("2006-01-02 15:04:05")).Where("code = ?",code.Code).Scan(&result)
				return c.JSON(http.StatusOK, result)
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
type Pramtypecoupon struct {
	Type_coupon string  `valid:"required,SqlInjection"`
}
func GetCouponByUserLogin(c echo.Context) error {
	if connectdb.Connnectdb() {
		var result []Coupon
		var type_coupon Pramtypecoupon
		type_coupon.Type_coupon = c.QueryParam("type")
		checkvalidation.SqlInjection()
		if _, err := govalidator.ValidateStruct(type_coupon); err != nil {
			
			return  c.JSON(http.StatusBadRequest, err)
		}else{
			if len(type_coupon.Type_coupon) == 0 {
				errorconnnet := messageapi.Objectapi{
					Status:  500,
					Message: "requied type",
				}
				writelog.Writelog(errorconnnet)
				return c.JSON(http.StatusBadRequest, errorconnnet)
			} else {
				connectdb.DB.Select("* ").Table("t_coupon").Where("is_delete = ?",0).Where("type = ?",type_coupon.Type_coupon).Scan(&result)
				return c.JSON(http.StatusOK, result)
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
