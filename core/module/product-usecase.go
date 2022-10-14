package module

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/fajartd02/golang-devcamp/core/entity"
	"github.com/fajartd02/golang-devcamp/core/repository"
)

type ProductUsecase interface {
	GetProducts(c *gin.Context) ([]entity.Product, error)
	GetProduct(c *gin.Context) (entity.Product, error)
	CreateProduct(c *gin.Context) error
	UpdateProduct(c *gin.Context) error
	DeleteProduct(c *gin.Context) error
}

type productUsecase struct {
	productRepo repository.ProductRepository
}

// NewProductUsecase use for initiate new product usecase
func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{repo}
}

var ErrProductNotFound = errors.New("product error: ")

func (em *productUsecase) GetProducts(c *gin.Context) ([]entity.Product, error) {
	data, err := em.productRepo.FindAll(c)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrProductNotFound, err)
	}
	return data, nil
}

func (em *productUsecase) GetProduct(c *gin.Context) (entity.Product, error) {
	data, err := em.productRepo.FindSingle(c)
	if err != nil {
		return entity.Product{}, fmt.Errorf("%w: %v", ErrProductNotFound, err)
	}
	return data, nil
}

func (em *productUsecase) CreateProduct(c *gin.Context) error {
	err := em.productRepo.Create(c)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrProductNotFound, err)
	}
	return nil
}

func (em *productUsecase) UpdateProduct(c *gin.Context) error {
	err := em.productRepo.Update(c)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrProductNotFound, err)
	}
	return nil
}

func (em *productUsecase) DeleteProduct(c *gin.Context) error {
	err := em.productRepo.Delete(c)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrProductNotFound, err)
	}
	return nil
}
