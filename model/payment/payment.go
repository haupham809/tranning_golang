package payment
import (
	"github.com/labstack/echo"
	"tranning_golang/model/connectdb"
	"net/http"
	"tranning_golang/model/messageapi"
	"tranning_golang/model/writelog"
	"time"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"fmt"
	"tranning_golang/model/order"
)
type Payment struct {
	ID int 
	UserID int `valid:"required"`
	Credit_card_id int `valid:"required"`
	Coupon_id int `valid:"required"`
	Total_amount float64 `valid:"required"`
	Payment_result_id string `valid:"required"`
	Payment_date time.Time 
	Refund_date time.Time 
	Is_refund int `valid:"required"`
	Date_created time.Time 
	Date_update time.Time 

}
type Creditcard struct {
	ID int 
	Cc_number string `valid:"required"`
	Cc_expiry time.Time 
	Cc_type_payment int `valid:"required"`
	Date_created time.Time 
	Date_update time.Time 
	Is_delete  int 

}
func checkpaymentresultid(Payment_result_id string) int {
	var result []Payment

	connectdb.DB.Raw("SELECT  * FROM t_payment where payment_result_id = "+Payment_result_id ).Scan(&result)
	
	return  len(result)
}

//lưu thông tin thanh toán 
func SavePayment(c echo.Context) error {
	if(connectdb.Connnectdb()){
		tx := connectdb.DB.Begin()
		var payment Payment
		json.NewDecoder(c.Request().Body).Decode(&payment)
		if _, err := govalidator.ValidateStruct(payment); err != nil {
			return  c.JSON(http.StatusCreated, err)
		}else{
			if(checkpaymentresultid(payment.Payment_result_id)==0){
				payment.Payment_date = time.Now()
				payment.Refund_date = time.Now()
				payment.Date_created = time.Now()
				payment.Date_update = time.Now()
	
				err := tx.Table("t_payment").Create(&payment).Error
	
					if err != nil {
						tx.Rollback()
						errorsave := messageapi.Objectapi{
							Status:500,
							Message:"Save payment error",
						}
						writelog.Writelog(errorsave)
						return  c.JSON(http.StatusBadRequest, errorsave)
					} else {
						tx.Commit()
						successsave := messageapi.Objectapi{
							Status:200,
							Message:"Save payment success",
							}
						return  c.JSON(http.StatusOK, successsave)
					}
			}else{
				errorsave := messageapi.Objectapi{
					Status:500,
					Message:" payment_result_id exist",
				}
				writelog.Writelog(errorsave)
				return  c.JSON(http.StatusBadRequest, errorsave)
			}

			

			
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

//hoàn tiền thành công
type Jsonbody struct {
	ID string `valid:"required"`
}
func checkpaymentid(id string) int {
	var result []Payment

	connectdb.DB.Raw("SELECT  * FROM t_payment where id = "+id ).Scan(&result)

	return  len(result)
}

func Updatepaymentcancel(c echo.Context) error {
	if(connectdb.Connnectdb()){
		tx := connectdb.DB.Begin()
			var id Jsonbody
			json.NewDecoder(c.Request().Body).Decode(&id)
			if _, err := govalidator.ValidateStruct(id); err != nil {
				return  c.JSON(http.StatusCreated, err)
			}else{
				if(checkpaymentid(id.ID) ==0){
					errorupdate := messageapi.Objectapi{
						Status:500,
						Message:"payment no exist",
					}
					writelog.Writelog(errorupdate)
					return  c.JSON(http.StatusBadRequest, errorupdate)
				}else{

					err := tx.Table("t_payment").Where(" id = ?", id.ID).Update("is_refund", 1).Update("refund_date",time.Now()).Update("date_update",time.Now()).Error
					if err != nil {
						tx.Rollback()
						errorupdate := messageapi.Objectapi{
							Status:500,
							Message:"update payment error",
						}
						writelog.Writelog(errorupdate)
						return  c.JSON(http.StatusBadRequest, errorupdate)
					} else {
						tx.Commit()
						successupdate := messageapi.Objectapi{
							Status:200,
							Message:"update payment success",
							}
						return  c.JSON(http.StatusOK, successupdate)
					}
				}
			}
	}else{
		errorsave := messageapi.Objectapi{
			Status:500,
			Message:" credit card exist",
		}
		writelog.Writelog(errorsave)
		return  c.JSON(http.StatusBadRequest, errorsave)
	}
}

//luu thông tin credit
func checkcreditcard(ccnumber string) int {
	var result []Creditcard
	connectdb.DB.Raw("SELECT  * FROM t_credit_card where cc_number = "+ccnumber ).Scan(&result)
	return  len(result)
}
func SaveCreditCard(c echo.Context) error {
	if(connectdb.Connnectdb()){
		tx := connectdb.DB.Begin()
		var card Creditcard
		json.NewDecoder(c.Request().Body).Decode(&card)
		
		if _, err := govalidator.ValidateStruct(card); err != nil {
			return  c.JSON(http.StatusCreated, err)
		}else{
			if(checkcreditcard(card.Cc_number)==0){
				card.Date_created = time.Now()
				card.Date_update = time.Now()
				card.Cc_expiry = time.Now()
				err := tx.Table("t_credit_card").Create(&card).Error
	
					if err != nil {
						tx.Rollback()
						errorsave := messageapi.Objectapi{
							Status:500,
							Message:"Save credit card error",
						}
						writelog.Writelog(errorsave)
						return  c.JSON(http.StatusBadRequest, errorsave)
					} else {
						tx.Commit()
						successsave := messageapi.Objectapi{
							Status:200,
							Message:"Save credit card success",
							}
						return  c.JSON(http.StatusOK, successsave)
					}
			}else{
				errorsave := messageapi.Objectapi{
					Status:500,
					Message:" credit card exist",
				}
				writelog.Writelog(errorsave)
				return  c.JSON(http.StatusBadRequest, errorsave)
			}

			

			
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
//hoan tiền momo
type Momobody struct {
		partnerCode string
		orderId string
		requestId string
		amount float64
		transId float64
		lang string
		description string
		signature string
	  
	  
}
func Getpayment(id string) order.T_order{
	var result []order.T_order

	connectdb.DB.Raw("SELECT  * FROM t_order where id = "+id ).Scan(&result)
	return  result[0]
}
func Refundmomo(c echo.Context) error {
	if(connectdb.Connnectdb()){
		var id Jsonbody
		json.NewDecoder(c.Request().Body).Decode(&id)
		if _, err := govalidator.ValidateStruct(id); err != nil {
			return  c.JSON(http.StatusCreated, err)
		}else{
			fmt.Println(id)
			if(checkpaymentid(id.ID) ==0){
				errorupdate := messageapi.Objectapi{
					Status:500,
					Message:"payment no exist",
				}
				writelog.Writelog(errorupdate)
				return  c.JSON(http.StatusBadRequest, errorupdate)
			}else{
				inforpayment:=Getpayment(id.ID)
				// resp, err := http.Post("https://test-payment.momo.vn/v2/gateway/api/refund","application/json", bytes.NewBuffer(jsonPayload)
				fmt.Println(inforpayment)
					successrefund := messageapi.Objectapi{
					Status:200,
					Message:"refund success",
					}
					return  c.JSON(http.StatusOK, successrefund)
			}
		
		}
	}else{
		errorsave := messageapi.Objectapi{
			Status:500,
			Message:" credit card exist",
		}
		writelog.Writelog(errorsave)
		return  c.JSON(http.StatusBadRequest, errorsave)
	}


}