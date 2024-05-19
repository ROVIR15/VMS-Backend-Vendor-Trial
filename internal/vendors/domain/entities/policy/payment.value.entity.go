package entities

type VendorPaymentPolicyValueCode string

const (
	AfterBookingIsFulfilled VendorPaymentPolicyValueCode = "after_booking_is_fulfilled"
	BeforeServiceDate       VendorPaymentPolicyValueCode = "before_service_date"
	AfterBookingConfirmed   VendorPaymentPolicyValueCode = "after_booking_confirmed"
)

type VendorPaymentPolicyValue struct {
	ID                           int
	VendorPaymentPolicyID        int
	VendorPaymentPolicyValueCode VendorPaymentPolicyValueCode
	ChargePercentage             float32
	Days                         int
	Hours                        int
	Minutes                      int
}

// NewVendorPaymentPolicyValue creates a new VendorPaymentPolicyValue instance.
// It takes the required fields as parameters and returns a pointer to the newly created instance.
func NewVendorPaymentPolicyValue(id int, vendorPaymentPolicyID int, vendorPaymentPolicyValueCode VendorPaymentPolicyValueCode, chargePercentage float32, days int, hours int, minutes int) *VendorPaymentPolicyValue {
	return &VendorPaymentPolicyValue{
		ID:                           id,
		VendorPaymentPolicyID:        vendorPaymentPolicyID,
		VendorPaymentPolicyValueCode: vendorPaymentPolicyValueCode,
		ChargePercentage:             chargePercentage,
		Days:                         days,
		Hours:                        hours,
		Minutes:                      minutes,
	}
}

// IsValid checks if the VendorPaymentPolicyValueCode is valid.
// It returns true if the code is valid, otherwise it returns false.
func (v *VendorPaymentPolicyValue) IsValid() bool {
	validCodes := []VendorPaymentPolicyValueCode{
		AfterBookingIsFulfilled,
		BeforeServiceDate,
		AfterBookingConfirmed,
	}

	for _, code := range validCodes {
		if v.VendorPaymentPolicyValueCode == code {
			return true
		}
	}

	return false
}
