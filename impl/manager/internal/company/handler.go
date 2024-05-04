package company

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Logger  *logrus.Logger
	Usecase Usecase
}

func NewHandler(l *logrus.Logger, u Usecase) Handler {
	return Handler{
		Logger:  l,
		Usecase: u,
	}
}

func (h *Handler) V1Create(c echo.Context) error {
	req := CreateRequest{}
	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		err := fmt.Errorf("error when receiving request err: %s", err)
		h.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	v := validator.New()
	if err := v.StructCtx(ctx, &req); err != nil {
		err := fmt.Errorf("error when validating request: %v, err: %s", req, err)
		h.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	userReq := Company{
		Name:        req.Name,
		ClusterName: req.ClusterName,
	}

	user, err := h.Usecase.Create(ctx, userReq)
	if err != nil {
		h.Logger.Errorf("error when creating users err: %s", err)
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResponse{Data: user})
}

func (h *Handler) V1GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	companyID, ok := c.Get("companyID").(uuid.UUID)

	if !ok {
		h.Logger.Errorf("error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "internal server error"})
	}

	user, err := h.Usecase.GetByID(ctx, companyID)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found id: %s, err: %s", companyID, err)
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting user with id: %s, err: %s", companyID, err)
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{Data: user})
}

func (h *Handler) V1AdminGetAll(c echo.Context) error {
	ctx := c.Request().Context()
	companies, err := h.Usecase.GetAll(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no result err: %s", err)
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting user with err: %s", err)
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{Data: companies})
}
