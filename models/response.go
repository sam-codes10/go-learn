package models

type AuthRes struct {
	AuthToken string `json:"authToken"`
	Email     string `json:"email"`
}
