package house

import (
	"encoding/json"
	"fmt"
	"github.com/mchirico/go-gov/pkg/httputils"
	"testing"
)

func TestHouse(t *testing.T) {
	url := "https://api.propublica.org/congress/v1/116/house/members.json"

	key := "X-API-Key"
	value := "1vtlJSvzaaB6bTjJKzyakYnjnxrRzM22Ex3j2SDR"

	h := httputils.NewHTTP()
	h.Header(key, value)

	r, err := h.Get(url)

	var gov House
	err = json.Unmarshal(r, &gov)
	if err != nil {
		fmt.Println("error:", err)
	}

	m := map[string]int{}
	for _, v := range gov.Results[0].Members {
		m[v.State] += 1
	}

	for k, v := range m {
		if v < 1 {
			t.Fatalf("Data bad: %v,%v\n", k, v)
		}
	}

	fmt.Println(gov.Results[0].Members[0].FacebookAccount)

}

func TestGetHouse(t *testing.T) {
	gov, err := GetHouse()
	if err != nil {
		t.Fatalf("We got error: %s\n", err)
	}
	fmt.Printf("%v\n", gov.Results[0].Members[0])
}

func TestSubcommittees(t *testing.T) {
   url := "https://api.propublica.org/congress/v1/members/A000374.json"
   sub, err := GetSubcommittees(url)
   if err != nil {
   	t.Fatalf("Didn't work: %s\n",err)
   }
   fmt.Printf("sub: %v\n",sub)

}
