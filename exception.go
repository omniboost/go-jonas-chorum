package jonas_chorum

import "fmt"

// <?xml version="1.0" encoding="utf-8"?>
// <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
//   <soap:Body>
//     <pmsint_GetFinancialReportResponse xmlns="http://tempuri.org/RLXSOAP19/RLXSOAP19">
//       <pmsint_GetFinancialReportResult>
//         <ExceptionCode>70004</ExceptionCode>
//         <ExceptionDescription>Access to this web method is not permitted for your user account. Method requested: pmsint_GetFinancialReport, User: </ExceptionDescription>
//       </pmsint_GetFinancialReportResult>
//     </pmsint_GetFinancialReportResponse>
//   </soap:Body>
// </soap:Envelope>

type ExceptionBlock struct {
	ExceptionCode        int    `xml:"ExceptionCode"`
	ExceptionDescription string `xml:"ExceptionDescription"`
	ResponseCode         int    `xml:"ResponseCode"`
	ResponseDescription  string `xml:"ResponseDescription"`
}

func (eb ExceptionBlock) Error() string {
	return fmt.Sprintf("%d: %s", eb.ExceptionCode, eb.ExceptionDescription)
}
