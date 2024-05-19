package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func New() Handler {
	return Handler{}
}

func (h *Handler) PingHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, SuccessResponse{Data: time.Now()})
}
