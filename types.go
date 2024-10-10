package jonas_chorum

type RequestHeader struct {
	BucketType    string `xml:"BucketType"`
	APIType       string `xml:"APIType"`
	APIVersion    string `xml:"APIVersion"`
	SecurityToken string `xml:"SecurityToken"`
}

type RequestParameters struct {
	// A unique code for the partner system
	// Provided by Jonas Chorum
	PartnerCode string `xml:"PartnerCode"`

	// A unique code for the hotel
	// Used by webservice layer to validate customer is active and determine where message needs to be sent
	// Provided by Jonas Chorum
	HotelCode string `xml:"HotelCode"`

	// Provides access privileges to connected PMS systems
	// Provided by Jonas Chorum
	PartnerToken string `xml:"PartnerToken"`
	EchoToken    string `xml:"EchoToken,omitempty"`
}

// <!DOCTYPE MsiXmlBucket []>
// <Content>
//   <Header>
//     <BucketType>ProcessRequestBucketError</BucketType>
//     <APIType>NA</APIType>
//     <APIVersion>NA</APIVersion>
//     <SecurityToken>NA</SecurityToken>
//     <Internal>NA</Internal>
//     <CustomDataA></CustomDataA>
//     <CustomDataB></CustomDataB>
//     <CustomDataC></CustomDataC>
//     <CustomDataD></CustomDataD>
//   </Header>
//   <Parameters>
//     <Error>Value cannot be null.
// Parameter name: s</Error>
//   </Parameters>
//   <Body />
// </Content>
