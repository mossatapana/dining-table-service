package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mossatapana/dining-table-service/models"
)

type TableService interface {
	Create(ctx context.Context, noOfTables int) error
	Reserve(ctx context.Context, noOfCus int) error
	Cancel(ctx context.Context, bookingID string) error
}

type TableHandler struct {
	tableService TableService
}

func NewTableHandlers(tableService TableService) *TableHandler {
	return &TableHandler{
		tableService: tableService,
	}
}

func (h TableHandler) Initial(c echo.Context) error {
	var req models.InitialRequest
	if err := c.Bind(&req); err != nil {
		errResp := models.CommonResponse{
			Success: false,
			Message: "Request body is invalid",
		}
		return c.JSON(http.StatusBadRequest, errResp)
	}

	var resp models.CommonResponse

	err := h.tableService.Create()

	resp.Success = true
	return c.JSON(http.StatusCreated, resp)
}

func (h TableHandler) Reserve(c echo.Context) error {
	var req models.ReserveRequest
	if err := c.Bind(&req); err != nil {
		errResp := models.CommonResponse{
			Success: false,
			Message: "Request body is invalid",
		}
		return c.JSON(http.StatusBadRequest, errResp)
	}

	var resp models.ReserveResponse

	resp.Success = true
	return c.JSON(http.StatusCreated, resp)
}

func (h TableHandler) Cancel(c echo.Context) error {
	var req models.CancelRequest
	if err := c.Bind(&req); err != nil {
		errResp := models.CommonResponse{
			Success: false,
			Message: "Request body is invalid",
		}
		return c.JSON(http.StatusBadRequest, errResp)
	}

	var resp models.CancelResponse

	resp.Success = true
	return c.JSON(http.StatusOK, resp)
}
