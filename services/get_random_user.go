package services

import (
	"encoding/json"
	"net/http"
)

type RandomUserResponse struct {
	Results []struct {
		Gender string `json:"gender"`
		Name   struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Email string `json:"email"`
	} `json:"results"`
}

func GetRandomUser() (*RandomUserResponse, error) {

	res, err := http.Get("https://randomuser.me/api")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var user RandomUserResponse
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
