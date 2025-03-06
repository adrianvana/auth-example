package models

type RandomUser struct {
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type RandomUserResponse struct {
	Results []RandomUser `json:"results"`
}
