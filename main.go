package main

import (
	"database/sql"
	"fmt"
	"food_delivery/handler"
	"food_delivery/repository"
	"github.com/gorilla/handlers"
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
	r.HandleFunc("/suppliers", supplierHandler.GetAllSuppliers).Methods(http.MethodGet)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryHandler := handler.NewCategoryHandler(categoryRepository)
	r.HandleFunc("/categories", categoryHandler.GetAllCategories).Methods(http.MethodGet)

	productRepository := repository.NewProductRepository(db)
	productHandler := handler.NewProductHandler(productRepository)
	r.HandleFunc("/product/{id}", productHandler.GetProductByID).Methods(http.MethodGet)
	r.HandleFunc("/products", productHandler.GetAllProductsBySupplierIDAndCategoryID).Queries("supplier_id", "{supplier_id}", "category_id", "{category_id}").Methods(http.MethodGet)
	r.HandleFunc("/products", productHandler.GetAllProductsByCategoryID).Queries("category_id", "{category_id}").Methods(http.MethodGet)
	r.HandleFunc("/products", productHandler.GetAllProductsBySupplierID).Queries("supplier_id", "{supplier_id}").Methods(http.MethodGet)
	r.HandleFunc("/products", productHandler.GetAllProducts).Methods(http.MethodGet)

	fmt.Println("Server is started...")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
