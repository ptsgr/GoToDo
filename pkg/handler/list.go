package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ptsgr/GoToDo"
)

type getAllListsResponse struct {
	Data []GoToDo.TodoList `json:"data"`
}

func (h *Handler) createList(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var input GoToDo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	listID, err := h.services.TodoList.Create(userID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": listID,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})

}

func (h *Handler) getListById(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetByID(userID, listID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
