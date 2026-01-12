package handler

import (
	"compro/config"
	"compro/internal/adapter/handler/request"
	"compro/internal/adapter/handler/response"
	"compro/internal/core/domain/entity"
	"compro/internal/core/service"
	"compro/utils/conv"
	"compro/utils/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type IHeroSectionHandler interface {
	CreateHeroSection(c echo.Context) error
}

type heroSectionHandler struct {
	heroSectionService service.IHeroSectionService
}

// CreateHeroSection implements [IHeroSectionHandler].
func (h *heroSectionHandler) CreateHeroSection(c echo.Context) error {
	var (
		req       = request.HeroSectionRequest{}
		resp      = response.DefaultSuccessResponse{}
		respError = response.ErrorResponseDefault{}
		ctx       = c.Request().Context()
	)

	user := conv.GetUserIDByContext(c)
	if user == 0 {
		log.Errorf("[HANDLER] CreateHeroSection - 1: Unauthorized")
		respError.Meta.Message = "Unauthorized"
		respError.Meta.Status = false
		return c.JSON(http.StatusUnauthorized, respError)
	}

	if err = c.Bind(&req); err != nil {
		log.Errorf("[HANDLER] CreateHeroSection - 2: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusUnprocessableEntity, respError)
	}

	if err = c.Validate(req); err != nil {
		log.Errorf("[HANDLER] CreateHeroSection - 3: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(http.StatusBadRequest, respError)
	}

	reqEntity := entity.HeroSectionEntity{
		Heading:    req.Heading,
		SubHeading: req.SubHeading,
		PathVideo:  req.PathVideo,
		Banner:     req.Banner,
		
	}

	err = h.heroSectionService.CreateHeroSection(ctx, reqEntity)
	if err != nil {
		log.Errorf("[HANDLER] CreateHeroSection - 4: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(conv.SetHTTPStatusCode(err), respError)
	}

	resp.Meta.Message = "Success create hero section"
	resp.Meta.Status = true
	resp.Data = nil
	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

func NewHeroSectionHandler(c *echo.Echo, cfg *config.Config, heroSectionService service.IHeroSectionService) IHeroSectionHandler {
	hero := &heroSectionHandler{
		heroSectionService: heroSectionService,
	}

	mid := middleware.NewMiddleware(cfg)
	heroApp := c.Group("/hero-sections")

	adminApp := heroApp.Group("/admin", mid.CheckToken())
	adminApp.POST("", hero.CreateHeroSection)

	return hero
}
