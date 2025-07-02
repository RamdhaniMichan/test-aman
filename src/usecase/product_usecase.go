package usecase

import (
	"errors"
	"test-aman/src/domain"
	"test-aman/src/dto"
)

type ProductUsecase struct {
	productRepo domain.ProductRepository
}

func NewProductUsecase(productRepo domain.ProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepo}
}

func (u *ProductUsecase) CreateProduct(product *domain.Product) error {
	if product.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}
	return u.productRepo.Create(product)
}

func (u *ProductUsecase) UpdateProduct(product *domain.Product) error {
	if product.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}
	return u.productRepo.Update(product)
}

func (u *ProductUsecase) DeleteProduct(id uint) error {
	return u.productRepo.Delete(id)
}

func (u *ProductUsecase) GetProduct(id uint) (*dto.ProductResponse, error) {
	product, err := u.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.Description,
		MerchantID:  product.MerchantID,
	}, nil
}

func (u *ProductUsecase) GetAllProducts() ([]dto.ProductResponse, error) {
	products, err := u.productRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dto.ProductResponse
	for _, p := range products {
		result = append(result, dto.ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Price:       p.Price,
			Stock:       p.Stock,
			Description: p.Description,
			MerchantID:  p.MerchantID,
		})
	}
	return result, nil
}

func (u *ProductUsecase) GetMerchantProducts(merchantID uint) ([]dto.ProductResponse, error) {
	products, err := u.productRepo.GetByMerchantID(merchantID)
	if err != nil {
		return nil, err
	}
	var result []dto.ProductResponse
	for _, p := range products {
		result = append(result, dto.ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Price:       p.Price,
			Stock:       p.Stock,
			Description: p.Description,
			MerchantID:  p.MerchantID,
		})
	}
	return result, nil
}
