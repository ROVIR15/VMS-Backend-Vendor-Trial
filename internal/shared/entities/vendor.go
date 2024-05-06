package entities

type Vendor struct {
	id         int     `json:"id"`
	name       string  `json:"name"`
	email      string  `json:"email"`
	commission float32 `json:"commission"`
}
