package commands

import (
	"context"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/controllers/rest/dtos"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/domain/entities"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/repositories"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
)

type VendorUpdateHandler struct {
	log        logger.Logger
	repository repositories.VendorRepository
}

func NewVendorUpdateHandler(log logger.Logger, repository repositories.VendorRepository) *VendorUpdateHandler {
	return &VendorUpdateHandler{
		log:        log,
		repository: repository,
	}
}

func (h *VendorUpdateHandler) Execute(ctx context.Context, VendorID int, dtos *dtos.UpdateVendorRequest) (int, error) {

	vendor := entities.NewVendor(
		dtos.Name,
		dtos.Email,
		dtos.Commission,
		dtos.VendorVisibility,
		dtos.Description,
		dtos.Logo,
	)

	_, err := h.repository.Update(ctx, VendorID, vendor, h.log)
	if err != nil {
		return 0, err
	}

	return vendor.ID, nil
}
