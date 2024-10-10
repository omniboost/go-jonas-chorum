package jonas_chorum

import "net/url"

type Request interface {
	Method() string
	QueryParams() QueryParams
	PathParams() PathParams
	RequestHeaderInterface() interface{}
	RequestBodyInterface() interface{}
	URL() *url.URL
}

type QueryParams interface {
	ToURLValues() (url.Values, error)
}

type PathParams interface {
	Params() map[string]string
}
