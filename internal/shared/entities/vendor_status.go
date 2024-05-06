package entities

type VendorStatus struct {
	id          int    `json:"id"`
	vendor_id   int    `json:"vendor_id"`
	user_id     int    `json:"user_id"`
	active      bool   `json:"active"`
	description string `json:"description"`
}
