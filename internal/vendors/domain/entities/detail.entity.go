package entities

type VendorDetail struct {
	ID                     int
	VendorID               *int
	Commission             float32
	IsCommissionIncludeTax bool
}

func NewVendorDetail(id int, vendorID *int, commission float32, isCommissionIncludeTax bool) *VendorDetail {
	return &VendorDetail{
		ID:                     id,
		VendorID:               vendorID,
		Commission:             commission,
		IsCommissionIncludeTax: isCommissionIncludeTax,
	}
}

func (v *VendorDetail) GetID() int {
	return v.ID
}

func (v *VendorDetail) FindOneVendorDetailByVendorID(vendorID int) (*VendorDetail, *NotFoundError) {
	if v.VendorID != nil && *v.VendorID == vendorID {
		return v, nil
	}

	return nil, &NotFoundError{Entity: "VendorDetail", ID: vendorID}
}
