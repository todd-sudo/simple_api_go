package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/todd-sudo/todo/internal/dto"
	"github.com/todd-sudo/todo/internal/helper"
	log "github.com/todd-sudo/todo/pkg/logger"
)

func (c *Handler) UpdateUser(ctx *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := ctx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.service.JWT.ValidateToken(authHeader)
	if errToken != nil {
		log.Errorf("token is valid: %v", errToken)
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	user, err := c.service.User.Update(ctx, userUpdateDTO)
	if err != nil {
		log.Errorf("update user error %v", err)
	}
	res := helper.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusOK, res)
}

func (c *Handler) ProfileUser(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.service.JWT.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user, err := c.service.User.Profile(ctx, id)
	if err != nil {
		log.Errorf("profile user error : %v", err)
	}
	res := helper.BuildResponse(true, "OK", user)
	ctx.JSON(http.StatusOK, res)

}
