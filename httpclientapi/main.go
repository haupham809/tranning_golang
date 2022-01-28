package main
import (
	"github.com/labstack/echo"
	"fmt"
	"net/http"
	"time"
	 "encoding/json"
	"io/ioutil"
	"strconv"
	"bytes"
)
func main() {
	e := echo.New()
	
	e.GET("/api/v1/detailproduct", Detailproduct)
	e.GET("/api/v1/toppingproduct", GetToppingproduct)
	e.GET("/api/v1/getcoupon", Getcoupon)
	e.PUT("/api/v1/cancelorder", Cancelorder)
	e.GET("/api/v1/getsizeproduct", Getsizeproduct)
	e.POST("/api/v1/savecreditcard", Savecreditcard)
	e.PUT("/api/v1/refundsuccess", Refunndsuccess)
	e.PUT("/api/v1/refundmomo", Refunndmomo)

	e.Logger.Fatal(e.Start(":1313"))

}
//detail product
type T_product struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Is_suggest   int       `json:"is_suggest"`
	price        float64   `json:"price"`
	Date_created time.Time `json:"date_created"`
	Date_updated time.Time `json:"date_updated"`
	Is_deleted   int       `json:"is_deleted"`
}
func Detailproduct(c echo.Context) error {
	
	client := &http.Client {Transport: 
		&http.Transport{ 
			MaxIdleConns: 10, 
			IdleConnTimeout: 30 * time.Second, 
			DisableCompression: true, 
		},

		}
	const apiMethod ="GET"
	requestURL := fmt.Sprintf("%s%s",  "http://localhost:1323", "/detailproduct")
	
	req, _ := http.NewRequest(apiMethod, requestURL, nil)
	q := req.URL.Query()
	fmt.Println(c.QueryParam("id"))
	q.Add("id",c.QueryParam("id"))
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	var detailproduct []T_product
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &detailproduct)
	return c.JSON(http.StatusOK, detailproduct)
}
//topping product
type ToppingProduct struct {
	ID int 
	Name string  
	Price float64 
}
func GetToppingproduct(c echo.Context) error {
	
	client := &http.Client {Transport: 
		&http.Transport{ 
			MaxIdleConns: 10, 
			IdleConnTimeout: 30 * time.Second, 
			DisableCompression: true, 
		},

		}
	const apiMethod ="GET"
	requestURL := fmt.Sprintf("%s%s",  "http://localhost:1323", "/toppingforproduct")
	
	req, _ := http.NewRequest(apiMethod, requestURL, nil)
	q := req.URL.Query()
	fmt.Println(c.QueryParam("product_id"))
	q.Add("product_id",c.QueryParam("product_id"))
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	var Topping []ToppingProduct
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, &Topping)
	return c.JSON(http.StatusOK, Topping)
}

//get coupon information
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
func Getcoupon(c echo.Context) error {
	
	client := &http.Client {Transport: 
		&http.Transport{ 
			MaxIdleConns: 10, 
			IdleConnTimeout: 30 * time.Second, 
			DisableCompression: true, 
		},

		}
	const apiMethod ="GET"
	requestURL := fmt.Sprintf("%s%s",  "http://localhost:1323", "/coupon")
	
	req, _ := http.NewRequest(apiMethod, requestURL, nil)
	q := req.URL.Query()
	fmt.Println(c.QueryParam("code"))
	q.Add("code",c.QueryParam("code"))
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	var Topping []Coupon
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, &Topping)
	return c.JSON(http.StatusOK, Topping)
}

//cancel order
type Objectapi struct {
	Status       int   
	Message string  
}
type Jsonbody struct {
	ID int 
}
func Cancelorder(c echo.Context) error {
	
	client := &http.Client {Transport: 
		&http.Transport{ 
			MaxIdleConns: 10, 
			IdleConnTimeout: 30 * time.Second, 
			DisableCompression: true, 
		},

		}
	const apiMethod ="PUT"
	requestURL := fmt.Sprintf("%s%s",  "http://localhost:1323", "/cancelorder")
	
	var id Jsonbody
	json.NewDecoder(c.Request().Body).Decode(&id)
	fmt.Println(strconv.Itoa(id.ID))
	postBody, _ := json.Marshal(map[string]int{
		"id":  id.ID,
	 })
	 responseBody := bytes.NewBuffer(postBody)
	 req, _ := http.NewRequest(apiMethod, requestURL, responseBody)
	 req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	var result Objectapi
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, &result)
	return c.JSON(http.StatusOK, result)
}

