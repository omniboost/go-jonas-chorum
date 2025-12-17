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
		BusinessDate          Date    `xml:"BusinessDate"`
		NetRoomRevenue        float64 `xml:"NetRoomRevenue"`
		NetOtherRevenue       float64 `xml:"NetOtherRevenue"`
		NetTotalRevenue       float64 `xml:"NetTotalRevenue"`
		TotalRooms            int     `xml:"TotalRooms"`
		VacantRooms           int     `xml:"VacantRooms"`
		TotalOccupiedRooms    int     `xml:"TotalOccupiedRooms"`
		GroupRoomsOccupied    int     `xml:"GroupRoomsOccupied"`
		OOORooms              int     `xml:"OOORooms"`
		DayUseRooms           int     `xml:"DayUseRooms"`
		GroupRoomsNotPickedUp int     `xml:"GroupRoomsNotPickedUp"`
		Arrivals              int     `xml:"Arrivals"`
		GroupArrivals         int     `xml:"GroupArrivals"`
		Departures            int     `xml:"Departures"`
		GroupDepartures       int     `xml:"GroupDepartures"`
		WalkInArrivals        int     `xml:"WalkInArrivals"`
		CompRooms             int     `xml:"CompRooms"`
		NoShows               int     `xml:"NoShows"`
		AverageLOS            int     `xml:"AverageLOS"`
		AdultGuestCount       int     `xml:"AdultGuestCount"`
		ChildOverCount        int     `xml:"ChildOverCount"`
		ChildUnderCount       int     `xml:"ChildUnderCount"`
		RoomsSold             int     `xml:"RoomsSold"`
		Occupancy             float64 `xml:"Occupancy"`
		OccupancyWithoutOOO   float64 `xml:"OccupancyWithoutOOO"`
		ESOC                  float64 `xml:"ESOC"`
		NetADR                float64 `xml:"NetADR"`
		NetAWR                float64 `xml:"NetAWR"`
		GrossADR              float64 `xml:"GrossADR"`
		GrossAWR              float64 `xml:"GrossAWR"`
		NetADRWithComps       float64 `xml:"NetADRWithComps"`
		GrossADRWithComps     float64 `xml:"GrossADRWithComps"`
		NetRevPAR             float64 `xml:"NetRevPAR"`
		NetRevPAW             float64 `xml:"NetRevPAW"`
		GrossRevPAR           float64 `xml:"GrossRevPAR"`
		GrossRevPAW           float64 `xml:"GrossRevPAW"`
		NetFBRevenue          float64 `xml:"NetFBRevenue"`
		AverageBookingWindow  float64 `xml:"AverageBookingWindow"`
	} `xml:"Body"`
}
