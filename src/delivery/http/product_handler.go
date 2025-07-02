package http

import (
	"net/http"
	"strconv"
	"test-aman/src/domain"
	"test-aman/src/lib"
	"test-aman/src/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
}

func NewProductHandler(productUsecase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		lib.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Set merchant ID from token
	merchantID := c.GetUint("userID")
	product.MerchantID = merchantID

	if err := h.productUsecase.CreateProduct(&product); err != nil {
		lib.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lib.SuccessResponse(c, http.StatusCreated, "Product created successfully", product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		lib.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		lib.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	product.ID = uint(id)
	product.MerchantID = c.GetUint("userID")

	if err := h.productUsecase.UpdateProduct(&product); err != nil {
		lib.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lib.SuccessResponse(c, http.StatusOK, "Product updated successfully", product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		lib.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	if err := h.productUsecase.DeleteProduct(uint(id)); err != nil {
		lib.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lib.SuccessResponse(c, http.StatusOK, "Product deleted successfully", nil)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		lib.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.productUsecase.GetProduct(uint(id))
	if err != nil {
		lib.ErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	lib.SuccessResponse(c, http.StatusOK, "Success", product)
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productUsecase.GetAllProducts()
	if err != nil {
		lib.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lib.SuccessResponse(c, http.StatusOK, "Success", products)
}

func (h *ProductHandler) GetMerchantProducts(c *gin.Context) {
	merchantID := c.GetUint("userID")
	products, err := h.productUsecase.GetMerchantProducts(merchantID)
	if err != nil {
		lib.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lib.SuccessResponse(c, http.StatusOK, "Success", products)
}
