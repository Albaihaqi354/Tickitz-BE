package controller

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/err"
	"github.com/Albaihaqi354/Tickitz-BE/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
// @Description  Get list of all movies for admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  dto.Response
// @Failure      401  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /admin [get]
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
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        id                path      int     true   "Movie ID"
// @Param        title             formData  string  false  "Movie Title"
// @Param        synopsis          formData  string  false  "Movie Synopsis"
// @Param        duration          formData  int     false  "Movie Duration"
// @Param        release_date      formData  string  false  "Movie Release Date (YYYY-MM-DD)"
// @Param        director_id       formData  int     false  "Director ID"
// @Param        poster            formData  file    false  "Poster Image"
// @Param        backdrop          formData  file    false  "Backdrop Image"
// @Param        popularity_score  formData  number  false  "Popularity Score"
// @Success      200               {object}  dto.Response
// @Failure      401               {object}  dto.Response
// @Failure      400               {object}  dto.Response
// @Failure      500               {object}  dto.Response
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
	if err := c.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	if req.Poster != nil {
		ext := path.Ext(req.Poster.Filename)
		re := regexp.MustCompile("^[.](jpg|png|jpeg)$")
		if !re.MatchString(ext) {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     "Invalid image extension",
				Error:   "Bad Request",
				Success: false,
				Data:    nil,
			})
			return
		}

		filename := fmt.Sprintf("%d_poster_%s", time.Now().UnixNano(), req.Poster.Filename)
		if e := c.SaveUploadedFile(req.Poster, filepath.Join("public", "movie", filename)); e != nil {
			log.Println("Poster Upload Error:", e.Error())
			c.JSON(http.StatusInternalServerError, dto.Response{
				Msg:     "Internal Server Error",
				Success: false,
				Error:   "failed to save poster",
				Data:    nil,
			})
			return
		}
		posterPath := fmt.Sprintf("/movie/%s", filename)
		req.PosterUrl = &posterPath
	}

	if req.Backdrop != nil {
		ext := path.Ext(req.Backdrop.Filename)
		re := regexp.MustCompile("^[.](jpg|png|jpeg)$")
		if !re.MatchString(ext) {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     "Invalid image extension",
				Error:   "Bad Request",
				Success: false,
				Data:    nil,
			})
			return
		}

		filename := fmt.Sprintf("%d_backdrop_%s", time.Now().UnixNano(), req.Backdrop.Filename)
		if e := c.SaveUploadedFile(req.Backdrop, filepath.Join("public", "movie", filename)); e != nil {
			log.Println("Backdrop Upload Error:", e.Error())
			c.JSON(http.StatusInternalServerError, dto.Response{
				Msg:     "Internal Server Error",
				Success: false,
				Error:   "failed to save backdrop",
				Data:    nil,
			})
			return
		}
		backdropPath := fmt.Sprintf("/movie/%s", filename)
		req.BackdropUrl = &backdropPath
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

// CreateMovieAdmin godoc
// @Summary      Create a movie
// @Description  Create a new movie (Requires admin token)
// @Tags         admin
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        title             formData  string  true   "Movie Title"
// @Param        synopsis          formData  string  true   "Movie Synopsis"
// @Param        duration          formData  int     true   "Movie Duration"
// @Param        release_date      formData  string  true   "Movie Release Date (YYYY-MM-DD)"
// @Param        director_id       formData  int     false  "Director ID"
// @Param        poster            formData  file    false  "Poster Image"
// @Param        backdrop          formData  file    false  "Backdrop Image"
// @Param        genre_ids         formData  []int   false  "Genre IDs"
// @Param        popularity_score  formData  number  false  "Popularity Score"
// @Success      201               {object}  dto.Response
// @Failure      401               {object}  dto.Response
// @Failure      400               {object}  dto.Response
// @Failure      500               {object}  dto.Response
// @Router       /admin/movies [post]
func (ctrl AdminController) CreateMovieAdmin(c *gin.Context) {
	var req dto.CreateMovieRequest
	if err := c.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	if req.GenreIds != "" {
		genreIdsStr := strings.Split(req.GenreIds, ",")
		for _, idStr := range genreIdsStr {
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				c.JSON(http.StatusBadRequest, dto.Response{
					Msg:     "Bad Request",
					Success: false,
					Error:   "invalid genre_id format: " + idStr,
					Data:    nil,
				})
				return
			}
			req.Genres = append(req.Genres, id)
		}
	}

	if req.Poster != nil {
		ext := path.Ext(req.Poster.Filename)
		re := regexp.MustCompile("^[.](jpg|png|jpeg)$")
		if !re.MatchString(ext) {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     err.ErrInvalidExt.Error(),
				Error:   "Bad Request",
				Success: false,
				Data:    nil,
			})
			return
		}

		filename := fmt.Sprintf("%d_poster_%s", time.Now().UnixNano(), req.Poster.Filename)
		if e := c.SaveUploadedFile(req.Poster, filepath.Join("public", "movie", filename)); e != nil {
			log.Println("Poster Upload Error:", e.Error())
			c.JSON(http.StatusInternalServerError, dto.Response{
				Msg:     "Internal Server Error",
				Success: false,
				Error:   "failed to save poster",
				Data:    nil,
			})
			return
		}
		posterPath := fmt.Sprintf("/movie/%s", filename)
		req.PosterUrl = &posterPath
	}

	if req.Backdrop != nil {
		ext := path.Ext(req.Backdrop.Filename)
		re := regexp.MustCompile("^[.](jpg|png|jpeg)$")
		if !re.MatchString(ext) {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     err.ErrInvalidExt.Error(),
				Error:   "Bad Request",
				Success: false,
				Data:    nil,
			})
			return
		}

		filename := fmt.Sprintf("%d_backdrop_%s", time.Now().UnixNano(), req.Backdrop.Filename)
		if e := c.SaveUploadedFile(req.Backdrop, filepath.Join("public", "movie", filename)); e != nil {
			log.Println("Backdrop Upload Error:", e.Error())
			c.JSON(http.StatusInternalServerError, dto.Response{
				Msg:     "Internal Server Error",
				Success: false,
				Error:   "failed to save backdrop",
				Data:    nil,
			})
			return
		}
		backdropPath := fmt.Sprintf("/movie/%s", filename)
		req.BackdropUrl = &backdropPath
	}

	data, err := ctrl.adminService.CreateMovieAdmin(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Msg:     "Create Movie Success",
		Success: true,
		Data:    []any{data},
	})
}
