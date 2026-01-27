package handlers

import (
	"boltdb-editor/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DBHandler handles database operations
type DBHandler struct {
	service *services.DBService
}

// NewDBHandler creates a new DBHandler
func NewDBHandler(service *services.DBService) *DBHandler {
	return &DBHandler{service: service}
}

// OpenDatabase opens a database
func (h *DBHandler) OpenDatabase(c *gin.Context) {
	var req struct {
		Path string `json:"path"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.OpenDatabase(req.Path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Database opened"})
}

// ListBuckets lists buckets
func (h *DBHandler) ListBuckets(c *gin.Context) {
	buckets, err := h.service.ListBuckets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, buckets)
}

// CreateBucket creates a bucket
func (h *DBHandler) CreateBucket(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateBucket(req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bucket created"})
}

// DeleteBucket deletes a bucket
func (h *DBHandler) DeleteBucket(c *gin.Context) {
	name := c.Param("name")
	if err := h.service.DeleteBucket(name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bucket deleted"})
}

// ListKeys lists keys in a bucket
func (h *DBHandler) ListKeys(c *gin.Context) {
	bucket := c.Param("name")
	keys, err := h.service.ListKeys(bucket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, keys)
}

// GetValue gets a value
func (h *DBHandler) GetValue(c *gin.Context) {
	bucket := c.Param("name")
	key := c.Param("key")
	value, err := h.service.GetValue(bucket, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": value})
}

// SetValue sets a value
func (h *DBHandler) SetValue(c *gin.Context) {
	bucket := c.Param("name")
	key := c.Param("key")
	var req struct {
		Value string `json:"value"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.SetValue(bucket, key, req.Value); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Value set"})
}

// DeleteKey deletes a key
func (h *DBHandler) DeleteKey(c *gin.Context) {
	bucket := c.Param("name")
	key := c.Param("key")
	if err := h.service.DeleteKey(bucket, key); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Key deleted"})
}
