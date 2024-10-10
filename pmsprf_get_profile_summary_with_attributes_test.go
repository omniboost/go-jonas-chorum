package jonas_chorum_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetProfileSummaryWithAttributes(t *testing.T) {
	req := client.NewGetProfileSummaryWithAttributesRequest()
	req.RequestBody().ProfileRequestor.ProfileUniqueID = "PF003033"
	req.RequestBody().ProfileRequestor.AuthenticationMethod = "PD"
	req.RequestBody().ProfileRequestor.AuthenticationCode = "Surname"
	req.RequestBody().ProfileRequestor.AuthenticationValue = "Bierreth"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
