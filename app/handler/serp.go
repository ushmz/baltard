package handler

import (
	"errors"
	"fmt"
	"net/http"

	"ratri/domain/model"
	"ratri/usecase"

	"github.com/labstack/echo/v4"
)

// Serp : Implemention of SERP handler
type Serp struct {
	usecase usecase.SerpUsecase
}

// NewSerpHandler : Return new SERP handler
func NewSerpHandler(serp usecase.SerpUsecase) *Serp {
	return &Serp{usecase: serp}
}

// FetchSERPParam : Request parameters for fetch search result pages
type FetchSERPParam struct {
	ID     int `json:"id" param:"id"`
	Offset int `json:"offset" query:"offset"`
	Top    int `json:"top" query:"top"`
}

// FetchSerpWithRatioByID : Return search result pages with similarweb information (such as icon)
// @Id fetch_serp_with_ratio_by_id
// @Summary Returns json which have a list of search results with data for RatioUI.
// @Description Returns json which have a list of search results with tracked pages of Similarweb top 2000 and its category.
// @Accept json
// @Produce json
// @Param taskId path int true "Task ID"
// @Param offset query int true "Number of offset to display"
// @Param top query int false "Number of categories to display"
// @Success 200 {object} []model.SerpWithRatio
// @Failure 400
// @Failure 500
// @Router /v1/serp/{taskId}/ratio [GET]
func (s *Serp) FetchSerpWithRatioByID(c echo.Context) error {
	if s == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, model.ErrNilReceiver)
	}

	p := FetchSERPParam{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			ErrWithMessage{
				error: fmt.Errorf("Invalid request body: %w", err),
				Why:   "Invalid request body",
			},
		)
	}

	if p.Top == 0 {
		p.Top = 3
	}

	serp, err := s.usecase.FetchSerpWithRatio(p.ID, p.Offset, p.Top)
	if err != nil {
		if errors.Is(err, model.ErrNoSuchData) {
			return echo.NewHTTPError(
				http.StatusNotFound,
				ErrWithMessage{
					error: fmt.Errorf("Page nout found: %w", err),
					Why:   "Page not found",
				},
			)
		}
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			ErrWithMessage{
				error: fmt.Errorf("Try to get search result: %w", err),
				Why:   "Request failed",
			},
		)
	}

	return c.JSON(http.StatusOK, serp)
}

// FetchSerpWithIconByID : Return search result pages with similarweb information (such as icon)
// @Id fetch_serp_with_icon_by_id
// @Summary Returns json which have a list of search results with data for IconUI.
// @Description Returns json which have a list of search results with tracked pages of Similarweb top 2000 pages and its favicon URL.
// @Accept json
// @Produce json
// @Param taskId path int true "Task ID"
// @Param offset query int true "Number of offset to display"
// @Param top query int false "Number of icons to display"
// @Success 200 {object} []model.SerpWithIcon
// @Failure 400
// @Failure 500
// @Router /v1/serp/{taskId}/icon [GET]
func (s *Serp) FetchSerpWithIconByID(c echo.Context) error {
	if s == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, model.ErrNilReceiver)
	}

	p := FetchSERPParam{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			ErrWithMessage{
				error: fmt.Errorf("Invalid request body: %w", err),
				Why:   "Invalid request body",
			},
		)
	}

	if p.Top == 0 {
		p.Top = 3
	}

	serp, err := s.usecase.FetchSerpWithIcon(p.ID, p.Offset, p.Top)
	if err != nil {
		if errors.Is(err, model.ErrNoSuchData) {
			return echo.NewHTTPError(
				http.StatusNotFound,
				ErrWithMessage{
					error: fmt.Errorf("Page nout found: %w", err),
					Why:   "Page not found",
				},
			)
		}
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			ErrWithMessage{
				error: fmt.Errorf("Try to get search result: %w", err),
				Why:   "Request failed",
			},
		)
	}

	return c.JSON(http.StatusOK, serp)
}

// FetchSerpByID : Return search result pages by task id .
// @Id fetch_serp_by_id
// @Summary Returns json which have a list of search results.
// @Description Returns json which have a list of search results with no additional data.
// @Accept json
// @Produce json
// @Param taskId path int true "Task ID"
// @Param offset query int true "Number of offset to display"
// @Success 200 {object} []model.SearchPage
// @Failure 400
// @Failure 500
// @Router /v1/serp/{taskId} [GET]
func (s *Serp) FetchSerpByID(c echo.Context) error {
	if s == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, model.ErrNilReceiver)
	}

	p := FetchSERPParam{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			ErrWithMessage{
				error: fmt.Errorf("Invalid request body: %w", err),
				Why:   "Invalid request body",
			},
		)
	}

	serp, err := s.usecase.FetchSerp(p.ID, p.Offset)
	if err != nil {
		if errors.Is(err, model.ErrNoSuchData) {
			return echo.NewHTTPError(
				http.StatusNotFound,
				ErrWithMessage{
					error: fmt.Errorf("Page nout found: %w", err),
					Why:   "Page not found",
				},
			)
		}
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			ErrWithMessage{
				error: fmt.Errorf("Try to get search result: %w", err),
				Why:   "Request failed",
			},
		)
	}

	return c.JSON(http.StatusOK, serp)
}
