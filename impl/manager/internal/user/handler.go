package user

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/auth"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/company"
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

	if _, err := h.CompanyUsecase.GetByID(ctx, req.CompanyID); err != nil {
		err := fmt.Errorf("invalid companyID %s, err: %s", req.CompanyID, err)
		h.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	userReq := User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CompanyID: req.CompanyID,
	}

	user, err := h.Usecase.Create(ctx, userReq)
	if err != nil {
		h.Logger.Errorf("error when creating users err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, handler.SuccessResponse{Data: user})
}

func (h *Handler) V1GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		h.Logger.Errorf("error when parsing id: %s, err: %s", idParam, err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	user, err := h.Usecase.GetByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found id: %s, err: %s", id, err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting user with id: %s, err: %s", id, err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: user})
}

func (h *Handler) V1Login(c echo.Context) error {
	req := LoginRequest{}
	ctx := c.Request().Context()

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

	user, err := h.Usecase.Login(ctx, req.Email, req.Password)
	if err != nil {
		h.Logger.Errorf("error when login users with email: %s, err: %s", req.Email, err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	myClaims := auth.MyClaims{
		UserID:    user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CompanyID: user.CompanyID,
	}

	accessToken, refreshToken, err := auth.CreateAndSignToken(myClaims, auth.Authentication)
	if err != nil {
		h.Logger.Errorf("error when creating token err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "server error"})
	}

	expiredAt := time.Now().Add(auth.LoginExpiration)
	atCookie := auth.CreateCookie("accessToken", accessToken, 60*60)
	rtCookie := auth.CreateCookie("refreshToken", refreshToken, 24*60*60)

	c.SetCookie(atCookie)
	c.SetCookie(rtCookie)

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt,
	}})
}

func (h *Handler) V1Refresh(c echo.Context) error {
	var refreshToken = ""
	cookies := c.Cookies()

	for _, ck := range cookies {
		if ck.Name == "refreshToken" {
			refreshToken = ck.Value
		}
	}

	if refreshToken == "" {
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "invalid token"})
	}

	rt, err := auth.ValidateToken(refreshToken)
	if err != nil {
		h.Logger.Errorf("error when validating refresh token: %v, err: %s", rt, err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "invalid token"})
	}

	claims, ok := rt.Claims.(*auth.JwtClaims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "server error"})

	}

	accessToken, refreshToken, err := auth.GeneratePairToken(*claims)
	if err != nil {
		h.Logger.Errorf("error when creating token err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "server error"})
	}

	expiredAt := time.Now().Add(auth.LoginExpiration)
	atCookie := auth.CreateCookie("accessToken", accessToken, 60*60)
	rtCookie := auth.CreateCookie("refreshToken", refreshToken, 24*60*60)

	c.SetCookie(atCookie)
	c.SetCookie(rtCookie)

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt,
	}})
}

func (h *Handler) V1GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	companyID, ok := c.Get("companyID").(uuid.UUID)

	if !ok {
		h.Logger.Errorf("error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	user, err := h.Usecase.GetAllByCompanyID(ctx, companyID)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found err: %s", err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting users with err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: user})
}

func (h *Handler) V1AdminGetAll(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := h.Usecase.GetAll(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found err: %s", err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting users with err: %s", err)
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
