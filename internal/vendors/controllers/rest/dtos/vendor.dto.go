package dtos

import "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/domain/entities"

type CreateVendorRequest struct {
	Name             string  `json:"name"`
	Email            string  `json:"email"`
	Commission       float64 `json:"commission"`
	VendorVisibility string  `json:"vendorVisibility"`
	Description      string  `json:"description"`
	Logo             string  `json:"logo"`
}

type CreateVendorResponse struct {
	Vendor *entities.Vendors
}

type UpdateVendorRequest struct {
	Name             string  `json:"name"`
	Email            string  `json:"email"`
	Commission       float64 `json:"commission"`
	VendorVisibility string  `json:"vendorVisibility"`
	Description      string  `json:"description"`
	Logo             string  `json:"logo"`
}

type UpdateVendorResponse struct {
	Name             string  `json:"name"`
	Email            string  `json:"email"`
	Commission       float64 `json:"commission"`
	VendorVisibility string  `json:"vendorVisibility"`
	Description      string  `json:"description"`
	Logo             string  `json:"logo"`
}

type DeleteVendorRequest struct {
	VendorID int
}

type DeleteVendorResponse struct {
	Deleted bool
}
