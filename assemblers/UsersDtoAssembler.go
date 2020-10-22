package assemblers

import (
	"test-api/services"
	"time"
)

type FormattedUsersData struct {
	Data []*User `json:"data"`
}

type User struct {
	Gender    string    `json:"gender"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Postcode  float64   `json:"postcode"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatUsersDataDto(dto *services.UsersData) FormattedUsersData {
	var result FormattedUsersData

	for _, user := range dto.Data {
		formatUser := &User{
			Gender:    user.Gender,
			FirstName: user.Name.First,
			LastName:  user.Name.Last,
			Postcode:  user.Location.Postcode,
			CreatedAt: user.Registered.Date,
		}
		result.Data = append(result.Data, formatUser)
	}

	return result
}
