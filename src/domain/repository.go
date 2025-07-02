package domain

type UserRepository interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
}

type ProductRepository interface {
	Create(product *Product) error
	Update(product *Product) error
	Delete(id uint) error
	GetByID(id uint) (*Product, error)
	GetAll() ([]Product, error)
	GetByMerchantID(merchantID uint) ([]Product, error)
}

type TransactionRepository interface {
	Create(transaction *Transaction) error
	GetByID(id uint) (*Transaction, error)
	GetByCustomerID(customerID uint) ([]Transaction, error)
	GetByMerchantID(merchantID uint) ([]Transaction, error)
}
