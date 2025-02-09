package handler

import (
	"net/http"
	"taskserver/dbserver/internal/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAll(c *gin.Context) {
	items, err := h.services.Containers.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{
		"results": items,
	})
}

func (h *Handler) update(c *gin.Context) {
	var input entity.UpdateContainers
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Containers.Update(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
