package commands

import (
	"context"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/domain/entities"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/repositories"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
)

type VendorFindAllHandler struct {
	log        logger.Logger
	repository repositories.VendorRepository
}

func NewVendorFindAllHandler(log logger.Logger, repository repositories.VendorRepository) *VendorFindAllHandler {
	return &VendorFindAllHandler{
		log:        log,
		repository: repository,
	}
}

func (h *VendorFindAllHandler) Execute(ctx context.Context) ([]*entities.Vendors, error) {

	vendor, err := h.repository.FindAll(ctx, h.log)

	if err != nil {
		return nil, err
	}

	return vendor, nil
}
