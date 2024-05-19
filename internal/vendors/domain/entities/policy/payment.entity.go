package entities

type VendorPaymentPolicy struct {
	ID                        int
	VendorID                  int
	IsCollectAfterFulfillment bool
}

type NotFoundError struct {
	Entity string
	ID     int
}

func NewVendorPaymentPolicy(id int, vendorID int, isCollectAfterFulfillment bool) *VendorPaymentPolicy {
	return &VendorPaymentPolicy{
		ID:                        id,
		VendorID:                  vendorID,
		IsCollectAfterFulfillment: isCollectAfterFulfillment,
	}
}

func (v *VendorPaymentPolicy) GetID() int {
	return v.ID
}

func (v *VendorPaymentPolicy) FindOneVendorPaymentPolicyByVendorID(vendorID int) (*VendorPaymentPolicy, *NotFoundError) {
	if v.VendorID == vendorID {
		return v, nil
	}

	return nil, &NotFoundError{Entity: "VendorPaymentPolicy", ID: vendorID}
}

func (v *VendorPaymentPolicy) GetVendorID() int {
	return v.VendorID
}
