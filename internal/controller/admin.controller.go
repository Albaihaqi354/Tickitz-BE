package controller

import (
	"net/http"
	"strconv"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/service"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService *service.AdminService
}

func NewAdminController(adminService *service.AdminService) *AdminController {
	return &AdminController{
		adminService: adminService,
	}
}

func (ctrl AdminController) GetAllMovieAdmin(c *gin.Context) {
	movies, err := ctrl.adminService.GetAllMovieAdmin(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   err.Error(),
			Data:    []any{},
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Get All Movies Success",
		Success: true,
		Data:    movies,
	})
}

func (ctrl AdminController) DeleteMovieAdmin(c *gin.Context) {
	idParam := c.Param("id")

	movieId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Invalid movie id",
			Success: false,
			Error:   err.Error(),
			Data:    []any{},
		})
		return
	}

	err = ctrl.adminService.DeleteMovieAdmin(c.Request.Context(), movieId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Failed to delete movie",
			Success: false,
			Error:   err.Error(),
			Data:    []any{},
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Delete Movie Success",
		Success: true,
	})
}
