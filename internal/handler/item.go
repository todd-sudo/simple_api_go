package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/todd-sudo/todo/internal/dto"
	"github.com/todd-sudo/todo/internal/helper"
	"github.com/todd-sudo/todo/internal/model"
	log "github.com/todd-sudo/todo/pkg/logger"
)

// Получениек всех item
func (c *Handler) AllItem(ctx *gin.Context) {
	var items []model.Item = c.service.Item.All(ctx)
	res := helper.BuildResponse(true, "OK", items)
	ctx.JSON(http.StatusOK, res)
}

// Поиск Item по ID
func (c *Handler) FindByIDItem(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var item model.Item = c.service.Item.FindByID(ctx, id)
	if (item == model.Item{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", item)
		ctx.JSON(http.StatusOK, res)
	}
}

// Добавление Item
func (c *Handler) InsertItem(ctx *gin.Context) {
	var itemCreateDTO dto.ItemCreateDTO
	errDTO := ctx.ShouldBind(&itemCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			itemCreateDTO.UserID = convertedUserID
		}
		result := c.service.Item.Insert(ctx, itemCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		ctx.JSON(http.StatusCreated, response)
	}
}

// Обновление Item
func (c *Handler) UpdateItem(ctx *gin.Context) {
	var itemUpdateDTO dto.ItemUpdateDTO
	errDTO := ctx.ShouldBind(&itemUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.service.JWT.ValidateToken(authHeader)
	if errToken != nil {
		log.Error(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.service.Item.IsAllowedToEdit(ctx, userID, itemUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			itemUpdateDTO.UserID = id
		}
		result := c.service.Item.Update(ctx, itemUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}
}

// Удаление Item
func (c *Handler) DeleteItem(ctx *gin.Context) {
	var item model.Item
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}
	item.ID = id
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.service.JWT.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.service.Item.IsAllowedToEdit(ctx, userID, item.ID) {
		c.service.Item.Delete(ctx, item)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}
}

// Получение User ID по токену
func (c *Handler) getUserIDByToken(token string) string {
	aToken, err := c.service.JWT.ValidateToken(token)
	if err != nil {
		log.Error(err)
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
