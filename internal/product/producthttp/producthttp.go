package producthttp

import (
	"net/http"

	"github.com/gabrieloliveiracrz/go-api-rest-medium/internal/encode"
	"github.com/gabrieloliveiracrz/go-api-rest-medium/internal/product/productdecode"
	"github.com/gabrieloliveiracrz/go-api-rest-medium/internal/product/productmain/productservices"
)

var productService = productservices.New()

func GetproductByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := productdecode.DecodeStringIdFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := productService.GetByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	encode.WriteJsonResponse(w, product, http.StatusOK)
}
func SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	productType := productdecode.DecodeTypeQueryString(r)

	products, err := productService.Search(ctx, productType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, products, http.StatusOK)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	productToCreate, err := productdecode.DecodeProductFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := productService.Create(ctx, productToCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, product, http.StatusOK)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := productdecode.DecodeStringIdFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productToUpdate, err := productdecode.DecodeProductFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productToUpdate.ID = id

	product, err := productService.Update(ctx, productToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, product, http.StatusOK)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := productdecode.DecodeStringIdFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = productService.Delete(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJsonResponse(w, nil, http.StatusNoContent)
}
