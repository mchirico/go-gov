package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/mchirico/go-gov/pkg/gov/senate"
	"github.com/mchirico/go-gov/pkg/httputils"
	"testing"
)

func TestSenate(t *testing.T) {
	url := "https://api.propublica.org/congress/v1/116/senate/members.json"

	key := "X-API-Key"
	value := "1vtlJSvzaaB6bTjJKzyakYnjnxrRzM22Ex3j2SDR"

	h := httputils.NewHTTP()
	h.Header(key, value)

	r, err := h.Get(url)

	var gov senate.Senate
	err = json.Unmarshal(r, &gov)
	if err != nil {
		fmt.Println("error:", err)
	}

	m := map[string]int{}
	for _, v := range gov.Results[0].Members {
		m[v.State] += 1
	}

	for k, v := range m {
		if v < 2 {
			t.Fatalf("Data bad: %v,%v\n", k, v)
		}
	}

	fmt.Println(gov.Results[0].Members[0].FacebookAccount)

}
