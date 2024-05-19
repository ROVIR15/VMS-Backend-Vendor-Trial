package entities

import "time"

type HostContract struct {
	ID            int
	HostPartnerId int
	URL           string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	StartPeriod   time.Time
	EndDate       time.Time
}

// NewHostContract creates a new instance of HostContract.
//
// The id parameter is the unique identifier of the host contract.
// The hostPartnerId parameter is the unique identifier of the host partner.
// The url parameter is the URL of the host contract.
// The startPeriod parameter is the start period of the host contract.
// The endDate parameter is the end date of the host contract.
//
// Returns a pointer to the newly created HostContract.
func NewHostContract(id int, hostPartnerId int, url string, startPeriod time.Time, endDate time.Time) *HostContract {
	currentTime := time.Now()
	return &HostContract{
		ID:            id,
		HostPartnerId: hostPartnerId,
		URL:           url,
		CreatedAt:     currentTime,
		UpdatedAt:     currentTime,
		DeletedAt:     nil,
		StartPeriod:   startPeriod,
		EndDate:       endDate,
	}
}

// FindOneHostContract returns the HostContract that matches the given id, or an error if no such contract exists.
func (h *HostContract) FindOneHostContract(id int) (*HostContract, *NotFoundError) {
	if h.ID == id {
		return h, nil
	}
	return nil, &NotFoundError{Entity: "HostContract", ID: id}
}

// FindAllHostContractsByHostPartnerID returns all HostContract instances that match the given hostPartnerId.
func (h *HostContract) FindAllHostContractsByHostPartnerID(hostPartnerId int) []*HostContract {
	if h.HostPartnerId == hostPartnerId {
		return []*HostContract{h}
	}
	return []*HostContract{}
}

// GetID returns the ID of the HostContract.
//
// Returns:
// - int: The ID of the HostContract.
func (h *HostContract) GetID() int {
	return h.ID
}
