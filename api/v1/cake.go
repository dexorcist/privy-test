package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"privy-test/api/http/request"
	"privy-test/enum"
	"privy-test/integration/logging"
	"privy-test/param/cake"
	"privy-test/service"
	"privy-test/utils"
	"strconv"
)

type cakeHandler struct {
	requestParser request.Parser
	logger        logging.Logger
	cakeService   service.CakeService
}

func NewCakeHandler(
	requestParser request.Parser,
	logger logging.Logger,
	cakeService service.CakeService,
) *cakeHandler {
	return &cakeHandler{
		requestParser: requestParser,
		logger:        logger,
		cakeService:   cakeService,
	}
}

func (ah *cakeHandler) extractParam(c echo.Context) (int64, error) {
	cakeIDParam := c.Param("id")
	if len(cakeIDParam) == 0 {
		return 0, c.JSON(http.StatusBadRequest, utils.MultiStringBadError(enum.HTTPErrorBadRequest))
	}

	cakeID, err := strconv.ParseInt(cakeIDParam, 10, 64)
	if err != nil {
		return 0, c.JSON(http.StatusInternalServerError, utils.DefaultMultiInternalError(enum.HTTPErrorInternalServerError))
	}

	return cakeID, nil
}

// GetDetailCake godoc
// @Summary Get Cake Detail
// @Description  Get Cake Detail
// @Tags Cake
// @Accept json
// @Produce json
// @Param id path int true "Cake ID"
// @Success 200 {object} cake.HTTPGetDetailCakeResponse
// @Failure 400 {object} param.CommonErrorResponse
// @Failure 401 {object} param.CommonErrorResponse
// @Failure 409 {object} param.CommonErrorResponse
// @Failure 500 {object} param.CommonErrorResponse
// @Router /cake/{id} [get]
func (ah *cakeHandler) GetDetailCake(c echo.Context) error {
	ctx := c.Request().Context()

	cakeID, err := ah.extractParam(c)
	if err != nil {
		return err
	}

	result, err := ah.cakeService.GetDetail(ctx, cakeID)
	if err != nil {
		return utils.IsMultiStringHTTPError(err, c)
	}
	return c.JSON(http.StatusOK, result)
}

// CreateCake godoc
// @Summary Create Cake
// @Description  Create Cake
// @Tags Cake
// @Accept json
// @Produce json
// @Param data body cake.CreateUpdateRequest true "Request Body"
// @Success 200 {object} cake.HTTPGetDetailCakeResponse
// @Failure 400 {object} param.CommonErrorResponse
// @Failure 401 {object} param.CommonErrorResponse
// @Failure 409 {object} param.CommonErrorResponse
// @Failure 500 {object} param.CommonErrorResponse
// @Router /cake [post]
func (ah *cakeHandler) CreateCake(c echo.Context) error {
	ctx := c.Request().Context()

	form := new(cake.CreateUpdateRequest)
	err := ah.requestParser.Form(c.Request(), form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.MultiStringBadError(enum.HTTPErrorInternalServerError))
	}

	err = form.Validate()
	if err != nil {
		ah.logger.ErrorWithContext(ctx, err, "validation error %s", err.Error())
		return c.JSON(http.StatusBadRequest, err.(*utils.MultiCommonBadError))
	}

	result, err := ah.cakeService.Create(ctx, form)
	if err != nil {
		return utils.IsMultiStringHTTPError(err, c)
	}
	return c.JSON(http.StatusOK, result)
}

// DeleteCake godoc
// @Summary Delete Cake Detail
// @Description  Delete Cake Detail
// @Tags Cake
// @Accept json
// @Produce json
// @Param id path int true "Cake ID"
// @Success 204
// @Failure 400 {object} param.CommonErrorResponse
// @Failure 401 {object} param.CommonErrorResponse
// @Failure 409 {object} param.CommonErrorResponse
// @Failure 500 {object} param.CommonErrorResponse
// @Router /cake/{id} [delete]
func (ah *cakeHandler) DeleteCake(c echo.Context) error {
	ctx := c.Request().Context()

	cakeID, err := ah.extractParam(c)
	if err != nil {
		return err
	}

	err = ah.cakeService.Delete(ctx, cakeID)
	if err != nil {
		return utils.IsMultiStringHTTPError(err, c)
	}
	return c.NoContent(http.StatusNoContent)
}

// UpdateCake godoc
// @Summary Update Cake
// @Description  Update Cake
// @Tags Cake
// @Accept json
// @Produce json
// @Param id path int true "Cake ID"
// @Param data body cake.CreateUpdateRequest true "Request Body"
// @Success 200 {object} cake.HTTPGetDetailCakeResponse
// @Failure 400 {object} param.CommonErrorResponse
// @Failure 401 {object} param.CommonErrorResponse
// @Failure 409 {object} param.CommonErrorResponse
// @Failure 500 {object} param.CommonErrorResponse
// @Router /cake/{id} [patch]
func (ah *cakeHandler) UpdateCake(c echo.Context) error {
	ctx := c.Request().Context()

	cakeID, err := ah.extractParam(c)
	if err != nil {
		return err
	}

	form := new(cake.CreateUpdateRequest)
	err = ah.requestParser.Form(c.Request(), form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.MultiStringBadError(enum.HTTPErrorInternalServerError))
	}

	err = form.Validate()
	if err != nil {
		ah.logger.ErrorWithContext(ctx, err, "validation error %s", err.Error())
		return c.JSON(http.StatusBadRequest, err.(*utils.MultiCommonBadError))
	}

	result, err := ah.cakeService.Update(ctx, cakeID, form)
	if err != nil {
		return utils.IsMultiStringHTTPError(err, c)
	}
	return c.JSON(http.StatusOK, result)
}

// GetList godoc
// @Summary Get List
// @Description  Get List
// @Tags Cake
// @Accept json
// @Produce json
// @Param title query string false "Title"
// @Param description query string false "Description"
// @Param min_rating query number false "Min Rating"
// @Param max_rating query number false "Max Rating"
// @Success 200 {object} cake.HTTPGetListCakeResponse
// @Failure 400 {object} param.CommonErrorResponse
// @Failure 401 {object} param.CommonErrorResponse
// @Failure 409 {object} param.CommonErrorResponse
// @Failure 500 {object} param.CommonErrorResponse
// @Router /cake [get]
func (ah *cakeHandler) GetList(c echo.Context) error {
	ctx := c.Request().Context()

	form := new(cake.FindAllRequest)
	err := ah.requestParser.Query(c.Request(), form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.MultiStringBadError(enum.HTTPErrorInternalServerError))
	}

	err = form.Validate()
	if err != nil {
		ah.logger.ErrorWithContext(ctx, err, "validation error %s", err.Error())
		return c.JSON(http.StatusBadRequest, err.(*utils.MultiCommonBadError))
	}

	result, err := ah.cakeService.GetList(ctx, form)
	if err != nil {
		return utils.IsMultiStringHTTPError(err, c)
	}
	return c.JSON(http.StatusOK, result)
}
