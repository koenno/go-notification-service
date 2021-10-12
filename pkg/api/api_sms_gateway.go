package api

type SMS struct {
	TelephoneNumber string `json:"telephoneNumber"`
	Message         string `json:"message"`
}
