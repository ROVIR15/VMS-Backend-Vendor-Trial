package commands

import (
	"context"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/controllers/rest/dtos"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/domain/entities"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/repositories"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
)

type VendorCreateHandler struct {
	log        logger.Logger
	repository repositories.VendorRepository
}

// type CreateNewVendorCommand struct {
// 	Name             string `json:name`
// 	Email            string `json:email`
// 	Commission       string `json:commission`
// 	VendorVisibility string `json:vendorVisibility`
// 	Description      string `json:description`
// 	Logo             string `json:logo`
// }

func NewVendorCreateHandler(log logger.Logger, repository repositories.VendorRepository) *VendorCreateHandler {
	return &VendorCreateHandler{
		log:        log,
		repository: repository,
	}
}

func (h *VendorCreateHandler) Execute(ctx context.Context, dtos *dtos.CreateVendorRequest) (int, error) {

	vendor := entities.NewVendor(
		dtos.Name,
		dtos.Email,
		dtos.Commission,
		dtos.VendorVisibility,
		dtos.Description,
		dtos.Logo,
	)

	_, err := h.repository.Create(ctx, vendor, h.log)
	if err != nil {
		return 0, err
	}

	return vendor.ID, nil
}
