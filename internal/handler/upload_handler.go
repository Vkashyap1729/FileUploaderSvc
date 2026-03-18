package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"file-uploader/internal/service"
)

type UploadHandler struct {
	service *service.UploadService
}

func NewUploadHandler(service *service.UploadService) *UploadHandler {
	return &UploadHandler{service: service}
}

func (h *UploadHandler) GetUploads(c *gin.Context) {
	data, err := h.service.GetAllUploads(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}