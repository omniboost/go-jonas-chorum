package jonas_chorum

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

func (c *Client) NewRetrieveDepartmentCodesRequest() RetrieveDepartmentCodesRequest {
	return RetrieveDepartmentCodesRequest{
		client: c,

		Header: JCHeader{
			BucketType: "RetrieveDepartmentCodesRQ",
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

type RetrieveDepartmentCodesRequest struct {
	client *Client

	XMLName    xml.Name     `xml:"Content"`
	Header     JCHeader     `xml:"Header"`
	Parameters JCParameters `xml:"Parameters"`
	Body       struct {
		BusinessDate Date `xml:"BusinessDate"`
	}
}

func (r RetrieveDepartmentCodesRequest) Do() (RetrieveDepartmentCodesResponseBody, error) {
	responseBody := RetrieveDepartmentCodesResponseBody{}

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

type RetrieveDepartmentCodesResponseBody struct {
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
		DepartmentCodes struct {
			DepartmentCode []struct {
				Code                 string `xml:"Code"`
				Description          string `xml:"Description"`
				GeneralLedgerAccount string `xml:"GeneralLedgerAccount"`
				Type                 string `xml:"Type"`
				TaxCodes             []struct {
					TaxCode struct {
						Code        string `xml:"Code"`
						IsInculsive string `xml:"IsInculsive"`
					} `xml:"TaxCode"`
				} `xml:"TaxCodes"`
			} `xml:"DepartmentCode"`
		} `xml:"DepartmentCodes"`
	} `xml:"Body"`
}
