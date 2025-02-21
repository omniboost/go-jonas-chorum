package jonas_chorum

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

func (c *Client) NewGetDailyTransactionsRequest() GetDailyTransactionsRequest {
	return GetDailyTransactionsRequest{
		client: c,

		Header: JCHeader{
			BucketType: "GetDailyTransactionsRQ",
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

type GetDailyTransactionsRequest struct {
	client *Client

	XMLName    xml.Name     `xml:"Content"`
	Header     JCHeader     `xml:"Header"`
	Parameters JCParameters `xml:"Parameters"`
	Body       struct {
		BusinessDate string `xml:"BusinessDate"`
	}
}

func (r GetDailyTransactionsRequest) Do() (GetDailyTransactionsResponseBody, error) {
	responseBody := GetDailyTransactionsResponseBody{}

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

type GetDailyTransactionsResponseBody struct {
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
		BusinessDate string `xml:"BusinessDate"`
		Transactions struct {
			Transaction []struct {
				AccountType             string `xml:"AccountType"`
				FolioId                 string `xml:"FolioId"`
				PMSConfirmationNumber   string `xml:"PMSConfirmationNumber"`
				GroupConfirmationNumber string `xml:"GroupConfirmationNumber"`
				MemberNumber            string `xml:"MemberNumber"`
				SponsorMemberNumber     string `xml:"SponsorMemberNumber"`
				TransactionID           string `xml:"TransactionID"`
				Description             string `xml:"Description"`
				DepartmentCode          string `xml:"DepartmentCode"`
				GLAccountNumber         string `xml:"GLAccountNumber"`
				Amount                  string `xml:"Amount"`
				CreditDebit             string `xml:"CreditDebit"`
				AcctgGroup              string `xml:"AcctgGroup"`
				PostingType             string `xml:"PostingType"`
				LinkedTransactionID     string `xml:"LinkedTransactionID"`
				RoomNumber              string `xml:"RoomNumber"`
				GuestFirstName          string `xml:"GuestFirstName"`
				GuestLastName           string `xml:"GuestLastName"`
			} `xml:"Transaction"`
		} `xml:"Transactions"`
	} `xml:"Body"`
}
