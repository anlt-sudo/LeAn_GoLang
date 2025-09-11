package handler

import (
	"net/http"
	"strconv"

	"go-shop-api/internal/dto"
	"go-shop-api/internal/model"
	"go-shop-api/internal/service"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	Service *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: s}
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req dto.CategoryRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    category := model.Category{
        Name:        req.Name,
        Description: req.Description,
    }

    if err := h.Service.Create(&category); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    res := dto.CategoryResponse{
        ID:          category.ID,
        Name:        category.Name,
        Description: category.Description,
        CreatedAt:   category.CreatedAt,
        UpdatedAt:   category.UpdatedAt,
    }

    c.JSON(http.StatusCreated, res)
}

func (h *CategoryHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    var req dto.CategoryRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    category, err := h.Service.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }

    category.Name = req.Name
    category.Description = req.Description

    if err := h.Service.Update(category); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    res := dto.CategoryResponse{
        ID:          category.ID,
        Name:        category.Name,
        Description: category.Description,
        CreatedAt:   category.CreatedAt,
        UpdatedAt:   category.UpdatedAt,
    }

    c.JSON(http.StatusOK, res)
}


func (h *CategoryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func (h *CategoryHandler) Search(c *gin.Context) {
    keyword := c.Query("q")
    categories, err := h.Service.SearchByName(keyword)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var res []dto.CategoryResponse
    for _, cat := range categories {
        res = append(res, dto.CategoryResponse{
            ID:          cat.ID,
            Name:        cat.Name,
            Description: cat.Description,
            CreatedAt:   cat.CreatedAt,
            UpdatedAt:   cat.UpdatedAt,
        })
    }

    c.JSON(http.StatusOK, res)
}

