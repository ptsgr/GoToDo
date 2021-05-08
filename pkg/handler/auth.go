package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ptsgr/GoToDo"
)

func (h *Handler) singIn(c *gin.Context) {
	var input GoToDo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) singUp(c *gin.Context) {

}
