package commands

import (
	"context"
	"errors"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/controllers/rest/dtos"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/repositories"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
)

type VendorPhoneUpdateHandler struct {
	log        logger.Logger
	repository repositories.VendorPhoneRepository
}

type VendorPhoneUpdateCommand struct {
	PhoneCountryCode int
	VendorPhoneType  string
}

func NewVendorPhoneUpdateCommand(
	PhoneCountryCode int,
	VendorPhoneType string,
) *VendorPhoneUpdateCommand {
	return &VendorPhoneUpdateCommand{
		PhoneCountryCode: PhoneCountryCode,
		VendorPhoneType:  VendorPhoneType,
	}
}

func NewVendorPhoneUpdateHandler(log logger.Logger, repository repositories.VendorPhoneRepository) *VendorPhoneUpdateHandler {
	return &VendorPhoneUpdateHandler{
		log:        log,
		repository: repository,
	}
}

func (h *VendorPhoneUpdateHandler) Execute(ctx context.Context, vendorID int, updateVendorPhoneRequest *dtos.UpdateVendorPhoneRequest) (int, error) {
	const (
		vendorPhoneNotFoundError = "vendor phone not found by vendor ID"
		vendorPhoneUpdateFailed  = "failed to update vendor phone"
	)

	//GET vendor phone by vendor ID
	vendorPhone, err := h.repository.FindPhoneNumberByVendorID(ctx, vendorID, h.log)
	if err != nil {
		return 0, errors.New(vendorPhoneNotFoundError)
	}

	vendorPhone.SetPhoneCountryCode(&updateVendorPhoneRequest.PhoneCountryCode)
	vendorPhone.SetVendorPhoneType(&updateVendorPhoneRequest.VendorPhoneType)

	//UPDATE vendor phone
	_, err = h.repository.Update(ctx, vendorID, vendorPhone, h.log)
	if err != nil {
		return 0, errors.New(vendorPhoneUpdateFailed)
	}

	return vendorID, nil
}
