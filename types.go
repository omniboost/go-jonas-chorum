package jonas_chorum

type JCHeader struct {
	BucketType    string `xml:"BucketType"`
	APIType       string `xml:"APIType"`
	APIVersion    string `xml:"APIVersion"`
	SecurityToken string `xml:"SecurityToken"`
}

type JCParameters struct {
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

	Error string `xml:"Error"`
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

type DailyStatistics struct {
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
}

type DailyFinancialSummary struct {
	BusinessDate          string `xml:"BusinessDate"`
	DepartmentCodeAmounts []struct {
		DepartmentCode string  `xml:"DepartmentCode"`
		Description    string  `xml:"Description"`
		GLAccount      int     `xml:"GLAccount"`
		Amount         float64 `xml:"Amount"`
		CreditDebit    string  `xml:"CreditDebit"`
		PostingType    string  `xml:"PostingType"`
	} `xml:"DepartmentCodeAmounts>DepartmentCodeAmount"`
	LedgerAmounts []struct {
		LedgerCode  string  `xml:"LedgerCode"`
		Description string  `xml:"Description"`
		GLAccount   int     `xml:"GLAccount"`
		Amount      float64 `xml:"Amount"`
	} `xml:"LedgerAmounts>LedgerAmount"`
	Status string `xml:"Status"`
}

type DailyTransactions struct {
	BusinessDate Date `xml:"BusinessDate"`
	Transactions []struct {
		AccountType             string  `xml:"AccountType"`
		FolioId                 int     `xml:"FolioId"`
		PMSConfirmationNumber   int     `xml:"PMSConfirmationNumber"`
		GroupConfirmationNumber int     `xml:"GroupConfirmationNumber"`
		MemberNumber            string  `xml:"MemberNumber"`
		SponsorMemberNumber     string  `xml:"SponsorMemberNumber"`
		TransactionID           string  `xml:"TransactionID"`
		Description             string  `xml:"Description"`
		DepartmentCode          string  `xml:"DepartmentCode"`
		GLAccountNumber         *int    `xml:"GLAccountNumber"`
		Amount                  float64 `xml:"Amount"`
		CreditDebit             string  `xml:"CreditDebit"`
		AcctgGroup              string  `xml:"AcctgGroup"`
		PostingType             string  `xml:"PostingType"`
		LinkedTransactionID     string  `xml:"LinkedTransactionID"`
		RoomNumber              int     `xml:"RoomNumber"`
		GuestFirstName          string  `xml:"GuestFirstName"`
		GuestLastName           string  `xml:"GuestLastName"`
	} `xml:"Transactions>Transaction"`
}

type DailyFinancialsByMarketSegment struct {
	BusinessDate          string `xml:"BusinessDate"`
	DepartmentCodeAmounts []struct {
		DepartmentCode string  `xml:"DepartmentCode"`
		Description    string  `xml:"Description"`
		GLAccount      int     `xml:"GLAccount"`
		Amount         float64 `xml:"Amount"`
		CreditDebit    string  `xml:"CreditDebit"`
		PostingType    string  `xml:"PostingType"`
		MarketSegment  string  `xml:"MarketSegment"`
	} `xml:"DepartmentCodeAmounts>DepartmentCodeAmount"`
	LedgerAmounts []struct {
		LedgerCode  string  `xml:"LedgerCode"`
		Description string  `xml:"Description"`
		GLAccount   int     `xml:"GLAccount"`
		Amount      float64 `xml:"Amount"`
	} `xml:"LedgerAmounts>LedgerAmount"`
	Status     string `xml:"Status"`
	StatusNote string `xml:"StatusNote"`
}
