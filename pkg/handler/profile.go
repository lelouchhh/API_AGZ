package handler

import (
	"AGZ/pkg/structures"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) RemovePurchase(c *gin.Context) {
	var user structures.Params
	c.BindJSON(&user.Purchase)
	err := h.services.Profile.RemovePurchase(user)
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

func (h *Handler) AddPurchase(c *gin.Context) {
	var user structures.Params
	c.BindJSON(&user.Purchase)
	err := h.services.Profile.AddPurchase(user)
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
func (h *Handler) GetBasket(c *gin.Context) {
	var user structures.Params
	c.BindJSON(&user.Purchase)
	res, err := h.services.Profile.GetBasket(user)
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
func (h *Handler) GetLinkBasket(c *gin.Context) {
	var user structures.Params
	c.BindJSON(&user.Link)
	res, err := h.services.Profile.GetLinksBasket(user)
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
func (h *Handler) AddLink(c *gin.Context) {
	var user structures.Params
	c.BindJSON(&user.Link)
	err := h.services.Profile.AddLink(user)
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
func (h *Handler) RemoveLink(c *gin.Context) {
	var user structures.Params
	c.BindJSON(&user.Link)
	err := h.services.Profile.RemoveLink(user)
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
