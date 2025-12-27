package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id          int
	Title       string
	Description string
	Price       float32
	ImgUrl      string
}

var productList []Product

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Abdullah, learning backend, kota km bol")
}

func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "plz give me GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProducts)
	fmt.Println("Server running on: 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server , haha", err)
	}

}

func init() {
	prd1 := Product{
		Id:          1,
		Title:       "Mango",
		Description: "I like mango",
		Price:       100,
		ImgUrl:      "https://img.freepik.com/premium-vector/ripe-delicious-cut-mango-slices-png-transparent-background_1236927-12905.jpg",
	}

	prd2 := Product{
		Id:          2,
		Title:       "Orange",
		Description: "I like orange",
		Price:       100,
		ImgUrl:      "https://pngimg.com/d/orange_PNG791.png",
	}

	prd3 := Product{
		Id:          3,
		Title:       "Banana",
		Description: "I like banana",
		Price:       100,
		ImgUrl:      "https://cdn.imgbin.com/0/7/16/3d-cartoon-fruit-happy-yellow-banana-cartoon-illustration-SjZHDscz.jpg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
}
