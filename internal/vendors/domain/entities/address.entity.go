package entities

type VendorAddress struct {
	ID        int     `boil:"id" json:"id" toml:"id" yaml:"id"`
	VendorID  *int    `boil:"vendor_id" json:"vendor_id" toml:"vendor_id" yaml:"vendor_id"`
	Address   *string `boil:"address" json:"address" toml:"address" yaml:"address"`
	Location  *string `boil:"location" json:"location" toml:"location" yaml:"location"`
	IsPrimary *bool   `boil:"is_primary" json:"is_primary" toml:"is_primary" yaml:"is_primary"`
}

func NewVendorAddress(id int, vendorID *int, address *string, location *string, isPrimary *bool) *VendorAddress {
	return &VendorAddress{
		ID:        id,
		VendorID:  vendorID,
		Address:   address,
		Location:  location,
		IsPrimary: isPrimary,
	}
}

// GetID returns the ID of the VendorAddress.
//
// This function returns the ID of the VendorAddress.
// The returned value is of type int.
func (v *VendorAddress) GetID() int {
	// Returns the ID of the VendorAddress.
	return v.ID
}

// GetVendorID returns the vendor ID of the VendorAddress.
//
// This function returns the vendor ID of the VendorAddress.
// The returned value is of type *int.
// It can be nil if the vendor ID is not set.
func (v *VendorAddress) GetVendorID() *int {
	// Returns the vendor ID of the VendorAddress.
	// It can be nil if the vendor ID is not set.
	return v.VendorID
}

// GetAddress returns the address of the VendorAddress.
//
// This function returns the address of the VendorAddress.
// The returned value is of type *string.
// It can be nil if the address is not set.
func (v *VendorAddress) GetAddress() *string {
	// Returns the address of the VendorAddress.
	// It can be nil if the address is not set.
	return v.Address
}

// GetLocation returns the location of the VendorAddress.
//
// This function returns the location of the VendorAddress.
// The returned value is of type *string.
// It can be nil if the location is not set.
func (v *VendorAddress) GetLocation() *string {
	// Returns the location of the VendorAddress.
	// It can be nil if the location is not set.
	return v.Location
}

// GetIsPrimary returns the is primary of the VendorAddress.
//
// This function returns the is primary of the VendorAddress.
// The returned value is of type *bool.
// It can be nil if the is primary is not set.
func (v *VendorAddress) GetIsPrimary() *bool {
	// Returns the is primary of the VendorAddress.
	// It can be nil if the is primary is not set.
	return v.IsPrimary
}
