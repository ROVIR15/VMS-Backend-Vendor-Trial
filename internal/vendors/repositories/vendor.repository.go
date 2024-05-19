package repositories

import (
	"context"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/domain/entities"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
)

type VendorRepository interface {
	FindOne(ctx context.Context, VendorID int, log logger.Logger) (*entities.Vendors, error)
	FindAll(ctx context.Context, log logger.Logger) ([]*entities.Vendors, error)
	Create(ctx context.Context, vendor *entities.Vendors, log logger.Logger) (int, error)
	Update(ctx context.Context, VendorID int, vendor *entities.Vendors, log logger.Logger) (int, error)
	Delete(ctx context.Context, VendorID int, log logger.Logger) error
}
