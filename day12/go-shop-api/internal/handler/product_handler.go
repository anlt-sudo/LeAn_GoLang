package handler

import (
	"net/http"
	"strconv"

	"go-shop-api/internal/dto"
	"go-shop-api/internal/model"
	"go-shop-api/internal/service"
	"go-shop-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{Service: s}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.Service.GetAll()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.ToProductResponses(products))
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.ToProductResponse(*product))
}

func (h *ProductHandler) Search(c *gin.Context) {
	q := c.Query("q")
	products, err := h.Service.SearchByName(q)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.ToProductResponses(products))
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req dto.ProductRequest
	if err := utils.BindAndValidate(c, &req); err != nil {
		c.Error(err)
		return
	}

	product := model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  &req.CategoryID, // nếu model dùng *uint
	}

	if err := h.Service.Create(&product); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, dto.ToProductResponse(product))
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req dto.ProductRequest
	if err := utils.BindAndValidate(c, &req); err != nil {
		c.Error(err)
		return
	}

	product, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.CategoryID = &req.CategoryID

	if err := h.Service.Update(product); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.ToProductResponse(*product))
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
