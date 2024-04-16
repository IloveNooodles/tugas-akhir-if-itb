package device

import (
	"fmt"
	"net/http"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/company"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/dto"
	"github.com/go-playground/validator/v10"
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

	if _, err := h.CompanyUsecase.GetByID(ctx, req.CompanyID); err != nil {
		err := fmt.Errorf("invalid companyID %s, err: %s", req.CompanyID, err)
		h.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	deviceReq := Device{
		Name:       req.Name,
		Type:       req.Type,
		Attributes: req.Attributes,
		CompanyID:  req.CompanyID,
	}

	user, err := h.Usecase.Create(ctx, deviceReq)
	if err != nil {
		h.Logger.Errorf("error when creating users err: %s", err)
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResponse{Data: user})
}

func (h *Handler) V1GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		h.Logger.Errorf("error when parsing id: %s, err: %s", idParam, err)
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	user, err := h.Usecase.GetByID(ctx, id)
	if err != nil {
		h.Logger.Errorf("error when getting user with id: %s, err: %s", id, err)
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse{Data: user})
}
