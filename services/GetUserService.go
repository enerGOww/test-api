package services

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"test-api/repositories"
	"time"
)

type UsersData struct {
	Data []struct {
		Gender string
		Name   struct {
			First string
			Last  string
		}
		Location struct {
			Postcode float64
		}
		Registered struct {
			Date time.Time
		}
	} `json:"results"`
}

func GetUsers() *UsersData {
	buffer, err := repositories.FetchUsers()
	if err != nil {
		log.Error(err)
		return nil
	}

	var users *UsersData
	err = json.Unmarshal(buffer, &users)
	if err != nil {
		log.Error(err)
	}

	return users
}

func FilterUsers(timeFrom, timeTo time.Time, users *UsersData) *UsersData {
	var result UsersData
	for _, user := range users.Data {
		if user.Registered.Date.After(timeFrom) && user.Registered.Date.Before(timeTo) {
			result.Data = append(result.Data, user)
		}
	}

	return &result
}
