package analysis

import (
	"encoding/json"
	"fmt"
	"github.com/mchirico/go-gov/pkg/gov/house"
	"github.com/mchirico/go-gov/pkg/httputils"
)

func GetHouseMembers() (house.House, error) {
	url := "https://api.propublica.org/congress/v1/116/senate/members.json"

	key := "X-API-Key"
	value := "1vtlJSvzaaB6bTjJKzyakYnjnxrRzM22Ex3j2SDR"

	h := httputils.NewHTTP()
	h.Header(key, value)

	r, err := h.Get(url)
	if err != nil {
		return house.House{}, err
	}

	var gov house.House
	err = json.Unmarshal(r, &gov)
	if err != nil {
		fmt.Println("error:", err)
	}
	return gov, err

}
