package jonas_chorum_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	jonas_chorum "github.com/omniboost/go-jonas_chorum"
)

var (
	client *jonas_chorum.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	partnerToken := os.Getenv("PARTNER_TOKEN")
	hotelCode := os.Getenv("HOTEL_CODE")
	partnerCode := os.Getenv("PARTNER_CODE")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL

	client = jonas_chorum.NewClient(nil)
	if debug != "" {
		client.SetDebug(true)
	}

	client.SetPartnerToken(partnerToken)
	client.SetHotelCode(hotelCode)
	client.SetPartnerCode(partnerCode)

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	m.Run()
}
