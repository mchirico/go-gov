package gov

import (
	"fmt"
	"testing"
)

func GetGOVs() (*Gov, GOV, GOV, error) {
	g := NewGov()
	senate := "https://api.propublica.org/congress/v1/116/senate/members.json"
	s, err := g.GetGov(senate)
	if err != nil {
		return g, s, s, err
	}

	house := "https://api.propublica.org/congress/v1/116/house/members.json"
	h, err := g.GetGov(house)
	if err != nil {
		return g, h, h, err
	}
	return g, h, s, err

}

func TestGetGov(t *testing.T) {
	_, house, senate, err := GetGOVs()
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	fmt.Printf("house: %v\nsenate: %v\n",
		house.Results[0].Members[0].APIURI,
		senate.Results[0].Members[0].APIURI,
	)

}

func TestGetSubcommittees(t *testing.T) {

	g, house, senate, err := GetGOVs()
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	_, err = g.GetSubcommittees(house.Results[0].Members[0].APIURI)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	_, err = g.GetSubcommittees(senate.Results[0].Members[0].APIURI)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}
