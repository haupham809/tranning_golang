package topping
import (
	"github.com/labstack/echo"
	"tranning_golang/model/connectdb"
	"net/http"
	"tranning_golang/model/messageapi"
	"tranning_golang/model/writelog"
	"tranning_golang/model/validation"
	"github.com/asaskevich/govalidator"
)
// type ToppingProduct struct {
// 	ID string `json:"id"`
// 	Product_id int `json:"name"`
// 	topping_id int `json:"price"`
// 	Date_created time.Time `json:"date_created"`
// 	Date_updated time.Time `json:"date_updated"`
// 	Is_deleted int `json:"is_deleted"`
	
// }
// type Topping struct {
// 	ID int `json:"id"`
// 	Name string `json:"name"`
// 	Price float64 `json:"price"`
// 	Date_created time.Time `json:"date_created"`
// 	Date_updated time.Time `json:"date_updated"`
// 	Is_deleted int `json:"is_deleted"`
// }
type ToppingProduct struct {
	ID int `valid:"required,SqlInjection"`
	Name string  `valid:"required,SqlInjection"`
	Price float64  `valid:"required,SqlInjection"`
}
type Pramproductid struct {
	Product_id string  `valid:"required,SqlInjection"`
}
//o       Danh sách topping theo sản phẩm.
func ToppingForProduct(c echo.Context) error {
	if(connectdb.Connnectdb()){
		var result []ToppingProduct
		var id Pramproductid
		id.Product_id = c.QueryParam("product_id")
		checkvalidation.SqlInjection()
		if _, err := govalidator.ValidateStruct(id); err != nil {
			
			return  c.JSON(http.StatusBadRequest, err)
		}else{
			if(len(id.Product_id) == 0){
				errorconnnet := messageapi.Objectapi{
					Status:500,
					Message:"requied product id ",
				}
				writelog.Writelog(errorconnnet)
				return  c.JSON(http.StatusBadRequest, errorconnnet)
			}else {
				connectdb.DB.Select("   t_topping.name ,t_topping.price ,t_product_topping.id  ").Table("t_topping ,t_product_topping ").Where("t_topping.id = t_product_topping.topping_id and t_topping.is_delete = 0 and t_product_topping.is_delete = 0 ").Where("t_product_topping.product_id =  ?",id.Product_id).Scan(&result)
				return c.JSON(http.StatusOK, result)

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
