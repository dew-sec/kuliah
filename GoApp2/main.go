package main

import (
	"goapp2/config"
	"goapp2/controllers/categorycontroller"
	"goapp2/controllers/homecontroller"
	"goapp2/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDb()

	http.HandleFunc("/", homecontroller.Welcome)

	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
