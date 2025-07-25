package handler

import (
	"main/dto"
	"main/middleware"
	"main/response"
	"main/service"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	itemService *service.ItemService
}

func NewItemHandler(itemService *service.ItemService) *ItemHandler {
	return &ItemHandler{
		itemService: itemService,
	}
}

func (h *ItemHandler) GetItems(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	boxID := c.Param("id")

	items, err := h.itemService.GetItems(userID.String(), boxID)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.SuccessResponse(c, 200, "items fetched", items)
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	boxID := c.Param("id")

	var req dto.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	item, err := h.itemService.CreateItem(userID.String(), boxID, &req)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.SuccessResponse(c, 201, "item created successfully", item)
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	boxID := c.Param("id")
	itemID := c.Param("itemid")
	if err := h.itemService.DeleteItem(userID.String(), boxID, itemID); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.SuccessResponse(c, 201, "item deleted", nil)
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	userID, err := middleware.GetUserIDFromContext(c)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	boxID := c.Param("id")

	var req dto.UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequestError(c, err.Error())
	}

	if err := h.itemService.UpdateItem(userID.String(), boxID, &req); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.SuccessResponse(c, 201, "item updated", nil)
}
