package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/todd-sudo/todo/internal/dto"
	"github.com/todd-sudo/todo/internal/helper"
	log "github.com/todd-sudo/todo/pkg/logger"
)

// Аутентификация
func (c *Handler) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.service.Auth.VerifyCredential(ctx, loginDTO.Email, loginDTO.Password)
	if err != nil {
		response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}

	generatedToken := c.service.JWT.GenerateToken(strconv.FormatUint(user.ID, 10))
	user.Token = generatedToken
	response := helper.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusOK, response)

}

// Регистрация
func (c *Handler) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	statusEmail, err := c.service.Auth.IsDuplicateEmail(ctx, registerDTO.Email)
	log.Info(statusEmail)
	if err == nil || statusEmail {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser, err := c.service.Auth.CreateUser(ctx, registerDTO)
		if err != nil {
			log.Errorf("create user failed: %v", err)
		}
		token := c.service.JWT.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
