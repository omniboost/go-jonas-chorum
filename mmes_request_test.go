package jonas_chorum_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestMMESRequest(t *testing.T) {
	req := client.NewMMESRequest()
	resp, err := req.Do()
	if err == nil {
		t.Errorf("TestMMESRequest should return an error")
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
