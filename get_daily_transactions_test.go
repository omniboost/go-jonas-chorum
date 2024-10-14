package jonas_chorum_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetDailyTransactions(t *testing.T) {
	req := client.NewGetDailyTransactionsRequest()
	req.Parameters.PartnerCode = client.PartnerCode()
	req.Parameters.HotelCode = client.HotelCode()
	req.Parameters.PartnerToken = client.PartnerToken()
	req.Body.BusinessDate = "2024-10-01"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
