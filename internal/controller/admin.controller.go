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

// GetAllMovieAdmin godoc
// @Summary      Get all movies
// @Description  Get list of all movies for admin management
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  dto.Response
// @Failure      401  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /admin/movies [get]
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

// DeleteMovieAdmin godoc
// @Summary      Delete a movie
// @Description  Delete a movie from the database (Requires admin token)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Movie ID"
// @Success      200  {object}  dto.Response
// @Failure      401  {object}  dto.Response
// @Failure      400  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /admin/movies/{id} [delete]
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

// UpdateMovieAdmin godoc
// @Summary      Update a movie
// @Description  Update movie details (Requires admin token)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int                     true  "Movie ID"
// @Param        movie body      dto.UpdateMovieRequest  true  "Movie Update Body"
// @Success      200   {object}  dto.Response
// @Failure      401   {object}  dto.Response
// @Failure      400   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /admin/movies/{id} [patch]
func (ctrl AdminController) UpdateMovieAdmin(c *gin.Context) {
	idParam := c.Param("id")

	movieId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Invalid movie id",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	var req dto.UpdateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	data, err := ctrl.adminService.UpdateMovieAdmin(c.Request.Context(), movieId, req)
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
		Msg:     "Update Movie Success",
		Success: true,
		Data:    data,
	})
}
