package jonas_chorum

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

func (c *Client) NewGetDailyStatisticsRequest() GetDailyStatisticsRequest {
	return GetDailyStatisticsRequest{
		client: c,

		Header: JCHeader{
			BucketType: "GetDailyStatisticsRQ",
			APIType:    "APIType",
			APIVersion: "1.0",
		},
		Parameters: JCParameters{
			PartnerCode:  c.PartnerCode(),
			HotelCode:    c.HotelCode(),
			PartnerToken: c.PartnerToken(),
		},
	}
}

type GetDailyStatisticsRequest struct {
	client *Client

	XMLName    xml.Name     `xml:"Content"`
	Header     JCHeader     `xml:"Header"`
	Parameters JCParameters `xml:"Parameters"`
	Body       struct {
		BusinessDate Date `xml:"BusinessDate"`
	}
}

func (r GetDailyStatisticsRequest) Do() (GetDailyStatisticsResponseBody, error) {
	responseBody := GetDailyStatisticsResponseBody{}

	// setup underlying soap request
	soapAction := r.client.NewMMESRequest()
	soapActionBody := soapAction.RequestBody()
	soapActionBody.TheRequest.Contents = r

	// execute soap request
	resp, err := soapAction.Do()
	if err != nil {
		return responseBody, errors.WithStack(err)
	}

	// use inner body of underlying soap response as our response body
	reader, err := resp.MMESResult.Decode()
	if err != nil {
		return responseBody, errors.WithStack(err)
	}

	// bodyFailure is already checked in underlying soap request (MMESRequest)
	// so only check for wanted response body
	err = r.client.Unmarshal(reader, []any{&responseBody}, []any{})
	if err != nil {
		return responseBody, errors.WithStack(err)
	}

	// else everything should be fine
	return responseBody, nil
}

type GetDailyStatisticsResponseBody struct {
	XMLName xml.Name `xml:"Content"`
	Header  struct {
		BucketType    string `xml:"BucketType"`
		APIType       string `xml:"APIType"`
		APIVersion    string `xml:"APIVersion"`
		SecurityToken string `xml:"SecurityToken"`
		Internal      string `xml:"Internal"`
		CustomDataA   string `xml:"CustomDataA"`
		CustomDataB   string `xml:"CustomDataB"`
		CustomDataC   string `xml:"CustomDataC"`
		CustomDataD   string `xml:"CustomDataD"`
	} `xml:"Header"`
	Parameters struct {
		HotelCode    string `xml:"HotelCode"`
		PartnerCode  string `xml:"PartnerCode"`
		EchoToken    string `xml:"EchoToken"`
		PartnerToken string `xml:"PartnerToken"`
	} `xml:"Parameters"`
	Body struct {
		BusinessDate          string `xml:"BusinessDate"`
		NetRoomRevenue        string `xml:"NetRoomRevenue"`
		NetOtherRevenue       string `xml:"NetOtherRevenue"`
		NetTotalRevenue       string `xml:"NetTotalRevenue"`
		TotalRooms            string `xml:"TotalRooms"`
		VacantRooms           string `xml:"VacantRooms"`
		TotalOccupiedRooms    string `xml:"TotalOccupiedRooms"`
		GroupRoomsOccupied    string `xml:"GroupRoomsOccupied"`
		OOORooms              string `xml:"OOORooms"`
		DayUseRooms           string `xml:"DayUseRooms"`
		GroupRoomsNotPickedUp string `xml:"GroupRoomsNotPickedUp"`
		Arrivals              string `xml:"Arrivals"`
		GroupArrivals         string `xml:"GroupArrivals"`
		Departures            string `xml:"Departures"`
		GroupDepartures       string `xml:"GroupDepartures"`
		WalkInArrivals        string `xml:"WalkInArrivals"`
		CompRooms             string `xml:"CompRooms"`
		NoShows               string `xml:"NoShows"`
		AverageLOS            string `xml:"AverageLOS"`
		AdultGuestCount       string `xml:"AdultGuestCount"`
		ChildOverCount        string `xml:"ChildOverCount"`
		ChildUnderCount       string `xml:"ChildUnderCount"`
		RoomsSold             string `xml:"RoomsSold"`
		Occupancy             string `xml:"Occupancy"`
		OccupancyWithoutOOO   string `xml:"OccupancyWithoutOOO"`
		ESOC                  string `xml:"ESOC"`
		NetADR                string `xml:"NetADR"`
		NetAWR                string `xml:"NetAWR"`
		GrossADR              string `xml:"GrossADR"`
		GrossAWR              string `xml:"GrossAWR"`
		NetADRWithComps       string `xml:"NetADRWithComps"`
		GrossADRWithComps     string `xml:"GrossADRWithComps"`
		NetRevPAR             string `xml:"NetRevPAR"`
		NetRevPAW             string `xml:"NetRevPAW"`
		GrossRevPAR           string `xml:"GrossRevPAR"`
		GrossRevPAW           string `xml:"GrossRevPAW"`
		NetFBRevenue          string `xml:"NetFBRevenue"`
		AverageBookingWindow  string `xml:"AverageBookingWindow"`
	} `xml:"Body"`
}
