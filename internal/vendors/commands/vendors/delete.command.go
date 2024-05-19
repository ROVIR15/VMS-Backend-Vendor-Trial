package commands

import (
	"context"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/repositories"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
)

type VendorDeleteHandler struct {
	log        logger.Logger
	repository repositories.VendorRepository
}

func NewVendorDeleteHandler(log logger.Logger, repository repositories.VendorRepository) *VendorDeleteHandler {
	return &VendorDeleteHandler{
		log:        log,
		repository: repository,
	}
}

func (h *VendorDeleteHandler) Execute(ctx context.Context, VendorID int) error {

	err := h.repository.Delete(ctx, VendorID, h.log)

	if err != nil {
		return err
	}

	return nil
}
