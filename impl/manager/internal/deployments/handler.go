package deployments

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/company"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/config"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/device"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/errx"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/groups"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/handler"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/history"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/util"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/validatorx"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Logger         *logrus.Logger
	Usecase        Usecase
	CompanyUsecase company.Usecase
	HistoryUsecase history.Usecase
	DeviceUsecase  device.Usecase
	GroupUsecase   groups.Usecase
	Config         config.Config
}

func NewHandler(l *logrus.Logger, u Usecase, cu company.Usecase, h history.Usecase, du device.Usecase, gu groups.Usecase, cfg config.Config) Handler {
	return Handler{
		Logger:         l,
		Usecase:        u,
		CompanyUsecase: cu,
		HistoryUsecase: h,
		DeviceUsecase:  du,
		GroupUsecase:   gu,
		Config:         cfg,
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

	deploymentRequest := Deployment{
		Name:         req.Name,
		RepositoryID: req.RepositoryID,
		Version:      req.Version,
		Target:       req.Target,
		CompanyID:    companyID,
	}

	user, err := h.Usecase.Create(ctx, deploymentRequest)
	if errx.IsDuplicateDatabase(err) {
		h.Logger.Errorf("device: error when creating deployments err: %s", err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "duplicate deployments"})
	}

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

	deployment, err := h.Usecase.GetByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found id: %s, err: %s", id, err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	companyID, ok := c.Get("companyID").(uuid.UUID)

	if !ok {
		h.Logger.Errorf("error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	if deployment.CompanyID != companyID {
		return c.JSON(http.StatusForbidden, handler.ErrorResponse{Message: "forbidden"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting deployments with id: %s, err: %s", id, err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: deployment})
}

func (h *Handler) V1AdminGetAll(c echo.Context) error {
	ctx := c.Request().Context()
	deployments, err := h.Usecase.GetAll(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found err: %s", err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting deployments with err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: deployments})
}

func (h *Handler) V1GetAllByCompanyID(c echo.Context) error {
	ctx := c.Request().Context()
	companyID, ok := c.Get("companyID").(uuid.UUID)

	if !ok {
		h.Logger.Errorf("error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	deployments, err := h.Usecase.Repo.GetAllByCompanyID(ctx, companyID)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found err: %s", err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when getting deployments with err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse{Data: deployments})
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

func (h *Handler) V1Deploy(c echo.Context) error {
	req := DeploymentRequest{}
	ctx := c.Request().Context()
	companyID, ok := c.Get("companyID").(uuid.UUID)

	if !ok {
		h.Logger.Errorf("error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	clusterName, ok := c.Get("clusterName").(string)

	if !ok {
		h.Logger.Errorf("company: cluster name not found %s", clusterName)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "invalid cluster name"})
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

	// Get all available deployments from IDS
	deployments, err := h.Usecase.GetDeploymentWithRepositoryByIDs(ctx, companyID, req.DeploymentIDs)
	if errors.Is(err, sql.ErrNoRows) {
		err := fmt.Errorf("deployments with ids %v not found, err: %s", req.DeploymentIDs, err)
		h.Logger.Error(err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "not found"})
	}

	if err != nil {
		err := fmt.Errorf("error when getting deployments with ids %v, err: %s", req.DeploymentIDs, err)
		h.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	devices := make([]device.Device, 0)
	switch req.Type {
	case "CUSTOM":
		switch req.Custom.Kind {
		case "DEVICE":
			deviceList, err := h.DeviceUsecase.GetAllByIDs(ctx, companyID, req.Custom.ListId)
			if err != nil {
				h.Logger.Errorf("error when getting deployment ids %v, err: %s", req.Custom.ListId.Strings(), err)
				return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
			}

			devices = append(devices, deviceList...)
		case "GROUP":
			deviceList, err := h.DeviceUsecase.GetAllByGroupIDs(ctx, companyID, req.Custom.ListId)
			if err != nil {
				h.Logger.Errorf("error when getting deployment ids %v, err: %s", req.Custom.ListId.Strings(), err)
				return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
			}

			devices = append(devices, deviceList...)
		}

		// Label nodes after get all possible from devicesList for custom type
		for _, deployment := range deployments {
			for _, device := range devices {
				labels := util.SplitByComma(deployment.Target)
				for _, label := range labels {
					k, v, err := util.SplitByEqual(label)
					if err != nil {
						h.Logger.Errorf("error when splitting by labels device k, v, err: %s %s %s", k, v, err)
						return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
					}

					err = h.Usecase.kc.LabelNodes(ctx, clusterName, device.NodeName, k, v)
					if err != nil {
						h.Logger.Errorf("error when labels device device: %v, %s, %s, err:  %s", device, k, v, err)
						return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
					}
				}
			}
		}
	case "TARGET":
		for _, d := range deployments {
			devicesList, err := h.DeviceUsecase.GetAllByLabels(ctx, companyID, d.Target)
			if err != nil {
				h.Logger.Errorf("error when getting target %s, err: %s", d.Target, err)
				return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
			}

			devices = append(devices, devicesList...)
		}
	}

	listDeployment, errs := h.Usecase.Deploy(ctx, deployments, clusterName)
	err = errors.Join(errs...)
	if errors.Is(err, sql.ErrNoRows) {
		h.Logger.Errorf("no rows found ids %v, err: %s", req.DeploymentIDs, err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "Not found"})
	}

	if err != nil {
		h.Logger.Errorf("error when deploying deployment err: %s", err)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: err.Error()})
	}

	for _, d := range listDeployment {
		for _, dev := range devices {
			hist := history.Histories{
				DeviceID:     dev.ID,
				DeploymentID: d.ID,
				CompanyID:    d.CompanyID,
				RepositoryID: d.RepositoryID,
			}

			hist, err := h.HistoryUsecase.Create(ctx, hist)
			if err != nil {
				h.Logger.Errorf("history: error when creating deployment histories err: %s", err)
				return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
			}

			go func(d DeploymentWithRepository, clusterName string) {
				pollingDuration := time.Duration(h.Config.PollingTimeout) * time.Second
				timeoutDuration := time.Duration(h.Config.PollingTimeout) * time.Second

				goCtx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
				defer cancel()
				for {
					select {
					case <-goCtx.Done():
						h.Logger.Errorf("timeout reached: deleting deployment %s", d.Name)
						h.Usecase.DeleteDeploy(context.Background(), []DeploymentWithRepository{d}, clusterName)
						h.HistoryUsecase.UpdateStatusById(context.Background(), hist.ID, "FAILED")
						return
					case <-time.After(pollingDuration):
						h.Logger.Infof("checking status deployment %s", d.Name)
						if h.Usecase.CheckDeploymentStatus(context.Background(), d.Name, clusterName) {
							h.HistoryUsecase.UpdateStatusById(context.Background(), hist.ID, "SUCCESS")
							return
						}
					}
				}
			}(d, clusterName)
		}
	}

	return c.JSON(http.StatusCreated, handler.SuccessResponse{Data: listDeployment})
}

func (h *Handler) V1DeleteDeploy(c echo.Context) error {
	req := DeleteDeploymentRequest{}
	ctx := c.Request().Context()
	companyID, ok := c.Get("companyID").(uuid.UUID)

	if !ok {
		h.Logger.Errorf("error when converting company id to string")
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "internal server error"})
	}

	clusterName, ok := c.Get("clusterName").(string)

	if !ok {
		h.Logger.Errorf("company: cluster name not found %s", clusterName)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "invalid cluster name"})
	}

	if err := c.Bind(&req); err != nil {
		h.Logger.Errorf("error when receiving request err: %s", err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	v := validatorx.New()
	if err := v.StructCtx(ctx, &req); err != nil {
		h.Logger.Errorf("error when validating request: %v, err: %s", req, err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	if _, err := h.CompanyUsecase.GetByID(ctx, companyID); err != nil {
		err := fmt.Errorf("invalid companyID %s, err: %s", companyID, err)
		h.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	// Get all available deployments from IDS
	deployments, err := h.Usecase.GetDeploymentWithRepositoryByIDs(ctx, companyID, req.DeploymentIDs)
	if errors.Is(err, sql.ErrNoRows) {
		err := fmt.Errorf("deployments with ids %v not found, err: %s", req.DeploymentIDs, err)
		h.Logger.Error(err)
		return c.JSON(http.StatusNotFound, handler.ErrorResponse{Message: "not found"})
	}

	if err != nil {
		err := fmt.Errorf("error when getting deployments with ids %v, err: %s", req.DeploymentIDs, err)
		h.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: err.Error()})
	}

	errs := h.Usecase.DeleteDeploy(ctx, deployments, clusterName)
	if len(errs) != 0 {
		errStr := ""
		for _, e := range errs {
			errStr += e.Error() + ","
		}
		h.Logger.Errorf("error when removing deployment err: %s", errStr)
		return c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: errStr})
	}

	return c.JSON(http.StatusNoContent, handler.SuccessResponse{Data: "success deleting all deployment"})
}
