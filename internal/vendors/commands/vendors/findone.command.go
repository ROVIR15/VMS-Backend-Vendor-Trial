package commands

import (
	"context"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/domain/entities"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/repositories"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
)

type VendorFindOneHandler struct {
	log        logger.Logger
	repository repositories.VendorRepository
}

func NewVendorFindOneHandler(log logger.Logger, repository repositories.VendorRepository) *VendorFindOneHandler {
	return &VendorFindOneHandler{
		log:        log,
		repository: repository,
	}
}

func (h *VendorFindOneHandler) Execute(ctx context.Context, VendorID int) (*entities.Vendors, error) {

	vendor, err := h.repository.FindOne(ctx, VendorID, h.log)

	if err != nil {
		return nil, err
	}

	return vendor, nil
}
