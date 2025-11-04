package controller

import (
	"go-posts/dto"
	"go-posts/errs"
	"go-posts/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	service services.PostService
}

func NewPostHandler(service services.PostService) *PostController {
	return &PostController{service}
}

func (h *PostController) GetPosts(c *gin.Context) {
	limitString := c.Param("limit")
	offsetString := c.Param("offset")

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}

	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset"})
		return
	}

	posts, total, err := h.service.GetPosts(offset, limit)
	if err != nil {
		if httpErr, ok := err.(*errs.HTTPError); ok {
			c.JSON(httpErr.StatusCode, gin.H{"message": httpErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  posts,
		"total": total,
	})
}

func (h *PostController) GetPostByID(c *gin.Context) {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	posts, err := h.service.GetPostByID(id)
	if err != nil {
		if httpErr, ok := err.(*errs.HTTPError); ok {
			c.JSON(httpErr.StatusCode, gin.H{"message": httpErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func (h *PostController) CreatePost(c *gin.Context) {
	var post dto.PublicPost
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err := h.service.CreatePost(post)
	if err != nil {
		if httpErr, ok := err.(*errs.HTTPError); ok {
			c.JSON(httpErr.StatusCode, gin.H{"message": httpErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}

func (h *PostController) UpdatePost(c *gin.Context) {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var post dto.PublicPost
	if err = c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err = h.service.UpdatePost(post, id)
	if err != nil {
		if httpErr, ok := err.(*errs.HTTPError); ok {
			c.JSON(httpErr.StatusCode, gin.H{"message": httpErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

func (h *PostController) DeletePost(c *gin.Context) {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.service.DeletePost(id)
	if err != nil {
		if httpErr, ok := err.(*errs.HTTPError); ok {
			c.JSON(httpErr.StatusCode, gin.H{"message": httpErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
