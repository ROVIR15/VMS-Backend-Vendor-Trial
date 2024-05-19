package entities

type VendorBookingPolicy struct {
	ID        int
	VendorID  int
	DaysPrior int
}

// NewVendorBookingPolicy creates a new VendorBookingPolicy instance.
//
// id: The ID of the policy.
// vendorID: The ID of the vendor.
// daysPrior: The number of days prior to the booking date.
// Returns a pointer to the newly created VendorBookingPolicy instance.
func NewVendorBookingPolicy(id int, vendorID int, daysPrior int) *VendorBookingPolicy {
	return &VendorBookingPolicy{
		ID:        id,
		VendorID:  vendorID,
		DaysPrior: daysPrior,
	}
}

func (v *VendorBookingPolicy) GetID() int {
	return v.ID
}
