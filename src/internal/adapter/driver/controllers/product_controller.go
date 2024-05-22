package controllers

import (
	"net/http"
	"strconv"
	httpserver "tech-challenge-fase-1/internal/adapter/driven/infra/http"
	"tech-challenge-fase-1/internal/core/applications/repositories"
	usecases "tech-challenge-fase-1/internal/core/applications/use_cases"
	"tech-challenge-fase-1/internal/core/domain/dtos"
	"tech-challenge-fase-1/internal/core/domain/errors"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productRepository repositories.ProductRepositoryInterface
}

type ProductRequest struct {
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image" `
}

func NewProductController(productRepository repositories.ProductRepositoryInterface) *ProductController {
	return &ProductController{
		productRepository: productRepository,
	}
}

func (pc *ProductController) CreateProduct(c httpserver.HTTPContext) {
	var product ProductRequest
	c.BindJSON(&product)
	if err := product.ValidateProduct(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	respProduct, err := usecases.NewCreateProductUseCase(pc.productRepository).Execute(&dtos.ProductDTO{
		ID:          0,
		Name:        product.Name,
		Category:    product.Category,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"product": respProduct,
	})

}

func (pc *ProductController) UpdateProduct(c httpserver.HTTPContext) {
	var product ProductRequest
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	respProduct, err := usecases.NewUpdateProductUseCase(pc.productRepository).Execute(&dtos.ProductDTO{
		ID:          ID,
		Name:        product.Name,
		Category:    product.Category,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": respProduct,
	})
}

func (pc *ProductController) DeleteProduct(c httpserver.HTTPContext) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err = usecases.NewDeleteProductUseCase(pc.productRepository).Execute(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}

func (pc *ProductController) ListProductsByCategory(c httpserver.HTTPContext) {
	category := c.Param("category")

	if _, ok := errors.ValidCategories[category]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.ErrInvalidCategory.Error()})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "50"))

	products, err := usecases.NewListProductsByCategoryUseCase(pc.productRepository).Execute(category, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
