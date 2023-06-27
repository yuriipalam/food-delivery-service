package main

import (
	"database/sql"
	"fmt"
	"food_delivery/handler"
	"food_delivery/repository"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	connStr := "postgres://food_delivery:password@localhost:5432/food_delivery?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	supplierRepository := repository.NewSupplierRepository(db)
	supplierHandler := handler.NewSupplierHandler(supplierRepository)
	r.HandleFunc("/suppliers", supplierHandler.GetListOfSuppliers)

	fmt.Println("Server is started...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
