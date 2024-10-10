package jonas_chorum

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/xml"
	"io"
	"log"

	"github.com/pkg/errors"
)

func (c *Client) NewGetDailyFinancialSummaryRequest() GetDailyFinancialSummaryRequest {
	return GetDailyFinancialSummaryRequest{
		client: c,

		Header: RequestHeader{
			BucketType: "GetDailyFinancialSummaryRQ",
			APIType:    "APIType",
			APIVersion: "1.0",
		},
	}
}

type GetDailyFinancialSummaryRequest struct {
	client *Client

	XMLName    xml.Name          `xml:"win:theRequest"`
	Header     RequestHeader     `xml:"Header"`
	Parameters RequestParameters `xml:"Parameters"`
	Body       struct {
		BusinessDate string `xml:"BusinessDate"`
	}
}

func (r GetDailyFinancialSummaryRequest) Do() (GetDailyFinancialSummaryResponseBody, error) {
	body := GetDailyFinancialSummaryResponseBody{}

	soapAction := r.client.NewMMESRequest()
	soapActionBody := soapAction.RequestBody()
	soapActionBody.Contents = r

	resp, err := soapAction.Do()
	if err != nil {
		return body, errors.WithStack(err)
	}

	// error handling...
	// xml decode into struct

	// base64decode
	b, err := base64.StdEncoding.DecodeString(resp.MMESResult)
	if err != nil {
		return body, errors.WithStack(err)
	}

	// gzip decode
	buf := bytes.NewBuffer(b)
	reader, err := gzip.NewReader(buf)
	if err != nil {
		return body, errors.WithStack(err)
	}

	b, err = io.ReadAll(reader)
	log.Println(err)
	log.Fatal(string(b))
	return body, nil
}

type GetDailyFinancialSummaryResponseBody struct{}
