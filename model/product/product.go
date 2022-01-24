package product
import (
	"github.com/labstack/echo"
	"tranning_golang/model/connectdb"
	"net/http"
	"tranning_golang/model/messageapi"
	"tranning_golang/model/writelog"
	
)

func DetailProduct(c echo.Context) error {
 
	if(connectdb.Connnectdb()){
		susscess := messageapi.Objectapi{
			Status:200,
			Message:"database connect",
		}
		writelog.Writelog(susscess)

		return  c.JSON(http.StatusOK, susscess)

		
	}else {
		errorconnnet := messageapi.Objectapi{
			Status:500,
			Message:"database disconnect",
		}
		writelog.Writelog(errorconnnet)
		return  c.JSON(http.StatusOK, errorconnnet)
	}
 

}

// func ToppingForProduct(c echo.Context) error {

// }