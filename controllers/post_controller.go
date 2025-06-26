package controllers

import (
	"net/http"
	"post_blog/models"
	"post_blog/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	service services.PostService
}

func NewPostController(service services.PostService) *PostController {
	return &PostController{service}
}

func (ctrl *PostController) Create(c *gin.Context) {
	var input models.BlogCreateUpdate
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := ctrl.service.Create(input)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (ctrl *PostController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := ctrl.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func (ctrl *PostController) Update(c *gin.Context) {
	var input models.BlogCreateUpdate
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := ctrl.service.Update(id, input)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func (ctrl *PostController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func (ctrl *PostController) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	items, total, err := ctrl.service.List(page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"data": items, "total": total})
}
