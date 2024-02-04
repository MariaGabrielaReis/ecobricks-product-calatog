package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MariaGabrielaReis/ecobricks-product-catalog/internal/database"
	"github.com/MariaGabrielaReis/ecobricks-product-catalog/internal/service"
	"github.com/MariaGabrielaReis/ecobricks-product-catalog/internal/webserver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ecobricks")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)
	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()

	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Get("/category", webCategoryHandler.GetCategories)
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product", webProductHandler.GetProducts)
	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product/category/{categoryID}", webProductHandler.GetProductsByCategory)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("Server is running on port 8080 🔥")
	http.ListenAndServe(":8080", c)
}
