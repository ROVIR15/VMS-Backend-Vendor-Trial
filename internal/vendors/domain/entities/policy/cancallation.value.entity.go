package entities

import "encoding/json"

type VendorRefundAndCancellationPolicyValue struct {
	ID                       int     `json:"id"`
	VendorID                 int     `json:"vendor_id"`
	IsFullRefund             bool    `json:"is_full_refund"`
	Percentage               float64 `json:"percentage"`
	DaysPriorMoreThanOrEqual int     `json:"days_prior_more_than_equal"`
	DaysPriorStart           int     `json:"days_prior_start"`
	DaysPriorEnd             int     `json:"days_prior_end"`
	DaysPriorLessThan        int     `json:"days_prior_less_than"`
	GracePeriod              int     `json:"grace_period"`
	Order                    int     `json:"order"`
}

func NewVendorRefundAndCancellationPolicyValue(
	id int,
	vendorID int,
	isFullRefund bool,
	percentage float64,
	daysPriorMoreThanOrEqual int,
	daysPriorStart int,
	daysPriorEnd int,
	daysPriorLessThan int,
	gracePeriod int,
	order int,
) *VendorRefundAndCancellationPolicyValue {
	return &VendorRefundAndCancellationPolicyValue{
		ID:                       id,
		VendorID:                 vendorID,
		IsFullRefund:             isFullRefund,
		Percentage:               percentage,
		DaysPriorMoreThanOrEqual: daysPriorMoreThanOrEqual,
		DaysPriorStart:           daysPriorStart,
		DaysPriorEnd:             daysPriorEnd,
		DaysPriorLessThan:        daysPriorLessThan,
		GracePeriod:              gracePeriod,
		Order:                    order,
	}
}

func (v *VendorRefundAndCancellationPolicyValue) String() string {
	b, _ := json.Marshal(v)
	return string(b)
}
