package models

type User struct {
	DisplayName string `json:"username"`
	Email       string `json:"email"`
	PhotoURL    string `json:"photo_url"`
}
