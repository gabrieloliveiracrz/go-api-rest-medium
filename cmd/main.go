package main

import (
	"net/http"

	"github.com/gabrieloliveiracrz/go-api-rest-medium/internal/product/productdb"
	"github.com/gabrieloliveiracrz/go-api-rest-medium/internal/product/producthttp"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	productdb.Build()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/products/{id}", producthttp.GetproductByIDHandler)
	r.Get("/products", producthttp.SearchProductsHandler)
	r.Post("/products", producthttp.CreateProductHandler)
	r.Put("/products/{id}", producthttp.UpdateProductHandler)
	r.Delete("/products/{id}", producthttp.DeleteProductHandler)

	http.ListenAndServe(":8081", r)
}
