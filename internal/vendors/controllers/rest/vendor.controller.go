package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/controllers/rest/dtos"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/interactor"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/response"
	"github.com/labstack/echo/v4"
)

type VendorsController struct {
	group   *echo.Group
	logger  logger.Logger
	usecase *interactor.Usecase
}

func NewVendorController(group *echo.Group, logger logger.Logger, usecase *interactor.Usecase) *VendorsController {
	return &VendorsController{
		group:   group,
		logger:  logger,
		usecase: usecase,
	}
}

func (h *VendorsController) RegisterPath() {
	h.group.GET("/:id", h.FindOneVendor)
	h.group.GET("", h.FindAll)
	h.group.POST("", h.CreateNewVendor)
	h.group.PUT("/:id", h.UpdateVendor)
	h.group.DELETE("/:id", h.DeleteVendor)
}

func (h *VendorsController) FindOneVendor(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	paramID := e.Param("id")

	vendorId, err := strconv.Atoi(paramID)

	if err != nil {
		return response.Success(e, http.StatusBadRequest, response.SuccessMessage("Invalid vendor ID"))
	}

	vendor, err := h.usecase.FindOneVendorCommand.Execute(ctx, vendorId)
	if err != nil {
		return response.Success(e, http.StatusNotFound, response.SuccessMessage(err.Error()))
	}

	return response.Success(e, 200, response.SuccessData(vendor))
}

func (h *VendorsController) FindAll(e echo.Context) error {
	ctx := e.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	vendor, err := h.usecase.FindAllVendorCommand.Execute(ctx)
	if err != nil {
		return response.Success(e, 404, response.SuccessMessage(err.Error()))
	}
	return response.Success(e, 200, response.SuccessData(vendor))

}

// Refactored code with error handling for decoding JSON payload into vendorDto
func (h *VendorsController) CreateNewVendor(e echo.Context) error {
	ctx := e.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var vendorDto dtos.CreateVendorRequest

	err := json.NewDecoder(e.Request().Body).Decode(&vendorDto)
	if err != nil {
		return err
	}

	// Perform additional error handling for individual fields if required
	// Example: if vendorDto.Name is nil or empty, handle the error accordingly

	// Proceed with further processing using the decoded vendorDto
	_, err = h.usecase.CreateNewVendorCommand.Execute(ctx, &vendorDto)

	if err != nil {
		return response.Success(e, 404, response.SuccessMessage(err.Error()))
	}

	return response.Success(e, 200, response.SuccessMessage("Vendor Created"))

}

func (h *VendorsController) UpdateVendor(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	param := e.Param("id")

	vendorID, err := strconv.Atoi(param)
	if err != nil {
		return err
	}

	var vendorDto dtos.UpdateVendorRequest

	err = json.NewDecoder(e.Request().Body).Decode(&vendorDto)
	if err != nil {
		return err
	}

	// Perform additional error handling for individual fields if required
	// Example: if vendorDto.Name is nil or empty, handle the error accordingly

	// Proceed with further processing using the decoded vendorDto
	_, err = h.usecase.UpdateVendorCommand.Execute(ctx, vendorID, &vendorDto)

	if err != nil {
		return response.Success(e, 404, response.SuccessMessage(err.Error()))
	}

	return response.Success(e, 200, response.SuccessMessage("Vendor Updated"))
}

func (h *VendorsController) DeleteVendor(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	param := e.Param("id")

	vendorID, err := strconv.Atoi(param)
	if err != nil {
		return err
	}

	err = h.usecase.DeleteVendorCommand.Execute(ctx, vendorID)

	if err != nil {
		return response.Success(e, 404, response.SuccessMessage(err.Error()))
	}

	return response.Success(e, 200, response.SuccessMessage("Vendor Deleted"))
}
