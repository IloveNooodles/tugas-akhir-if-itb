package user

import (
	"fmt"
	"net/http"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/auth"
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

	userReq := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
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

func (h *Handler) V1Login(c echo.Context) error {
	req := LoginRequest{}
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

	user, err := h.Usecase.Login(ctx, req.Email, req.Password)
	if err != nil {
		h.Logger.Errorf("error when login users with email: %s, err: %s", req.Email, err)
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	myClaims := auth.MyClaims{
		UserID: user.ID,
		Email:  user.Email,
		Name:   user.Name,
	}

	token, err := auth.CreateAndSignToken(myClaims, auth.Authentication)
	if err != nil {
		h.Logger.Errorf("error when creating token err: %s", err)
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "server error"})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResponse{Data: token})
}
