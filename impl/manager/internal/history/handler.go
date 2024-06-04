package history

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/company"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/groups"
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
	GroupUsecase   groups.Usecase
}

func NewHandler(l *logrus.Logger, u Usecase, cu company.Usecase, gu groups.Usecase) Handler {
	return Handler{
		Logger:         l,
		Usecase:        u,
		CompanyUsecase: cu,
		GroupUsecase:   gu,
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
		h.Logger.Errorf("error when receiving request err: %s", err)
		return err
	}

	v := validatorx.New()
	if err := v.StructCtx(ctx, &req); err != nil {
		h.Logger.Errorf("error when validating request: %v, err: %s", req, err)
		return err
	}

	if _, err := h.CompanyUsecase.GetByID(ctx, companyID); err != nil {
		err := fmt.Errorf("invalid companyID %s, err: %s", companyID, err)
		h.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	historyRequest := Histories{
		DeviceID:     req.DeviceID,
		RepositoryID: req.RepositoryID,
		DeploymentID: req.DeploymentID,
		Status:       req.Status,
	}

	user, err := h.Usecase.Create(ctx, historyRequest)
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

	history, err := h.Usecase.GetByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found id: %s, err: %s", id, err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	companyID, ok := c.Get("companyID").(uuid.UUID)

	if !ok {
		h.Logger.Errorf("error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	if history.CompanyID != companyID {
		return c.JSON(http.StatusForbidden, handler.ErrorResponse{Message: "forbidden"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting histories with id: %s, err: %s", id, err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: history})
}

func (h *Handler) V1AdminGetAll(c echo.Context) error {
	ctx := c.Request().Context()
	histories, err := h.Usecase.GetAll(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found err: %s", err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting histories with err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: histories})
}

func (h *Handler) V1GetAllByCompanyID(c echo.Context) error {
	ctx := c.Request().Context()
	companyID, ok := c.Get("companyID").(uuid.UUID)
	if !ok {
		h.Logger.Errorf("error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	p := GetAllParams{
		CompanyID: companyID,
		DeviceID:  make(uuid.UUIDs, 0),
	}

	var groupID = uuid.Nil
	var deviceID = uuid.Nil
	var deploymentID = uuid.Nil
	var repositoryID = uuid.Nil

	var parseError error
	qGroupID := c.QueryParam("group_id")
	if qGroupID != "" {
		groupID, parseError = uuid.Parse(qGroupID)
		if parseError != nil {
			h.Logger.Errorf("error when converting group id to uuid: %s, err: %s", qGroupID, parseError)
			return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "invalid group id"})
		}

		deviceList, err := h.GroupUsecase.GetDevices(ctx, companyID, groupID)
		if err != nil {
			h.Logger.Errorf("error when getting devices id from groupid %s, err: %s", groupID, err)
			return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "invalid group id"})
		}

		for _, d := range deviceList {
			p.DeviceID = append(p.DeviceID, d.DeviceID)
		}
	}

	qDeviceID := c.QueryParam("device_id")
	if qDeviceID != "" {
		deviceID, parseError = uuid.Parse(qDeviceID)
		if parseError != nil {
			h.Logger.Errorf("error when converting device id to uuid: %s, err: %s", qDeviceID, parseError)
			return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "internal device id"})
		}

		p.DeviceID = append(p.DeviceID, deviceID)
	}

	qDeploymentID := c.QueryParam("deployment_id")
	if qDeploymentID != "" {
		deploymentID, parseError = uuid.Parse(qDeploymentID)
		if parseError != nil {
			h.Logger.Errorf("error when converting device id to uuid: %s, err: %s", qDeploymentID, parseError)
			return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "internal deployment id"})
		}

		p.DeploymentID = deploymentID
	}

	qRepositoryID := c.QueryParam("repository_id")
	if qRepositoryID != "" {
		repositoryID, parseError = uuid.Parse(qRepositoryID)
		if parseError != nil {
			h.Logger.Errorf("error when converting device id to uuid: %s, err: %s", qRepositoryID, parseError)
			return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "internal deployment id"})
		}

		p.RepositoryID = repositoryID
	}

	histories, err := h.Usecase.GetAllByCompanyID(ctx, p)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found err: %s", err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting histories with err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: histories})
}
