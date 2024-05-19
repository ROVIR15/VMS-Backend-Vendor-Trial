package entities

import "errors"

type Host struct {
	ID                     int
	VendorId               int
	Code                   string
	Commission             float32
	IsCommissionIncludeTax bool
	IsAllowBooking         bool
}

// NewHost creates a new Host entity with the given parameters.
// It returns a pointer to the newly created Host entity.
//
// If any of the following conditions are met, it returns nil and an error:
// - If the id or vendorID is less than 1.
// - If the code is an empty string.
// - If the commission is less than 0.
// - If the isCommissionIncludeTax is false and isAllowBooking is true.
func NewHost(
	id int,
	vendorID int,
	code string,
	commission float32,
	isCommissionIncludeTax bool,
	isAllowBooking bool,
) (*Host, error) {
	if id < 1 || vendorID < 1 {
		return nil, errors.New("id and vendorID must be greater than 0")
	}
	if code == "" {
		return nil, errors.New("code cannot be empty")
	}
	if commission < 0 {
		return nil, errors.New("commission must be greater than or equal to 0")
	}
	if !isCommissionIncludeTax && isAllowBooking {
		return nil, errors.New("isCommissionIncludeTax must be true if isAllowBooking is true")
	}
	return &Host{
		ID:                     id,
		VendorId:               vendorID,
		Code:                   code,
		Commission:             commission,
		IsCommissionIncludeTax: isCommissionIncludeTax,
		IsAllowBooking:         isAllowBooking,
	}, nil
}

// FindAllHostByVendorId returns a slice of Host entities that belong to the given vendorID.
// If no hosts are found, it returns an empty slice.
// If an error occurs during the operation, it returns nil and the error.
func (v *Host) FindAllHostByVendorId(vendorID int) ([]*Host, error) {
	if vendorID < 1 {
		return nil, errors.New("vendorID must be greater than 0")
	}

	if v.VendorId == vendorID {
		return []*Host{v}, nil
	}

	return make([]*Host, 0), nil
}

// FindOneHost returns a Host entity with the given hostID.
// If no host is found, it returns nil and an error.
func (v *Host) FindOneHost(hostID int) (*Host, error) {
	if v.ID == hostID {
		return v, nil
	}
	return nil, errors.New("host not found")
}
