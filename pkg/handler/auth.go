package handler

import (
	"AGZ/pkg/structures"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) RegisterUserEmail(c *gin.Context) {
	var input structures.Params
	c.BindJSON(&input.EmailReg)
	fmt.Println(input)
	err := h.services.Authorization.RegisterUserEmail(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": "1",
		"message": "OK",
		"data":    "",
	})
}
func (h *Handler) ConfirmUserEmail(c *gin.Context) {
	var input structures.Params
	c.BindJSON(&input.EmailReg)
	fmt.Println(input)
	err := h.services.Authorization.ConfirmUserEmail(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": "1",
		"message": "OK",
		"data":    "",
	})
}
func (h *Handler) ResetUserEmailPass(c *gin.Context) {
	var input structures.Params
	c.BindJSON(&input.EmailReg)
	fmt.Println(input)
	err := h.services.Authorization.ResetUserEmailPass(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": "1",
		"message": "OK",
		"data":    "",
	})
}
func (h *Handler) SignIn(c *gin.Context) {
	var input structures.Params
	c.BindJSON(&input.User)

	fmt.Println(input)
	res, err := h.services.Authorization.SignIn(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": "1",
		"message": "OK",
		"data":    res,
	})
}

func (h *Handler) ActionUserAccess(c *gin.Context) {
	var token structures.Tokens
	c.BindJSON(&token)
	err := h.services.Authorization.ActionUserAccess(token)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": "1",
		"message": "OK",
		"data":    "",
	})
}

func (h *Handler) ActionUserRefresh(c *gin.Context) {
	var token structures.Tokens
	c.BindJSON(&token)
	token, err := h.services.Authorization.ActionUserRefresh(token)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": "1",
		"message": "OK",
		"data":    token.AccessR,
	})
}
