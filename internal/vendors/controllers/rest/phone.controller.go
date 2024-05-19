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

type VendorPhoneControoller struct {
	group   *echo.Group
	logger  logger.Logger
	usecase *interactor.UseCaseVendorPhone
}

func NewVendorPhoneController(group *echo.Group, logger logger.Logger, usecase *interactor.UseCaseVendorPhone) *VendorPhoneControoller {
	return &VendorPhoneControoller{
		group:   group,
		logger:  logger,
		usecase: usecase,
	}
}

func (h *VendorPhoneControoller) RegisterPath() {
	h.group.PUT("/:id/phone", h.UpdatePhoneNumber)
}

func (h *VendorPhoneControoller) UpdatePhoneNumber(e echo.Context) error {
	ctx := e.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	paramID := e.Param("id")

	vendorId, err := strconv.Atoi(paramID)
	if err != nil {
		return response.Success(e, http.StatusBadRequest, response.SuccessMessage("Invalid vendor ID"))
	}

	var vendorPhoneDto dtos.UpdateVendorPhoneRequest
	err = json.NewDecoder(e.Request().Body).Decode(&vendorPhoneDto)
	if err != nil {
		return err
	}

	// Perform additional error handling for individual fields if required
	// Example: if vendorPhoneDto.PhoneCountryCode is nil or empty, handle the error accordingly

	// Proceed with further processing using the decode vendorDto

	_, err = h.usecase.UpdateVendorPhoneCommand.Execute(ctx, vendorId, &vendorPhoneDto)

	if err != nil {
		return response.Success(e, http.StatusNotFound, response.SuccessMessage(err.Error()))
	}

	return response.Success(e, http.StatusCreated, response.SuccessMessage("Vendor phone updated successfully"))
}
