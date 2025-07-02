package http

import (
	"net/http"
	"strconv"
	"test-aman/src/domain"
	"test-aman/src/lib"
	"test-aman/src/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionUsecase *usecase.TransactionUsecase
}

func NewTransactionHandler(transactionUsecase *usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{transactionUsecase}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var transaction domain.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		lib.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Set customer ID from token
	transaction.CustomerID = c.GetUint("userID")

	if err := h.transactionUsecase.CreateTransaction(&transaction); err != nil {
		lib.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lib.SuccessResponse(c, http.StatusCreated, "Transaction created successfully", nil)
}

func (h *TransactionHandler) GetTransaction(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		lib.ErrorResponse(c, http.StatusBadRequest, "Invalid transaction ID")
		return
	}

	transaction, err := h.transactionUsecase.GetTransaction(uint(id))
	if err != nil {
		lib.ErrorResponse(c, http.StatusNotFound, "Transaction not found")
		return
	}

	lib.SuccessResponse(c, http.StatusOK, "Success", transaction)
}

func (h *TransactionHandler) GetCustomerTransactions(c *gin.Context) {
	customerID := c.GetUint("userID")
	transactions, err := h.transactionUsecase.GetCustomerTransactions(customerID)
	if err != nil {
		lib.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lib.SuccessResponse(c, http.StatusOK, "Success", transactions)
}

func (h *TransactionHandler) GetMerchantTransactions(c *gin.Context) {
	merchantID := c.GetUint("userID")
	transactions, err := h.transactionUsecase.GetMerchantTransactions(merchantID)
	if err != nil {
		lib.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lib.SuccessResponse(c, http.StatusOK, "Success", transactions)
}
