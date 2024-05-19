package entities

import "time"

type Language struct {
	ID          int
	Code        string
	DisplayName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewLanguage(id int, code string, displayName string) *Language {
	return &Language{
		ID:          id,
		Code:        code,
		DisplayName: displayName,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Now(),
	}
}

// GetID returns the ID of the Language.
//
// This function returns the ID of the Language.
// The returned value is of type int.
func (l *Language) GetID() int {
	// Returns the ID of the Language.
	return l.ID
}

// GetDetails returns the ID and code of the Language.
//
// This function returns the ID and code of the Language.
// The returned values are of type int and string respectively.
func (l *Language) GetDetails() (int, string) {
	// Returns the ID and code of the Language.
	return l.ID, l.Code
}
