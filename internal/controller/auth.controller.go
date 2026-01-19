package controller

import (
	"net/http"
	"strings"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Register godoc
// @Summary      Register a new user
// @Description  Create a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      dto.NewUser  true  "User Registration Body"
// @Success      201   {object}  dto.Response{data=dto.RegisterResponse}
// @Failure      400   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /auth/register [post]
func (a AuthController) Register(c *gin.Context) {
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
	data, err := a.authService.Register(c.Request.Context(), newUser)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     "Bad Request",
				Success: false,
				Error:   "Email already in use",
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

// Login godoc
// @Summary      User login
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      dto.LoginRequest  true  "Login Credentials"
// @Success      200          {object}  dto.Response{data=dto.LoginResponse}
// @Failure      400          {object}  dto.Response
// @Failure      401          {object}  dto.Response
// @Router       /auth/login [post]
func (a AuthController) Login(c *gin.Context) {
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

	data, err := a.authService.Login(c.Request.Context(), loginReq)
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
