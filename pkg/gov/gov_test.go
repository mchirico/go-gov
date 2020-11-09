package gov

import (
	"fmt"
	"testing"
)

func GetGOVs() (GOV, GOV, error) {
	senate := "https://api.propublica.org/congress/v1/116/senate/members.json"
	s, err := GetGov(senate)
	if err != nil {
		return s, s, err
	}

	house := "https://api.propublica.org/congress/v1/116/house/members.json"
	h, err := GetGov(house)
	if err != nil {
		return h, h, err
	}
	return h, s, err

}

func TestGetGov(t *testing.T) {
	house, senate, err := GetGOVs()
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	fmt.Printf("house: %v\nsenate: %v\n",
		house.Results[0].Members[0].APIURI,
		senate.Results[0].Members[0].APIURI,
	)

}

func TestGetSubcommittees(t *testing.T) {
	house, senate, err := GetGOVs()
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	_, err = GetSubcommittees(house.Results[0].Members[0].APIURI)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	_, err = GetSubcommittees(senate.Results[0].Members[0].APIURI)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}
