package company

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/errx"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/handler"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/validatorx"
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
		err := fmt.Errorf("company: invalid request err: %s", err)
		h.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	v := validatorx.New()
	if err := v.StructCtx(ctx, &req); err != nil {
		err := fmt.Errorf("company: error when validating request: %v, err: %s", req, err)
		h.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	userReq := Company{
		Name:        req.Name,
		ClusterName: req.ClusterName,
	}

	user, err := h.Usecase.Create(ctx, userReq)
	if errors.Is(err, ErrClusterNotAvailable) {
		h.Logger.Errorf("company: error when creating err: %s", err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	if errx.IsDuplicateDatabase(err) {
		h.Logger.Errorf("company: error when creating err: %s", err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "duplicate combination name and cluster_name"})
	}

	if err != nil {
		h.Logger.Errorf("company: error when creating err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	return c.JSON(http.StatusCreated, handler.SuccessResponse{Data: user})
}

func (h *Handler) V1GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	companyID, ok := c.Get("companyID").(uuid.UUID)

	if !ok {
		h.Logger.Errorf("company: error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	user, err := h.Usecase.GetByID(ctx, companyID)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("company: no rows found id: %s, err: %s", companyID, err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "not found"})
	}

	if err != nil {
		h.Logger.Errorf("company: error when getting user with id: %s, err: %s", companyID, err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: user})
}

func (h *Handler) V1AdminGetAll(c echo.Context) error {
	ctx := c.Request().Context()
	companies, err := h.Usecase.GetAll(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no result err: %s", err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting user with err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: companies})
}

func (h *Handler) V1GetCompanyAndLoggedInUser(c echo.Context) error {
	ctx := c.Request().Context()
	companyID, companyOk := c.Get("companyID").(uuid.UUID)
	userID, userOK := c.Get("userID").(uuid.UUID)

	if !userOK || !companyOk {
		h.Logger.Errorf("error when converting context info to string, companyID: %s, userID: %s", companyID, userID)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	user, err := h.Usecase.GetCompanyAndLoggedInUser(ctx, companyID, userID)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found id: %s, err: %s", companyID, err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting user with id: %s, err: %s", companyID, err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: user})
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

func (h *Handler) V1CheckClusterStatus(c echo.Context) error {
	ctx := c.Request().Context()

	err := h.Usecase.CheckClusterStatus(ctx)
	if err != nil {
		h.Logger.Errorf("company: cluster server check: %s", err)
		return c.JSON(http.StatusOK, handler.SuccessResponse{Data: "Server is down, please try again later"})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: "Server is healthy"})
}
