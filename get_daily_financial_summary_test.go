package jonas_chorum_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetDailyFinancialSummary(t *testing.T) {
	req := client.NewGetDailyFinancialSummaryRequest()
	req.Parameters.PartnerCode = client.PartnerCode()
	req.Parameters.HotelCode = client.HotelCode()
	req.Parameters.PartnerToken = client.PartnerToken()
	req.Body.BusinessDate = "03-10-2022"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
