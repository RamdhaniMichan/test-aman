package repository

import (
	"test-aman/src/domain"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Create(transaction *domain.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetByID(id uint) (*domain.Transaction, error) {
	var transaction domain.Transaction
	err := r.db.Preload("Product").Preload("Customer").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionRepository) GetByCustomerID(customerID uint) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Preload("Product").Where("customer_id = ?", customerID).Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) GetByMerchantID(merchantID uint) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Preload("Product").
		Joins("JOIN products ON transactions.product_id = products.id").
		Where("products.merchant_id = ?", merchantID).
		Find(&transactions).Error
	return transactions, err
}
