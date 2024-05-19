package interactor

import (
	commands "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/commands/phone"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/repositories"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
)

type UseCaseVendorPhone struct {
	UpdateVendorPhoneCommand *commands.VendorPhoneUpdateHandler
}

func NewUseCaseVendorPhone(
	logger logger.Logger,
	repository repositories.VendorPhoneRepository,
) *UseCaseVendorPhone {
	return &UseCaseVendorPhone{
		UpdateVendorPhoneCommand: commands.NewVendorPhoneUpdateHandler(logger, repository),
	}
}
