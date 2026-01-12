package controller

import (
	"net/http"
	"strings"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u UserController) AddUser(c *gin.Context) {
	var newUser dto.NewUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
			Data:    []any{},
		})
		return
	}
	data, err := u.userService.AddUser(c.Request.Context(), newUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     "Bad Request",
				Success: false,
				Error:   "Name already in use",
				Data:    []any{},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "internal server error",
			Data:    []any{},
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Msg:     "Create User Success",
		Success: true,
		Data:    []any{data},
	})
}

func (u UserController) Login(c *gin.Context) {
	var loginReq dto.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   "invalid request body",
			Data:    []any{},
		})
		return
	}

	data, err := u.userService.Login(c.Request.Context(), loginReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Msg:     "Unauthorized",
			Success: false,
			Error:   err.Error(),
			Data:    []any{},
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Login Success",
		Success: true,
		Data:    []any{data},
	})
}

func (u UserController) GetProfile(c *gin.Context) {
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Msg:     "Unauthorized",
			Success: false,
			Error:   "User Id Not Found",
			Data:    nil,
		})
		return
	}

	userIdInt, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "Invalid User Id",
			Data:    nil,
		})
		return
	}

	profile, err := u.userService.GetProfile(c.Request.Context(), userIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Get Profile Success",
		Success: true,
		Data:    profile,
	})
}

func (u UserController) GetHistory(c *gin.Context) {
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Msg:     "Unauthorized",
			Success: false,
			Error:   "User Id Not Found",
			Data:    nil,
		})
		return
	}

	userIdInt, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "Invalid User Id",
			Data:    nil,
		})
		return
	}

	history, err := u.userService.GetHistory(c.Request.Context(), userIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Get History Success",
		Success: true,
		Data:    history,
	})
}

func (u UserController) UpdatePassword(c *gin.Context) {
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Msg:     "Unauthorized",
			Success: false,
			Error:   "User Id Not Found",
			Data:    nil,
		})
		return
	}

	userIdInt, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   "Invalid User Id",
			Data:    nil,
		})
		return
	}

	var req dto.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   "Invalid request body",
			Data:    nil,
		})
		return
	}

	err := u.userService.UpdatePassword(c.Request.Context(), userIdInt, req)
	if err != nil {
		if err.Error() == "invalid old password" {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     "Bad Request",
				Success: false,
				Error:   err.Error(),
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Update Password Success",
		Success: true,
		Data:    nil,
	})
}
