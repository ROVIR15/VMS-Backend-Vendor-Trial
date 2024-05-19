package repositories

import (
	"context"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/domain/entities"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
)

type VendorPhoneRepository interface {
	Create(ctx context.Context, VendorPhone *entities.VendorPhone, log logger.Logger) (int, error)
	Update(ctx context.Context, VendorPhoneID int, vendorPhone *entities.VendorPhone, log logger.Logger) (int, error)
	Delete(ctx context.Context, VendorPhoneID int, log logger.Logger) error
	FindPhoneNumberByVendorID(ctx context.Context, VendorID int, log logger.Logger) (*entities.VendorPhone, error)
}
