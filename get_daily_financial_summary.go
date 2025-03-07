package jonas_chorum

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

func (c *Client) NewGetDailyFinancialSummaryRequest() GetDailyFinancialSummaryRequest {
	return GetDailyFinancialSummaryRequest{
		client: c,

		Header: JCHeader{
			BucketType: "GetDailyFinancialSummaryRQ",
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

type GetDailyFinancialSummaryRequest struct {
	client *Client

	XMLName    xml.Name     `xml:"Content"`
	Header     JCHeader     `xml:"Header"`
	Parameters JCParameters `xml:"Parameters"`
	Body       struct {
		BusinessDate Date `xml:"BusinessDate"`
	}
}

func (r GetDailyFinancialSummaryRequest) Do() (GetDailyFinancialSummaryResponseBody, error) {
	responseBody := GetDailyFinancialSummaryResponseBody{}

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

type GetDailyFinancialSummaryResponseBody struct {
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
		DepartmentCodeAmounts struct {
			DepartmentCodeAmount []struct {
				DepartmentCode string `xml:"DepartmentCode"`
				Description    string `xml:"Description"`
				GLAccount      string `xml:"GLAccount"`
				Amount         string `xml:"Amount"`
				CreditDebit    string `xml:"CreditDebit"`
				PostingType    string `xml:"PostingType"`
			} `xml:"DepartmentCodeAmount"`
		} `xml:"DepartmentCodeAmounts"`
		LedgerAmounts struct {
			LedgerAmount []struct {
				LedgerCode  string `xml:"LedgerCode"`
				Description string `xml:"Description"`
				GLAccount   string `xml:"GLAccount"`
				Amount      string `xml:"Amount"`
			} `xml:"LedgerAmount"`
		} `xml:"LedgerAmounts"`
		Status string `xml:"Status"`
	} `xml:"Body"`
}
