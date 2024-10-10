package jonas_chorum

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-jonas_chorum/utils"
)

func (c *Client) NewGetDailyFinancialSummary() GetDailyFinancialSummary {
	r := GetDailyFinancialSummary{
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

type GetDailyFinancialSummary struct {
	client        *Client
	queryParams   *GetDailyFinancialSummaryQueryParams
	pathParams    *GetDailyFinancialSummaryPathParams
	method        string
	headers       http.Header
	requestBody   GetDailyFinancialSummaryBody
	requestHeader GetDailyFinancialSummaryHeader
}

func (r GetDailyFinancialSummary) NewQueryParams() *GetDailyFinancialSummaryQueryParams {
	return &GetDailyFinancialSummaryQueryParams{}
}

type GetDailyFinancialSummaryQueryParams struct {
}

func (p GetDailyFinancialSummaryQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetDailyFinancialSummary) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetDailyFinancialSummary) NewPathParams() *GetDailyFinancialSummaryPathParams {
	return &GetDailyFinancialSummaryPathParams{}
}

type GetDailyFinancialSummaryPathParams struct {
}

func (p *GetDailyFinancialSummaryPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetDailyFinancialSummary) PathParams() PathParams {
	return r.pathParams
}

func (r *GetDailyFinancialSummary) SetMethod(method string) {
	r.method = method
}

func (r *GetDailyFinancialSummary) Method() string {
	return r.method
}

func (r GetDailyFinancialSummary) NewRequestHeader() GetDailyFinancialSummaryRequestHeader {
	return GetDailyFinancialSummaryRequestHeader{}
}

func (r *GetDailyFinancialSummary) RequestHeader() *GetDailyFinancialSummaryRequestHeader {
	return &r.requestHeader
}

func (r *GetDailyFinancialSummary) RequestHeaderInterface() interface{} {
	return &r.requestHeader
}

type GetDailyFinancialSummaryRequestHeader struct{}

func (r GetDailyFinancialSummary) NewRequestBody() GetDailyFinancialSummaryRequestBody {
	return GetDailyFinancialSummaryRequestBody{
		ProfileRequestor: ProfileRequestor{
			AuthenticationMethod: "PD",
		},
	}
}

type GetDailyFinancialSummaryRequestBody struct {
	XMLName          xml.Name         `xml:"http://tempuri.org/RLXSOAP19/RLXSOAP19 pmsprf_GetProfileSummaryWithAttributes"`
	SessionID        string           `xml:"SessionId"`
	ProfileRequestor ProfileRequestor `xml:"ProfileRequestor,omitempty"`
}

func (r *GetDailyFinancialSummary) RequestBody() *GetDailyFinancialSummaryRequestBody {
	return &r.requestBody
}

func (r *GetDailyFinancialSummary) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetDailyFinancialSummary) SetRequestBody(body GetDailyFinancialSummaryRequestBody) {
	r.requestBody = body
}

func (r *GetDailyFinancialSummary) NewResponseBody() *GetDailyFinancialSummaryResponseBody {
	return &GetDailyFinancialSummaryResponseBody{}
}

type GetDailyFinancialSummaryResponseBody struct {
	XMLName                                     xml.Name                `xml:GetProfileSummaryWithAttributesResponse`
	PmsintGetProfileSummaryWithAttributesResult ExceptionBlock          `xml:"pmsprf_GetProfileSummaryWithAttributesResult"`
	ProfileSummary                              ProfileSummary          `xml:"Profile>Profile"`
	CustomAttributes                            ProfileCustomAttributes `xml:"Profile>CustomAttributes>ProfileCustomAttributes"`
}

func (rb GetDailyFinancialSummaryResponseBody) ExceptionBlock() ExceptionBlock {
	return rb.PmsintGetProfileSummaryWithAttributesResult
}

func (r *GetDailyFinancialSummary) URL() *url.URL {
	u := r.client.GetEndpointURL("rlxsoap.asmx?op=pmsprf_DailyFinancialSummary", r.PathParams())
	return &u
}

func (r *GetDailyFinancialSummary) Do() (GetDailyFinancialSummaryResponseBody, error) {
	var err error

	// fetch a new token if it isn't set already
	r.requestBody.SessionID, err = r.client.SessionID()
	if err != nil {
		return *r.NewResponseBody(), err
	}

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

type ProfileCustomAttributes []ProfileCustomAttribute

type ProfileCustomAttribute struct {
	AttributeCode               string `xml:"AttributeCode"`
	Value                       string `xml:"Value"`
	Param1                      string `xml:"Param1"`
	Param2                      string `xml:"Param2"`
	ProfileAttributeCode        string `xml:"ProfileAttributeCode"`
	ProfileAttributeDescription string `xml:"ProfileAttributeDescription"`
	ProfileAttributeValue       string `xml:"ProfileAttributeValue"`
}
