package handler

import (
	"encoding/json"
	"fmt"
	"food_delivery/model"
	"food_delivery/repository"
	"food_delivery/response"
	"food_delivery/utils"
	"net/http"
)

type ProductHandler struct {
	repo repository.ProductRepositoryI
}

func NewProductHandler(repo repository.ProductRepositoryI) *ProductHandler {
	return &ProductHandler{
		repo: repo,
	}
}

func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromMuxVars(r)
	if err != nil {
		response.SendBadRequestError(w, err) // "id must be integer"
		return
	}

	product, err := ph.repo.GetProductByID(id)
	if err != nil {
		response.SendBadRequestError(w, fmt.Errorf("cannot fetch product"))
		return
	} else if product == nil {
		response.SendNotFoundError(w, fmt.Errorf("product not found"))
		return
	}

	productRes, err := ph.getProductResponseFromModel(product)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot create response"))
		return
	}

	response.SendOK(w, productRes)
}

func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := ph.repo.GetAllProducts()
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot fetch products"))
		return
	} else if len(products) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no products found"))
		return
	}

	productsRes, err := ph.getProductResponsesFromModels(products)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot create responses"))
		return
	}

	response.SendOK(w, productsRes)
}

func (ph *ProductHandler) GetAllProductsBySupplierID(w http.ResponseWriter, r *http.Request) {
	supplierID, err := utils.GetIntValueByKeyFromMuxVars("supplier_id", r)
	if err != nil {
		response.SendBadRequestError(w, err) // "key must be integer"
		return
	}

	products, err := ph.repo.GetAllProductsBySupplierID(supplierID)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot fetch products"))
		return
	} else if len(products) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no products found"))
		return
	}

	productsRes, err := ph.getProductResponsesFromModels(products)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot create responses"))
		return
	}

	response.SendOK(w, productsRes)
}

func (ph *ProductHandler) GetAllProductsByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID, err := utils.GetIntValueByKeyFromMuxVars("category_id", r)
	if err != nil {
		response.SendBadRequestError(w, err) // "key must be integer"
		return
	}

	products, err := ph.repo.GetAllProductsByCategoryID(categoryID)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot fetch products"))
		return
	} else if len(products) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no products found"))
		return
	}

	productsRes, err := ph.getProductResponsesFromModels(products)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot create responses"))
		return
	}

	response.SendOK(w, productsRes)
}

func (ph *ProductHandler) GetAllProductsBySupplierIDAndCategoryID(w http.ResponseWriter, r *http.Request) {
	supplierID, err := utils.GetIntValueByKeyFromMuxVars("supplier_id", r)
	if err != nil {
		response.SendBadRequestError(w, err) // "key must be integer"
		return
	}

	categoryID, err := utils.GetIntValueByKeyFromMuxVars("category_id", r)
	if err != nil {
		response.SendBadRequestError(w, err) // "key must be integer"
		return
	}

	products, err := ph.repo.GetAllProductsBySupplierIDAndCategoryID(supplierID, categoryID)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot fetch products"))
		return
	} else if len(products) == 0 {
		response.SendNotFoundError(w, fmt.Errorf("no products found"))
		return
	}

	productsRes, err := ph.getProductResponsesFromModels(products)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("cannot create responses"))
		return
	}

	response.SendOK(w, productsRes)
}

func (ph *ProductHandler) getProductResponseFromModel(product *model.Product) (*response.ProductResponse, error) {
	var productRes response.ProductResponse

	productMarshaled, err := json.Marshal(product)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal product from db")
	}

	err = json.Unmarshal(productMarshaled, &productRes)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal product from db into response")
	}

	productRes.SupplierName, err = ph.repo.GetSupplierNameByID(product.SupplierID)
	if err != nil {
		return nil, err
	}
	productRes.CategoryName, err = ph.repo.GetCategoryNameByID(product.CategoryID)
	if err != nil {
		return nil, err
	}

	productRes.URL = fmt.Sprintf("/products/%d", product.ID)
	productRes.ImageURL = fmt.Sprintf("http://localhost:8080/images/products/%s", product.Image)

	return &productRes, nil
}

func (ph *ProductHandler) getProductResponsesFromModels(products []model.Product) ([]response.ProductResponse, error) {
	var productsRes []response.ProductResponse

	for _, product := range products {
		productRes, err := ph.getProductResponseFromModel(&product)
		if err != nil {
			return nil, err
		}

		productsRes = append(productsRes, *productRes)
	}

	return productsRes, nil
}
