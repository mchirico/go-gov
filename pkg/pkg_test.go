package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/mchirico/go-gov/pkg/gov"
	"github.com/mchirico/go-gov/pkg/httputils"
	"testing"
)

func TestGov(t *testing.T) {
	url := "https://api.propublica.org/congress/v1/116/senate/members.json"

	key := "X-API-Key"
	value := "1vtlJSvzaaB6bTjJKzyakYnjnxrRzM22Ex3j2SDR"

	httputils.Header(key,value)


	r, err := httputils.Get(url)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}



	var gov gov.Gov
	err = json.Unmarshal(r, &gov)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Printf("\n%+v", gov)

	fmt.Println(gov.Results[0].Members[0].FacebookAccount)


}
