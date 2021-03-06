package httputils

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://api.propublica.org/congress/v1/116/senate/members.json"

	key := "X-API-Key"
	value := "1vtlJSvzaaB6bTjJKzyakYnjnxrRzM22Ex3j2SDR"

	h := NewHTTP()
	h.Header(key, value)

	r, err := h.Get(url)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	fmt.Println(string(r))

}
