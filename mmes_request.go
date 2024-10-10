package jonas_chorum

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-jonas_chorum/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewMMESRequest() MMESRequest {
	r := MMESRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	r.requestHeader = r.NewRequestHeader()
	return r
}

type MMESRequest struct {
	client        *Client
	queryParams   *MMESRequestQueryParams
	pathParams    *MMESRequestPathParams
	method        string
	headers       http.Header
	requestBody   MMESRequestBody
	requestHeader MMESRequestHeader
}

func (r MMESRequest) SOAPAction() string {
	return "https://winpm.com/MMES"
}

func (r MMESRequest) NewQueryParams() *MMESRequestQueryParams {
	return &MMESRequestQueryParams{}
}

type MMESRequestQueryParams struct {
}

func (p MMESRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *MMESRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r MMESRequest) NewPathParams() *MMESRequestPathParams {
	return &MMESRequestPathParams{}
}

type MMESRequestPathParams struct {
}

func (p *MMESRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *MMESRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *MMESRequest) SetMethod(method string) {
	r.method = method
}

func (r *MMESRequest) Method() string {
	return r.method
}

func (r MMESRequest) NewRequestHeader() MMESRequestHeader {
	return MMESRequestHeader{}
}

func (r *MMESRequest) RequestHeader() *MMESRequestHeader {
	return &r.requestHeader
}

func (r *MMESRequest) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type MMESRequestHeader struct{}

func (r MMESRequest) NewRequestBody() MMESRequestBody {
	return MMESRequestBody{}
}

type MMESRequestBody struct {
	XMLName    xml.Name            `xml:"https://winpm.com/ MMES"`
	TheRequest MMESRequestContents `xml:"theRequest"`
}

type MMESRequestContents struct {
	Contents any `xml:",any"`
}

func (b MMESRequestContents) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "theRequest"

	contents, err := xml.Marshal(b.Contents)
	if err != nil {
		return errors.WithStack(err)
	}

	// first gzip b.Contents
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err = w.Write(contents)
	if err != nil {
		return errors.WithStack(err)
	}

	err = w.Close()
	if err != nil {
		return errors.WithStack(err)
	}

	// base64 decode buf.Bytes()
	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	// use gzipped, based64 encoded, xml marshalled struct in encoding
	return e.EncodeElement(b64, start)
}

func (r *MMESRequest) RequestBody() *MMESRequestBody {
	return &r.requestBody
}

func (r *MMESRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *MMESRequest) SetRequestBody(body MMESRequestBody) {
	r.requestBody = body
}

func (r *MMESRequest) NewResponseBody() *MMESRequestResponseBody {
	return &MMESRequestResponseBody{}
}

type MMESRequestResponseBody struct {
	XMLName xml.Name `xml:"MMESResponse"`

	MMESResult string `xml:"MMESResult"`
}

// func (rb MMESRequestResponseBody) ExceptionBlock() ExceptionBlock {
// 	return rb.PmsintGetProfileSummaryWithAttributesResult
// }

func (r *MMESRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *MMESRequest) Do() (MMESRequestResponseBody, error) {
	var err error

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
