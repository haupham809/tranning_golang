package product

import (
	"net/http"
	"time"
	"tranning_golang/model/connectdb"
	"tranning_golang/model/messageapi"
	"tranning_golang/model/writelog"

	"github.com/labstack/echo"
)

type T_product struct {
	ID           int       `json:"id"`
	category_id  int       `json:"category"`
	Name         string    `json:"name"`
	Is_suggest   int       `json:"is_suggest"`
	price        float64   `json:"price"`
	Date_created time.Time `json:"date_created"`
	Date_updated time.Time `json:"date_updated"`
	Is_deleted   int       `json:"is_deleted"`
}
type T_size_product struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Product_id   int       `json:"product_id"`
	Price        float64   `json:"price"`
	Date_created time.Time `json:"date_created"`
	Date_updated time.Time `json:"date_updated"`
	Is_deleted   int       `json:"is_deleted"`
}

// o       Lấy chi tiết sản phẩm.
func DetailProduct(c echo.Context) error {

	if connectdb.Connnectdb() {
		var result []T_product
		product_id := c.QueryParam("id")
		if len(product_id) == 0 {
			errorconnnet := messageapi.Objectapi{
				Status:  500,
				Message: "requied product id ",
			}
			writelog.Writelog(errorconnnet)
			return c.JSON(http.StatusOK, errorconnnet)
		} else {
			connectdb.DB.Raw("select * from t_product where is_delete = 0 and id = " + product_id).Scan(&result)
			return c.JSON(http.StatusOK, result)

		}

	} else {
		errorconnnet := messageapi.Objectapi{
			Status:  500,
			Message: "database disconnect",
		}
		writelog.Writelog(errorconnnet)
		return c.JSON(http.StatusOK, errorconnnet)
	}

}

// o       Lấy sản phẩm.
func GetProduct(c echo.Context) error {

	if connectdb.Connnectdb() {
		var result []T_product
		connectdb.DB.Raw("select * from t_product where is_delete = 0").Scan(&result)
		return c.JSON(http.StatusOK, result)

	} else {
		errorconnnet := messageapi.Objectapi{
			Status:  500,
			Message: "database disconnect",
		}
		writelog.Writelog(errorconnnet)
		return c.JSON(http.StatusOK, errorconnnet)
	}

}

// o       Lấy gơij ý sản phẩm.
func GetProductSuggest(c echo.Context) error {

	if connectdb.Connnectdb() {
		var result []T_product
		connectdb.DB.Raw("select * from t_product where is_delete = 0 and id_suggest = 1").Scan(&result)
		return c.JSON(http.StatusOK, result)

	} else {
		errorconnnet := messageapi.Objectapi{
			Status:  500,
			Message: "database disconnect",
		}
		writelog.Writelog(errorconnnet)
		return c.JSON(http.StatusOK, errorconnnet)
	}

}

// o       Danh sách size.
func SizeProduct(c echo.Context) error {

	if connectdb.Connnectdb() {
		var result []T_size_product
		product_id := c.QueryParam("id")
		if len(product_id) == 0 {
			errorconnnet := messageapi.Objectapi{
				Status:  500,
				Message: "requied product id ",
			}
			writelog.Writelog(errorconnnet)
			return c.JSON(http.StatusBadRequest, errorconnnet)
		} else {
			connectdb.DB.Raw("select * from t_size_product where is_delete = 0 and product_id = " + product_id).Scan(&result)
			return c.JSON(http.StatusOK, result)

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
