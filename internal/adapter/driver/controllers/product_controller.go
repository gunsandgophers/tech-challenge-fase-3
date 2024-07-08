package controllers

import (
	"net/http"
	"strconv"
	httpserver "tech-challenge-fase-1/internal/adapter/driven/infra/http"
	"tech-challenge-fase-1/internal/core/applications/repositories"
	usecases "tech-challenge-fase-1/internal/core/applications/use_cases"
	"tech-challenge-fase-1/internal/core/domain/dtos"
	"tech-challenge-fase-1/internal/core/domain/errors"
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

// CreateProduct godoc
// @Summary      Create a product
// @Description  register the product information
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product   body      ProductRequest  true  "Insert Product"
// @Success      200  {object}  dtos.ProductDTO
// @Failure      400 {string} string "when invalid params"
// @Failure      500 {string} string "when create process error"
// @Router       /product/ [post]
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
		c.JSON(http.StatusInternalServerError, httpserver.Payload{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, httpserver.Payload{
		"product": respProduct,
	})

}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  change the product information
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product   body      ProductRequest  true  "Update Product"
// @Success      200  {object}  dtos.ProductDTO
// @Failure      400 {string} string "when invalid params"
// @Failure      500 {string} string "when update process error"
// @Router       /product/{id}/ [put]
func (pc *ProductController) UpdateProduct(c httpserver.HTTPContext) {
	var product ProductRequest
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, httpserver.Payload{
			"error": err.Error()})
		return
	}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, httpserver.Payload{
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
		c.JSON(http.StatusInternalServerError, httpserver.Payload{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, httpserver.Payload{
		"product": respProduct,
	})
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  remove the product information
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product   body      ProductRequest  true  "Delete Product"
// @Success      200 {string} string
// @Failure      400 {string} string "when invalid params"
// @Failure      500 {string} string "when delete process error"
// @Router       /product/{id}/ [delete]
func (pc *ProductController) DeleteProduct(c httpserver.HTTPContext) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, httpserver.Payload{
			"error": err.Error()})
		return
	}

	err = usecases.NewDeleteProductUseCase(pc.productRepository).Execute(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpserver.Payload{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, httpserver.Payload{
		"message": "Product deleted successfully",
	})
}

// ListProductsByCategory godoc
// @Summary      List products
// @Description  List a set of products information over followed categories
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        category   path      string  true  "List Product By Category"
// @Success      200 {array} dtos.ProductDTO
// @Failure      400 {string} string "when invalid params"
// @Failure      500 {string} string "when list process error"
// @Router       /product/{category}/ [get]
func (pc *ProductController) ListProductsByCategory(c httpserver.HTTPContext) {
	category := c.Param("category")

	if _, ok := errors.ValidCategories[category]; !ok {
		c.JSON(http.StatusBadRequest, httpserver.Payload{
			"error": errors.ErrInvalidCategory.Error()})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "50"))

	products, err := usecases.NewListProductsByCategoryUseCase(pc.productRepository).Execute(category, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpserver.Payload{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, httpserver.Payload{
		"products": products,
	})
}
