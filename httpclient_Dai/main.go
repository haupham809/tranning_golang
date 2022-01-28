package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/products", GetProduct)
	http.HandleFunc("/order", Orders)
	http.HandleFunc("/updates", Update)
	log.Fatal(http.ListenAndServe(":10000", nil))

}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("http://localhost:1323/products")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}

func Orders(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:1323/order"
	fmt.Println("URL:>", url)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("X-Custom-Header", "myvalue")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	// bodys, err := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(bodys))
}

func Update(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	fmt.Println("response Body:", string(body))

	url := "http://localhost:1323/updateorder"
	fmt.Println("URL:>", url)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	client := &http.Client{Transport: &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	},
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	// fmt.Println(resp)
}
