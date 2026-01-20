package controller

import (
	"net/http"
	"strconv"
	"strings"

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

// GetUpcomingMovies godoc
// @Summary      Get upcoming movies
// @Description  Get a list of upcoming movies
// @Tags         movies
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /movies/upcoming [get]
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
	if len(data) == 0 {
		c.JSON(http.StatusNotFound, dto.Response{
			Msg:     "Data not found",
			Success: true,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Get Upcoming Movies Success",
		Success: true,
		Data:    data,
	})
}

// GetPopularMovie godoc
// @Summary      Get popular movies
// @Description  Get a list of popular movies
// @Tags         movies
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /movies/popular [get]
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
	if len(data) == 0 {
		c.JSON(http.StatusFound, dto.Response{
			Msg:     "data not found",
			Success: true,
			Data:    nil,
		})
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Get Upcoming Movies Success",
		Success: true,
		Data:    data,
	})
}

// GetMovieWithFilter godoc
// @Summary      Filter movies
// @Description  Get movies with search, genre filter, and pagination
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        search    query     string   false  "Search by title"
// @Param        genre_id  query     []int    false  "Filter by genre ID"
// @Param        page      query     int      false  "Page number (default: 1)"
// @Success      200       {object}  dto.Response
// @Failure      400       {object}  dto.Response
// @Failure      500       {object}  dto.Response
// @Router       /movies [get]
func (ctrl MovieController) GetMovieWithFilter(c *gin.Context) {
	var search *string
	if s := c.Query("search"); s != "" {
		search = &s
	}

	var genreIds []int
	queryGenres := c.QueryArray("genre_id")
	for _, q := range queryGenres {
		parts := strings.Split(q, ",")
		for _, p := range parts {
			if id, err := strconv.Atoi(strings.TrimSpace(p)); err == nil {
				genreIds = append(genreIds, id)
			}
		}
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "16"))

	data, meta, err := ctrl.movieService.GetMovieWithFilter(c.Request.Context(), search, genreIds, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   err.Error(),
			Data:    []any{},
		})
		return
	}
	if len(data) == 0 {
		c.JSON(http.StatusNotFound, dto.Response{
			Msg:     "Data not found",
			Success: true,
			Data:    nil,
			Meta:    meta,
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

// GetMovieDetail godoc
// @Summary      Get movie detail
// @Description  Get detailed information about a specific movie
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Movie ID"
// @Success      200  {object}  dto.Response
// @Failure      500  {object}  dto.Response
// @Router       /movies/detail/{id} [get]
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
	if len(data) == 0 {
		c.JSON(http.StatusNotFound, dto.Response{
			Msg:     "Data not found",
			Success: true,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Get Detail Movie Succes",
		Success: true,
		Data:    data,
	})
}
