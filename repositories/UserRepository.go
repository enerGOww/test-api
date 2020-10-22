package repositories

import (
	"io/ioutil"
	"net/http"
)

func FetchUsers() ([]byte, error) {
	res, err := http.Get("https://randomuser.me/api?inc=gender,name,location,registered&noinfo&results=4")
	if err != nil {
		return nil, err
	}

	users, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return users, nil
}
