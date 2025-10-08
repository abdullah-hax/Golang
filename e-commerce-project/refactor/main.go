package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product // int type er slice e onkgulo int variable thake, struct product typer slice e onkgulo struct product type er variable thakbe

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Abdullah, learning backend, kota km bol ")
}

func getProducts(w http.ResponseWriter, r *http.Request) {

	handleCorsError(w)
	allowPreFlightReq(w, r)

	if r.Method != "GET" { // http.MethodGet
		http.Error(w, "plz give me GET request", 400)
		return
	}

	w.WriteHeader(200)
	sendData(w, productList)

}

func createProduct(w http.ResponseWriter, r *http.Request) {

	handleCorsError(w)
	allowPreFlightReq(w, r)

	if r.Method != "POST" {
		http.Error(w, "Plz give me POST request", 400)
		return
	}

	/*

		jdi frontend pathai tahle r.Body er vetor ei 4ta info ache.
		r.Body => description, imageUrl, price, title

		1. take body information from r.Body
		2. create an instance using struct Product with the body information
		3. append the instance into productList

	*/

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	w.WriteHeader(200) // eta if er pore dile website e statusCode 200 count hbena , 400 tke jabe

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me a valid json", 400)
		return
	}

	newProduct.Id = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct)
}

func handleCorsError(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //  যেকোনো frontend থেকে call দাও , আমি এলাওড
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Abdullah, shafi")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")

}

func allowPreFlightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" { // option successful holei post req krbe , post req i krbe karon FE e leka ache  "const response = await api.get('/products');"
		w.WriteHeader(201)
	}
}

func sendData(w http.ResponseWriter, data interface{}) {

	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()                // router
	mux.HandleFunc("/hello", helloHandler)   // route, helloHandler = handler
	mux.HandleFunc("/about", aboutHandler)   // aboutHandler = handler
	mux.HandleFunc("/products", getProducts) // getProducts = handler
	mux.HandleFunc("/create-products", createProduct)
	fmt.Println("Server running on : 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server haha", err)
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

	// jackfruit img link : https://png.pngtree.com/png-vector/20240722/ourlarge/pngtree-a-beautiful-jackfruit-with-leaf-on-transparent-png-image_13109531.png

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)

}

/*


	type Product struct {
	Id          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
	}

	var Hello int => hello er vetor ekta int thakbe
	var hello []int => hello er vetor onkgulo int thakbe

	var hello Product => hello er vetor er ekta Product thakbe
	var hello []Product => hello er vetor onkgulo Product thakbe


*/
