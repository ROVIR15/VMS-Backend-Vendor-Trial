package entities

type VendorPhone struct {
	ID               int
	VendorID         *int
	PhoneCountryCode *string
	VendorPhoneType  *string
}

type NotFoundError struct {
	Entity string
	ID     int
}

func NewVendorPhone(id int, vendorID *int, phoneCountryCode *string, vendorPhoneType *string) *VendorPhone {
	return &VendorPhone{
		ID:               id,
		VendorID:         vendorID,
		PhoneCountryCode: phoneCountryCode,
		VendorPhoneType:  vendorPhoneType,
	}
}

func (v *VendorPhone) GetID() int {
	return v.ID
}

func (v *VendorPhone) GetVendorID() *int {
	return v.VendorID
}

func (v *VendorPhone) SetPhoneCountryCode(phoneCountryCode *string) {
	v.PhoneCountryCode = phoneCountryCode
}

func (v *VendorPhone) SetVendorPhoneType(vendorPhoneType *string) {
	v.VendorPhoneType = vendorPhoneType
}
