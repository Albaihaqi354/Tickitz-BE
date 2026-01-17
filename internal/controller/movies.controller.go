package controller

import (
	"net/http"
	"strconv"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/service"
	"github.com/gin-gonic/gin"
)

type MovieController struct {
	movieService *service.MovieService
}

func NewMovieController(movieService *service.MovieService) *MovieController {
	return &MovieController{
		movieService: movieService,
	}
}

func (ctrl MovieController) GetUpcomingMovies(c *gin.Context) {
	data, err := ctrl.movieService.GetUpcomingMovies(c.Request.Context())
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
		Msg:     "Get Upcoming Movies Success",
		Success: true,
		Data:    data,
	})
}

func (ctrl MovieController) GetPopularMovie(c *gin.Context) {
	data, err := ctrl.movieService.GetPopularMovie(c.Request.Context())
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
		Msg:     "Get Upcoming Movies Success",
		Success: true,
		Data:    data,
	})
}

func (ctrl MovieController) GetMovieWithFilter(c *gin.Context) {
	var search *string
	if s := c.Query("search"); s != "" {
		search = &s
	}

	var genreId *int
	if genre := c.Query("genre_id"); genre != "" {
		id, err := strconv.Atoi(genre)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.Response{
				Msg:     "Invalid genre_id parameter",
				Success: false,
				Error:   err.Error(),
				Data:    []any{},
			})
			return
		}
		genreId = &id
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "16"))

	data, meta, err := ctrl.movieService.GetMovieWithFilter(c.Request.Context(), search, genreId, page, limit)
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
		Msg:     "Get Filter Movie Success",
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

func (ctr MovieController) GetMovieDetail(c *gin.Context) {
	idParam := c.Param("id")

	movieId, err := strconv.Atoi(idParam)
	data, err := ctr.movieService.GetMovieDetail(c.Request.Context(), movieId)
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
		Msg:     "Get Detail Movie Succes",
		Success: true,
		Data:    data,
	})
}
