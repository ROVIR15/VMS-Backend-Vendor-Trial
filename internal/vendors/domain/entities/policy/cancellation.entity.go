package entities

type VendorRefundAndCancellationPolicy struct {
	ID                    int
	VendorID              int
	IsGuestPayTransaction bool
}

// NewVendorRefundAndCancellationPolicy creates a new instance of VendorRefundAndCancellationPolicy
// with the given parameters.
//
// Parameters:
// - id: The ID of the vendor refund and cancellation policy.
// - vendorID: The ID of the vendor.
// - isGuestPayTransaction: Flag indicating whether the policy is for guest pay transaction.
//
// Returns:
// - *VendorRefundAndCancellationPolicy: A pointer to the newly created VendorRefundAndCancellationPolicy instance.
func NewVendorRefundAndCancellationPolicy(
	id int,
	vendorID int,
	isGuestPayTransaction bool,
) *VendorRefundAndCancellationPolicy {
	return &VendorRefundAndCancellationPolicy{
		ID:                    id,
		VendorID:              vendorID,
		IsGuestPayTransaction: isGuestPayTransaction,
	}
}

// FindOneVendorRefundAndCancellationPolicy returns the VendorRefundAndCancellationPolicy that matches the given vendor_id,
// or an error if no such policy exists.
func (v *VendorRefundAndCancellationPolicy) FindOneVendorRefundAndCancellationPolicy(vendor_id int) (*VendorRefundAndCancellationPolicy, *NotFoundError) {
	if v.VendorID == vendor_id {
		return v, nil
	}
	return nil, &NotFoundError{Entity: "VendorRefundAndCancellationPolicy", ID: vendor_id}
}

// GetID returns the ID of the VendorRefundAndCancellationPolicy.
//
// Returns:
// - int: The ID of the VendorRefundAndCancellationPolicy.
func (v *VendorRefundAndCancellationPolicy) GetID() int {
	return v.ID
}
