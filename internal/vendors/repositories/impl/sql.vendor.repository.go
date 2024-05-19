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

type SqlVendorsRepository struct {
	db  *sql.DB
	trx *sql.Tx
}

var _ repositories.VendorRepository = (*SqlVendorsRepository)(nil)

func NewSQLVendorRepositoryConn(db *sql.DB) *SqlVendorsRepository {
	return &SqlVendorsRepository{db: db}
}

// SetTransaction sets the transaction on the repository
// This can be used to group multiple queries into a single transaction
//
// The transaction is stored in the repository and will be passed to any subsequent
// calls to the repository that use the WithTransaction context method
//
// If the transaction is not set, the repository will default to using
// a new transaction for each query
//
// If the transaction is set and the context is cancelled, the transaction
// will be rolled back
func (r *SqlVendorsRepository) SetTransaction(trx *sql.Tx) {
	r.trx = trx
}

func (r *SqlVendorsRepository) FindOne(ctx context.Context, id int, log logger.Logger) (*entities.Vendors, error) {
	// Simulate fetching a vendor from a database or some other data source
	// For demonstration purposes, we'll return a hard-coded vendor

	// Current timestamp
	// now := time.Now().Unix()
	var exec boil.ContextExecutor = r.db
	if r.trx != nil {
		exec = r.trx
	}

	vendorModel, err := models.Vendors(qm.Where("id = ?", id)).One(ctx, exec)
	if err != nil {
		return nil, err
	}

	__vendor := &entities.Vendors{
		ID:               int(vendorModel.ID),
		Name:             vendorModel.Name.String,
		Email:            vendorModel.Email.String,
		Commission:       vendorModel.Commission.Float64,
		Description:      vendorModel.Description.String,
		Logo:             vendorModel.Logo.String,
		VendorVisibility: vendorModel.VendorVisibility.Val.String(),
		TrialEndDateAt:   vendorModel.TrialEndDateAt.Time,
		CreatedAt:        vendorModel.CreatedAt.Time,
		UpdatedAt:        vendorModel.UpdatedAt.Time,
		DeletedAt:        vendorModel.DeletedAt.Time,
	}

	return __vendor, nil
}

func (r *SqlVendorsRepository) FindAll(ctx context.Context, log logger.Logger) ([]*entities.Vendors, error) {

	var exec boil.ContextExecutor = r.db
	if r.trx != nil {
		exec = r.trx
	}

	vendorList, err := models.Vendors().All(ctx, exec)
	if err != nil {
		return nil, err
	}
	var vendors []*entities.Vendors
	for _, vendor := range vendorList {
		vendors = append(vendors, &entities.Vendors{
			ID:               int(vendor.ID),
			Name:             vendor.Name.String,
			Email:            vendor.Email.String,
			Commission:       vendor.Commission.Float64,
			Description:      vendor.Description.String,
			Logo:             vendor.Logo.String,
			VendorVisibility: vendor.VendorVisibility.Val.String(),
			TrialEndDateAt:   vendor.TrialEndDateAt.Time,
			CreatedAt:        vendor.CreatedAt.Time,
			UpdatedAt:        vendor.UpdatedAt.Time,
			DeletedAt:        vendor.DeletedAt.Time,
		})
	}
	return vendors, nil

}

func (r *SqlVendorsRepository) Create(ctx context.Context, vendor *entities.Vendors, log logger.Logger) (int, error) {

	var exec boil.ContextExecutor = r.db
	if r.trx != nil {
		exec = r.trx
	}

	vendorModel := &models.Vendor{
		Name:             null.StringFrom(vendor.Name),
		Email:            null.StringFrom(vendor.Email),
		Commission:       null.Float64From(vendor.Commission),
		VendorVisibility: models.NullVendorVisibilityFrom(models.VendorVisibility(vendor.VendorVisibility)),
		Description:      null.StringFrom(vendor.Description),
		Logo:             null.StringFrom(vendor.Logo),
	}

	err := vendorModel.Insert(ctx, exec, boil.Infer())

	if err != nil {
		return 0, err
	}

	return vendorModel.ID, nil
}

func (r *SqlVendorsRepository) Update(ctx context.Context, vendorId int, vendor *entities.Vendors, log logger.Logger) (int, error) {

	var exec boil.ContextExecutor = r.db
	if r.trx != nil {
		exec = r.trx
	}

	updateColumns := models.M{
		"name":              vendor.Name,
		"email":             vendor.Email,
		"commission":        vendor.Commission,
		"vendor_visibility": vendor.VendorVisibility,
		"description":       vendor.Description,
		"logo":              vendor.Logo,
	}

	var err error

	_, err = models.Vendors(qm.Where("id = ?", vendorId)).UpdateAll(ctx, exec, updateColumns)

	if err != nil {
		return 0, err
	}
	return int(vendorId), nil
}

func (r *SqlVendorsRepository) Delete(ctx context.Context, id int, log logger.Logger) error {

	var exec boil.ContextExecutor = r.db
	if r.trx != nil {
		exec = r.trx
	}

	_, err := models.Vendors(qm.Where("id =?", id)).DeleteAll(ctx, exec)
	if err != nil {
		return err
	}

	return nil
}
