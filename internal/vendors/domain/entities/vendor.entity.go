package entities

import "time"

type Vendors struct {
	ID               int
	Code             string
	Name             string
	Email            string
	Commission       float64
	VendorVisibility string
	Description      string
	Logo             string
	TrialEndDateAt   time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

func NewVendor(
	name string,
	email string,
	commission float64,
	vendorVisibility string,
	description string,
	logo string,
) *Vendors {

	vendor := &Vendors{
		ID:               0,
		Name:             name,
		Email:            email,
		Commission:       commission,
		VendorVisibility: vendorVisibility,
		Description:      description,
		Logo:             logo,
		TrialEndDateAt:   time.Time{},
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		DeletedAt:        time.Time{},
	}

	return vendor
}

func (v *Vendors) GetID() int {
	return v.ID
}

func (v *Vendors) OneVendor() *Vendors {
	return v
}

func (v *Vendors) ListVendors() []*Vendors {
	return []*Vendors{v}
}
