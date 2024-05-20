package device

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/company"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/errx"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/handler"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/validatorx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Logger         *logrus.Logger
	Usecase        Usecase
	CompanyUsecase company.Usecase
}

func NewHandler(l *logrus.Logger, u Usecase, cu company.Usecase) Handler {
	return Handler{
		Logger:         l,
		Usecase:        u,
		CompanyUsecase: cu,
	}
}

func (h *Handler) V1Create(c echo.Context) error {
	req := CreateRequest{}
	ctx := c.Request().Context()
	companyID, ok := c.Get("companyID").(uuid.UUID)

	if !ok {
		h.Logger.Errorf("error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	if err := c.Bind(&req); err != nil {
		err := fmt.Errorf("error when receiving request err: %s", err)
		h.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	v := validatorx.New()
	if err := v.StructCtx(ctx, &req); err != nil {
		err := fmt.Errorf("error when validating request: %v, err: %s", req, err)
		h.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	if _, err := h.CompanyUsecase.GetByID(ctx, companyID); err != nil {
		err := fmt.Errorf("invalid companyID %s, err: %s", companyID, err)
		h.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	var nodeName = req.Name
	if req.NodeName != "" {
		nodeName = req.NodeName
	}

	deviceReq := Device{
		Name:      req.Name,
		Type:      req.Type,
		Labels:    req.Labels,
		CompanyID: companyID,
		NodeName:  nodeName,
	}

	device, err := h.Usecase.Create(ctx, deviceReq)
	if errx.IsNodeNotFound(err) {
		h.Logger.Errorf("device: error when creating devices err: %s", err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "node not found"})
	}

	if errx.IsClusterDown(err) {
		h.Logger.Errorf("device: error when creating devices err: %s", errx.ErrClusterDown)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: errx.ErrClusterDown.Error()})
	}

	if errx.IsDuplicateDatabase(err) {
		h.Logger.Errorf("device: error when creating devices err: %s", err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "duplicate node name"})
	}

	if err != nil {
		h.Logger.Errorf("device: error when creating devices err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, handler.SuccessResponse{Data: device})
}

func (h *Handler) V1GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")
	companyID, ok := c.Get("companyID").(uuid.UUID)

	if !ok {
		h.Logger.Errorf("error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		h.Logger.Errorf("error when parsing id: %s, err: %s", idParam, err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	device, err := h.Usecase.GetByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found id: %s, err: %s", id, err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if device.CompanyID != companyID {
		return c.JSON(http.StatusForbidden, handler.ErrorResponse{Message: "forbidden"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting device with id: %s, err: %s", id, err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: device})
}

func (h *Handler) V1AdminGetAll(c echo.Context) error {
	ctx := c.Request().Context()

	devices, err := h.Usecase.GetAll(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found err: %s", err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting device with err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: devices})
}

func (h *Handler) V1GetGroupByDeviceID(c echo.Context) error {
	ctx := c.Request().Context()
	ctxCompanyID, ok := c.Get("companyID").(uuid.UUID)
	idParam := c.Param("id")

	deviceID, err := uuid.Parse(idParam)
	if err != nil {
		h.Logger.Errorf("error when parsing id: %s, err: %s", idParam, err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	if !ok {
		h.Logger.Errorf("error when converting company id to uuid")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	devices, err := h.Usecase.GetGroups(ctx, ctxCompanyID, deviceID)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found err: %s", err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting device with err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: devices})
}

func (h *Handler) V1GetAllByCompanyID(c echo.Context) error {
	ctx := c.Request().Context()
	companyID, ok := c.Get("companyID").(uuid.UUID)
	if !ok {
		h.Logger.Errorf("error when converting company id to uuid")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	devices, err := h.Usecase.GetAllByCompanyID(ctx, companyID)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found err: %s", err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting device with err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: devices})
}

func (h *Handler) V1Delete(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		h.Logger.Errorf("error when parsing id: %s, err: %s", idParam, err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	err = h.Usecase.Delete(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found id: %s, err: %s", id, err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting user with id: %s, err: %s", id, err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusNoContent, handler.SuccessResponse{})
}
