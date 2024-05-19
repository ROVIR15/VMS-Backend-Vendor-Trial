package entities

import (
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type VendorLog struct {
	ID         int        `boil:"id"`
	VendorID   int        `boil:"vendor_id"`
	UserID     null.Int   `boil:"user_id"`
	LogSection types.JSON `boil:"vendor_log_section"`
	OldValue   types.JSON `boil:"old_value"`
	NewValue   types.JSON `boil:"new_value"`
	CreatedAt  time.Time  `boil:"created_at"`
	DeletedAt  null.Time  `boil:"deleted_at"`
}

func NewVendorLog(id int, vendorID int, userID null.Int, logSection types.JSON, oldValue types.JSON, newValue types.JSON, createdAt time.Time, deletedAt null.Time) *VendorLog {
	return &VendorLog{
		ID:         id,
		VendorID:   vendorID,
		UserID:     userID,
		LogSection: logSection,
		OldValue:   oldValue,
		NewValue:   newValue,
		CreatedAt:  createdAt,
		DeletedAt:  deletedAt,
	}
}

func (l *VendorLog) GetID() int {
	return l.ID
}

func (l *VendorLog) FindAllVendorLog() []*VendorLog {
	return []*VendorLog{l}
}
