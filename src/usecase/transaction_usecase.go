package usecase

import (
	"errors"
	"test-aman/src/domain"
	"test-aman/src/dto"
)

type TransactionUsecase struct {
	transactionRepo domain.TransactionRepository
	productRepo     domain.ProductRepository
}

func NewTransactionUsecase(transactionRepo domain.TransactionRepository, productRepo domain.ProductRepository) *TransactionUsecase {
	return &TransactionUsecase{
		transactionRepo: transactionRepo,
		productRepo:     productRepo,
	}
}

func (u *TransactionUsecase) CreateTransaction(transaction *domain.Transaction) error {
	product, err := u.productRepo.GetByID(transaction.ProductID)
	if err != nil {
		return err
	}

	if product.Stock < transaction.Quantity {
		return errors.New("insufficient stock")
	}

	// Calculate total price
	transaction.TotalPrice = float64(transaction.Quantity) * product.Price

	// Apply free shipping and discount rules
	if transaction.TotalPrice >= 15000 {
		transaction.FreeShipping = true
	}
	if transaction.TotalPrice >= 50000 {
		transaction.Discount = transaction.TotalPrice * 0.1 // 10% discount
		transaction.TotalPrice = transaction.TotalPrice - transaction.Discount
	}

	// Update product stock
	product.Stock -= transaction.Quantity
	if err := u.productRepo.Update(product); err != nil {
		return err
	}

	return u.transactionRepo.Create(transaction)
}

func (u *TransactionUsecase) GetTransaction(id uint) (*dto.TransactionResponse, error) {
	transaction, err := u.transactionRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.TransactionResponse{
		ID:           transaction.ID,
		ProductID:    transaction.ProductID,
		CustomerID:   transaction.CustomerID,
		Quantity:     transaction.Quantity,
		TotalPrice:   transaction.TotalPrice,
		FreeShipping: transaction.FreeShipping,
		Discount:     transaction.Discount,
	}, nil
}

func (u *TransactionUsecase) GetCustomerTransactions(customerID uint) ([]dto.TransactionResponse, error) {
	transactions, err := u.transactionRepo.GetByCustomerID(customerID)
	if err != nil {
		return nil, err
	}
	var result []dto.TransactionResponse
	for _, t := range transactions {
		result = append(result, dto.TransactionResponse{
			ID:           t.ID,
			ProductID:    t.ProductID,
			CustomerID:   t.CustomerID,
			Quantity:     t.Quantity,
			TotalPrice:   t.TotalPrice,
			FreeShipping: t.FreeShipping,
			Discount:     t.Discount,
		})
	}
	return result, nil
}

func (u *TransactionUsecase) GetMerchantTransactions(merchantID uint) ([]dto.TransactionResponse, error) {
	transactions, err := u.transactionRepo.GetByMerchantID(merchantID)
	if err != nil {
		return nil, err
	}
	var result []dto.TransactionResponse
	for _, t := range transactions {
		result = append(result, dto.TransactionResponse{
			ID:           t.ID,
			ProductID:    t.ProductID,
			CustomerID:   t.CustomerID,
			Quantity:     t.Quantity,
			TotalPrice:   t.TotalPrice,
			FreeShipping: t.FreeShipping,
			Discount:     t.Discount,
		})
	}
	return result, nil
}
