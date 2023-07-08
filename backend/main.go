package main

import (
	"database/sql"
	"fmt"
	"food_delivery/config"
	"food_delivery/handler"
	"food_delivery/middleware"
	"food_delivery/repository"
	"food_delivery/utils"
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

	cfg := config.NewConfig()

	r := mux.NewRouter()
	r.Use(utils.SendCfgToMiddleware(cfg))

	customerRepository := repository.NewCustomerRepository(db)
	customerHandler := handler.NewCustomerHandler(customerRepository, cfg)
	r.HandleFunc("/customer/{id}", customerHandler.GetCustomerByID).Methods(http.MethodGet)
	r.HandleFunc("/customer/{id}", customerHandler.UpdateCustomerByID).Methods(http.MethodPut)
	r.HandleFunc("/customer/{id}/change-password", customerHandler.UpdateCustomerPasswordByID).Methods(http.MethodPut)
	r.HandleFunc("/customer/{id}", customerHandler.DeleteCustomerByID).Methods(http.MethodDelete)
	//r.HandleFunc("/customer", customerHandler.CreateCustomer).Methods(http.MethodPost)

	authHandler := handler.NewAuthHandler(customerRepository, cfg)
	r.HandleFunc("/register", authHandler.Register).Methods(http.MethodPost)
	r.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	refToken := r.PathPrefix("/refresh").Subrouter()
	refToken.HandleFunc("", authHandler.Refresh).Methods(http.MethodGet)
	refToken.Use(middleware.ValidateRefreshToken)

	supplierRepository := repository.NewSupplierRepository(db)
	supplierHandler := handler.NewSupplierHandler(supplierRepository)
	r.HandleFunc("/supplier/{id}", supplierHandler.GetSupplierByID).Methods(http.MethodGet)
	r.HandleFunc("/suppliers", supplierHandler.GetSuppliersByCategoryID).Queries("category_id", "{category_id}").Methods(http.MethodGet)
	r.HandleFunc("/suppliers", supplierHandler.GetAllSuppliers).Methods(http.MethodGet)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryHandler := handler.NewCategoryHandler(categoryRepository)
	r.HandleFunc("/category/{id}", categoryHandler.GetCategoryByID).Methods(http.MethodGet)
	r.HandleFunc("/categories", categoryHandler.GetCategoriesBySupplierID).Queries("supplier_id", "{supplier_id}").Methods(http.MethodGet)
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