// get size product 
type T_size_product struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Product_id   int       `json:"product_id"`
	Price        float64   `json:"price"`
	Date_created time.Time `json:"date_created"`
	Date_updated time.Time `json:"date_updated"`
	Is_deleted   int       `json:"is_deleted"`
}
func Getsizeproduct(c echo.Context) error {
	
	client := &http.Client {Transport: 
		&http.Transport{ 
			MaxIdleConns: 10, 
			IdleConnTimeout: 30 * time.Second, 
			DisableCompression: true, 
		},

		}
	const apiMethod ="GET"
	requestURL := fmt.Sprintf("%s%s",  "http://localhost:1323", "/sizeproduct")
	
	req, _ := http.NewRequest(apiMethod, requestURL, nil)
	q := req.URL.Query()
	q.Add("id",c.QueryParam("id"))
	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)
	var result []T_size_product
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, &result)
	return c.JSON(http.StatusOK, result)
}

// save credit 
type Creditcard struct {
	ID int 
	Cc_number string `valid:"required,SqlInjection"`
	Cc_expiry string  
	Cc_type_payment int `valid:"required,SqlInjection"`
	Date_created string 
	Date_update string 
	Is_delete  int 

}
func Savecreditcard(c echo.Context) error {
	
	client := &http.Client {Transport: 
		&http.Transport{ 
			MaxIdleConns: 10, 
			IdleConnTimeout: 30 * time.Second, 
			DisableCompression: true, 
		},

		}
	const apiMethod ="POST"
	requestURL := fmt.Sprintf("%s%s",  "http://localhost:1323", "/savecreditcard")
	
	var inforbody Creditcard
	json.NewDecoder(c.Request().Body).Decode(&inforbody)
	fmt.Println(inforbody)
	postBody, _ := json.Marshal(map[string]interface{}{
		"Cc_number":inforbody.Cc_number,
        "Cc_expiry":inforbody.Cc_number,
        "Cc_type_payment":inforbody.Cc_type_payment,
	 })
	 responseBody := bytes.NewBuffer(postBody)
	 req, _ := http.NewRequest(apiMethod, requestURL, responseBody)
	
	resp, _ := client.Do(req)
	var result Objectapi
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, &result)
	return c.JSON(http.StatusOK, result)
}

//resfund success
func Refunndsuccess(c echo.Context) error {
	
	client := &http.Client {Transport: 
		&http.Transport{ 
			MaxIdleConns: 10, 
			IdleConnTimeout: 30 * time.Second, 
			DisableCompression: true, 
		},

		}
	const apiMethod ="PUT"
	requestURL := fmt.Sprintf("%s%s",  "http://localhost:1323", "/updatepaymentcancel")
	
	var id Jsonbody
	json.NewDecoder(c.Request().Body).Decode(&id)
	
	postBody, _ := json.Marshal(map[string]string{
		"id":  strconv.Itoa(id.ID),
	 })
	 responseBody := bytes.NewBuffer(postBody)
	 req, _ := http.NewRequest(apiMethod, requestURL, responseBody)
	 req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	var result Objectapi
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, &result)
	return c.JSON(http.StatusOK, result)
}

//refund momo
func Refunndmomo(c echo.Context) error {
	
	client := &http.Client {Transport: 
		&http.Transport{ 
			MaxIdleConns: 10, 
			IdleConnTimeout: 30 * time.Second, 
			DisableCompression: true, 
		},

		}
	const apiMethod ="PUT"
	requestURL := fmt.Sprintf("%s%s",  "http://localhost:1323", "/refundmomo")
	
	var id Jsonbody
	json.NewDecoder(c.Request().Body).Decode(&id)
	
	postBody, _ := json.Marshal(map[string]interface{}{
		"id":  strconv.Itoa(id.ID),
	 })
	 responseBody := bytes.NewBuffer(postBody)
	 req, _ := http.NewRequest(apiMethod, requestURL, responseBody)
	 req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	var result Objectapi
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, &result)
	return c.JSON(http.StatusOK, result)
}