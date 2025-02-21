package jonas_chorum_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	jonas_chorum "github.com/omniboost/go-jonas-chorum"
)

func TestGetDailyStatistics(t *testing.T) {
	req := client.NewGetDailyStatisticsRequest()
	req.Parameters.PartnerCode = client.PartnerCode()
	req.Parameters.HotelCode = client.HotelCode()
	req.Parameters.PartnerToken = client.PartnerToken()
	req.Body.BusinessDate = jonas_chorum.Date{time.Date(2022, 10, 03, 0, 0, 0, 0, time.UTC)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
