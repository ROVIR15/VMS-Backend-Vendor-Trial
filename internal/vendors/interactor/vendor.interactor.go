package interactor

import (
	companyCommand "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/commands/vendors"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/repositories"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
)

type Usecase struct {
	FindOneVendorCommand   *companyCommand.VendorFindOneHandler
	FindAllVendorCommand   *companyCommand.VendorFindAllHandler
	CreateNewVendorCommand *companyCommand.VendorCreateHandler
	UpdateVendorCommand    *companyCommand.VendorUpdateHandler
	DeleteVendorCommand    *companyCommand.VendorDeleteHandler
}

func NewUseCase(
	logger logger.Logger,
	repository repositories.VendorRepository,
) *Usecase {
	return &Usecase{
		FindOneVendorCommand:   companyCommand.NewVendorFindOneHandler(logger, repository),
		FindAllVendorCommand:   companyCommand.NewVendorFindAllHandler(logger, repository),
		CreateNewVendorCommand: companyCommand.NewVendorCreateHandler(logger, repository),
		UpdateVendorCommand:    companyCommand.NewVendorUpdateHandler(logger, repository),
		DeleteVendorCommand:    companyCommand.NewVendorDeleteHandler(logger, repository),
	}
}
