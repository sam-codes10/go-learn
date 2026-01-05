package models

type SignUpRes struct {
	AuthToken string `json:"authToken"`
	Email     string `json:"email"`
}
