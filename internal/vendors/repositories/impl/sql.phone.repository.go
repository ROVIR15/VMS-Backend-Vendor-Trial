package impl

import (
	"context"
	"database/sql"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/shared/models"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/domain/entities"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/repositories"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SqlVendorPhoneRepository struct {
	db  *sql.DB
	trx *sql.Tx
}

var _ repositories.VendorPhoneRepository = (*SqlVendorPhoneRepository)(nil)

func NewSqlvendorPhoneRepository(db *sql.DB) *SqlVendorPhoneRepository {
	return &SqlVendorPhoneRepository{db: db}
}

func (r *SqlVendorPhoneRepository) SetTransaction(trx *sql.Tx) {
	r.trx = trx
}

func (r *SqlVendorPhoneRepository) Create(ctx context.Context, vendorPhone *entities.VendorPhone, log logger.Logger) (int, error) {
	var exec boil.ContextExecutor = r.db
	if r.trx != nil {
		exec = r.trx
	}

	vendorPhoneModel := &models.VendorPhone{
		VendorID:         null.IntFromPtr(vendorPhone.VendorID),
		PhoneCountryCode: null.StringFromPtr(vendorPhone.PhoneCountryCode),
		VendorPhoneType:  models.NullVendorPhoneTypeFrom(models.VendorPhoneType(*vendorPhone.VendorPhoneType)),
	}

	err := vendorPhoneModel.Insert(ctx, exec, boil.Infer())

	if err != nil {
		return 0, err
	}

	return vendorPhoneModel.ID, nil
}

func (r *SqlVendorPhoneRepository) Delete(ctx context.Context, VendorPhoneID int, log logger.Logger) error {
	var exec boil.ContextExecutor = r.db
	if r.trx != nil {
		exec = r.trx
	}

	_, err := models.VendorPhones(qm.Where("id= ?", VendorPhoneID)).DeleteAll(ctx, exec)
	if err != nil {
		return err
	}

	return nil
}

func (r *SqlVendorPhoneRepository) FindPhoneNumberByVendorID(ctx context.Context, VendorID int, log logger.Logger) (*entities.VendorPhone, error) {
	var exec boil.ContextExecutor = r.db
	if r.trx != nil {
		exec = r.trx
	}

	vendorPhoneModel, err := models.VendorPhones(qm.Where("vendorId=?", VendorID)).One(ctx, exec)

	if err != nil {
		return nil, err
	}

	phoneType := vendorPhoneModel.VendorPhoneType.Val.String()

	vendorPhone := entities.NewVendorPhone(
		vendorPhoneModel.ID,
		&VendorID,
		&vendorPhoneModel.PhoneCountryCode.String,
		&phoneType,
	)

	return vendorPhone, nil

}

func (r *SqlVendorPhoneRepository) Update(ctx context.Context, VendorPhoneID int, vendorPhone *entities.VendorPhone, log logger.Logger) (int, error) {
	var exec boil.ContextExecutor = r.db
	if r.trx != nil {
		exec = r.trx
	}

	updateColumnValue := models.M{
		"phone_country_code": vendorPhone.PhoneCountryCode,
		"vendor_phone_type":  vendorPhone.VendorPhoneType,
	}

	_, err := models.VendorPhones(qm.Where("id=?", VendorPhoneID)).UpdateAll(ctx, exec, updateColumnValue)

	if err != nil {
		return 0, err
	}

	return VendorPhoneID, nil
}
