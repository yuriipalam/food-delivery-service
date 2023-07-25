package server

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"food_delivery/config"
	"food_delivery/middleware"
	"food_delivery/repository"
	"food_delivery/server/handler"
	"food_delivery/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start(cfg *config.Config) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DbName, cfg.DbPassword, cfg.DbServer, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.Use(utils.SendCfgToNextMiddleware(cfg))

	customerRepository := repository.NewCustomerRepository(db)
	customerHandler := handler.NewCustomerHandler(customerRepository, cfg)
	customerRouter := r.PathPrefix("/customer").Subrouter()
	customerRouter.Use(middleware.ValidateAccessToken)
	customerRouter.HandleFunc("", customerHandler.GetCustomer).Methods(http.MethodGet)
	customerRouter.HandleFunc("/{id}", customerHandler.UpdateCustomer).Methods(http.MethodPut)
	customerRouter.HandleFunc("/{id}/change-password", customerHandler.UpdateCustomerPassword).Methods(http.MethodPut)
	customerRouter.HandleFunc("/{id}", customerHandler.DeleteCustomer).Methods(http.MethodDelete)

	authHandler := handler.NewAuthHandler(customerRepository, cfg)
	r.HandleFunc("/register", authHandler.Register).Methods(http.MethodPost)
	r.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	refTokenRouter := r.PathPrefix("/refresh").Subrouter()
	refTokenRouter.HandleFunc("", authHandler.Refresh).Methods(http.MethodGet)
	refTokenRouter.Use(middleware.ValidateRefreshToken)

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

	orderRepository := repository.NewOrderRepository(db)
	orderHandler := handler.NewOrderHandler(orderRepository, cfg)
	ordersRouter := r.PathPrefix("/orders").Subrouter()
	ordersRouter.Use(middleware.ValidateAccessToken)
	ordersRouter.HandleFunc("", orderHandler.GetOrders).Methods(http.MethodGet)
	ordersRouter.HandleFunc("", orderHandler.CreateOrder).Methods(http.MethodPost)


	r.HandleFunc("/images/{folder}/{id}/{name}", handler.GetImage).Methods(http.MethodGet)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:5173", "http://127.0.0.1:8888"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	r.Use(handlers.CORS(originsOk, headersOk, methodsOk))

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",

		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: r, // Pass our instance of gorilla/mux in.

	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	fmt.Println("Server is started...")

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
