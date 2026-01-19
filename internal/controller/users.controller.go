package controller

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"regexp"
	"time"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/err"
	"github.com/Albaihaqi354/Tickitz-BE/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// GetProfile godoc
// @Summary      Get user profile
// @Description  Get current user profile information (Requires user token)
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  dto.Response{data=dto.GetProfile}
// @Failure      401  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /user/profile [get]
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

// GetHistory godoc
// @Summary      Get order history
// @Description  Get user's order history (Requires user token)
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  dto.Response{data=[]dto.GetHistory}
// @Failure      401  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /user/history [get]
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
		Msg:     "Get Seats Success",
		Success: true,
		Data:    history,
	})
}

// UpdatePassword godoc
// @Summary      Update user password
// @Description  Change user password (Requires user token)
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        password  body      dto.UpdatePasswordRequest  true  "Password Update Body"
// @Success      200       {object}  dto.Response
// @Failure      400       {object}  dto.Response
// @Failure      401       {object}  dto.Response
// @Failure      500       {object}  dto.Response
// @Router       /user/password [patch]
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

// UpdateProfile godoc
// @Summary      Update user profile
// @Description  Update user profile information including image upload (Requires user token)
// @Tags         user
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        first_name    formData  string  false  "First Name"
// @Param        last_name     formData  string  false  "Last Name"
// @Param        phone_number  formData  string  false  "Phone Number"
// @Param        image         formData  file    false  "Profile Image"
// @Success      200           {object}  dto.Response{data=dto.UpdateProfileResponse}
// @Failure      400           {object}  dto.Response
// @Failure      401           {object}  dto.Response
// @Failure      500           {object}  dto.Response
// @Router       /user/profile [patch]
func (u UserController) UpdateProfile(c *gin.Context) {
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

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	if req.Image != nil {
		ext := path.Ext(req.Image.Filename)
		re := regexp.MustCompile("^[.](jpg|png)$")
		if !re.Match([]byte(ext)) {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     err.ErrInvalidExt.Error(),
				Error:   "Bad Request",
				Success: false,
				Data:    []any{},
			})
			return
		}

		filename := fmt.Sprintf("%d_profile_%d%s", time.Now().UnixNano(), userIdInt, ext)
		if e := c.SaveUploadedFile(req.Image, filepath.Join("public", "profile", filename)); e != nil {
			log.Println(e.Error())
			c.JSON(http.StatusInternalServerError, dto.Response{
				Msg:     "Internal Server Error",
				Success: false,
				Error:   "internal server error",
				Data:    []any{},
			})
			return
		}
		profileImagePath := fmt.Sprintf("/profile/%s", filename)
		req.ProfileImage = &profileImagePath
	}

	data, err := u.userService.UpdateProfile(c.Request.Context(), userIdInt, req)
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
		Msg:     "Update Profile Success",
		Success: true,
		Data:    []any{data},
	})
}
