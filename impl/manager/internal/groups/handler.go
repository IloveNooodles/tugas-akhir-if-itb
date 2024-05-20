package groups

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

	groupRequest := Group{
		Name:      req.Name,
		CompanyID: companyID,
	}

	group, err := h.Usecase.Create(ctx, groupRequest)
	if errx.IsDuplicateDatabase(err) {
		h.Logger.Errorf("device: error when creating devices err: %s", err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: fmt.Sprintf("name %s already exists, please use another name", req.Name)})
	}

	if err != nil {
		h.Logger.Errorf("error when creating users err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, handler.SuccessResponse{Data: group})
}

func (h *Handler) V1GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		h.Logger.Errorf("error when parsing id: %s, err: %s", idParam, err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	group, err := h.Usecase.GetByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found id: %s, err: %s", id, err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	companyID, ok := c.Get("companyID").(uuid.UUID)

	if !ok {
		h.Logger.Errorf("error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	if group.CompanyID != companyID {
		return c.JSON(http.StatusForbidden, handler.ErrorResponse{Message: "forbidden"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting groups with id: %s, err: %s", id, err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: group})
}

func (h *Handler) V1AdminGetAll(c echo.Context) error {
	ctx := c.Request().Context()
	groups, err := h.Usecase.GetAll(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found err: %s", err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting groups with err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: groups})
}

func (h *Handler) V1GetAllByCompanyID(c echo.Context) error {
	ctx := c.Request().Context()
	companyID, ok := c.Get("companyID").(uuid.UUID)

	if !ok {
		h.Logger.Errorf("error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	groups, err := h.Usecase.GetAllByCompanyID(ctx, companyID)

	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found err: %s", err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting groups with err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: groups})
}

func (h *Handler) V1GetByDeviceID(c echo.Context) error {
	ctx := c.Request().Context()
	companyID, ok := c.Get("companyID").(uuid.UUID)
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		h.Logger.Errorf("error when parsing id: %s, err: %s", idParam, err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	if !ok {
		h.Logger.Errorf("error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	devices, err := h.Usecase.GetDevices(ctx, companyID, id)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found err: %s", err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting devices with err: %s", err)
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
