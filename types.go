package jonas_chorum

import (
	"encoding/xml"
	"io"
)

type PeriodType int

const (
	PeriodTypeEndOfShit  PeriodType = 1
	PeriodTypeEndOfDay   PeriodType = 2
	PeriodTypeEndOfWeek  PeriodType = 3
	PeriodTypeEndOfMonth PeriodType = 4
	PeriodTypeEndOfYear  PeriodType = 5
)

type DocumentType string

const (
	DocumentTypeAll             DocumentType = "All"
	DocumentTypeInvoicesOnly    DocumentType = "InvoicesOnly"
	DocumentTypeCreditNotesOnly DocumentType = "CreditNotesOnly"
	DocumentTypeFolioBillsOnly  DocumentType = "FolioBillsOnly"
)

type VoucherType struct {
	ID      int    `json:"id"`
	Version int    `json:"version"`
	Name    string `json:"name"`
}

type Postings []Posting

type Posting struct {
	ID                  int         `json:"id,omitempty"`
	Version             int         `json:"version,omitempty"`
	URL                 string      `json:"url"`
	Date                string      `json:"date"`
	Description         string      `json:"description"`
	Account             *Account    `json:"account"`
	Customer            *Customer   `json:"customer,omitempty"`
	Supplier            *Supplier   `json:"supplier,omitempty"`
	Employee            *Employee   `json:"employee,omitempty"`
	Project             *Project    `json:"project,omitempty"`
	Product             *Product    `json:"project,omitempty"`
	Department          *Department `json:"department,omitempty"`
	VATType             *VATType    `json:"vatType,omitempty"`
	Amount              int         `json:"amount,omitempty"`
	AmountCurrency      int         `json:"amountCurrency,omitempty"`
	AmountGross         float64     `json:"amountGross"`
	AmountGrossCurrency float64     `json:"amountGrossCurrency"`
	Currency            *Currency   `json:"currency,omitempty"`
	CloseGroup          *CloseGroup `json:"closeGroup,omitempty"`
	InvoiceNumber       string      `json:"invoiceNumber,omitempty"`
	TermOfPayment       string      `json:"termOfPayment,omitempty"`
	Row                 int         `json:"row,omitempty"`
}

type Accounts []Account

type Account struct {
	ID                             int       `json:"id,omitempty"`
	Version                        int       `json:"version,omitempty"`
	URL                            string    `json:"url"`
	Number                         int       `json:"number"`
	Name                           string    `json:"name"`
	Description                    string    `json:"description"`
	VATType                        VATType   `json:"vatType"`
	VATLocked                      bool      `json:"vatLocked"`
	Currency                       *Currency `json:"currency"`
	IsCloseable                    bool      `json:"isCloseable"`
	IsApplicableForSupplierInvoice bool      `json:"isApplicableForSupplierInvoice"`
	RequireReconciliation          bool      `json:"requireReconciliation"`
	IsInactive                     bool      `json:"isInactive"`
	IsBankAccount                  bool      `json:"isBankAccount"`
	IsInvoiceAccount               bool      `json:"isInvoiceAccount"`
	BankAccountNumber              string    `json:"bankAccountNumber"`
	BankAccountCountry             struct {
		ID      int `json:"id"`
		Version int `json:"version"`
	} `json:"bankAccountCountry"`
	BankName         string `json:"bankName"`
	BankAccountIBAN  string `json:"bankAccountIBAN"`
	BankAccountSWIFT string `json:"bankAccountSWIFT"`
}

type Customers []Customer

type Customer struct {
	ID                    int             `json:"id,omitempty"`
	Version               int             `json:"version,omitempty"`
	URL                   string          `json:"url"`
	Name                  string          `json:"name"`
	OrganizationNumber    string          `json:"organizationNumber,omitempty"`
	SupplierNumber        int             `json:"supplierNumber,omitempty"`
	CustomerNumber        int             `json:"customerNumber,omitempty"`
	IsSupplier            bool            `json:"isSupplier,omitempty"`
	IsCustomer            bool            `json:"isCustomer,omitempty"`
	IsInactive            bool            `json:"isInactive,omitempty"`
	AccountManager        *AccountManager `json:"accountManager,omitempty"`
	Email                 string          `json:"email,omitempty"`
	InvoiceEmail          string          `json:"invoiceEmail,omitempty"`
	OverdueNoticeEmail    string          `json:"overdueNoticeEmail,omitempty"`
	BankAccounts          []string        `json:"bankAccounts,omitempty"`
	PhoneNumber           string          `json:"phoneNumber,omitempty"`
	PhoneNumberMobile     string          `json:"phoneNumberMobile,omitempty"`
	Description           string          `json:"description,omitempty"`
	IsPrivateIndividual   bool            `json:"isPrivateIndividual,omitempty"`
	SingleCustomerInvoice bool            `json:"singleCustomerInvoice,omitempty"`
	InvoiceSendMethod     string          `json:"invoiceSendMethod,omitempty"`
	EmailAttachmentType   string          `json:"emailAttachmentType,omitempty"`
	PostalAddress         *Address        `json:"postalAddress,omitempty"`
	PhysicalAddress       *Address        `json:"physicalAddress,omitempty"`
	DeliveryAddress       *Address        `json:"deliveryAddress,omitempty"`
	Category1             *Category       `json:"category1,omitempty"`
	Category2             *Category       `json:"category2,omitempty"`
	Category3             *Category       `json:"category3,omitempty"`
	InvoicesDueIn         int             `json:"invoicesDueIn,omitempty"`
	InvoicesDueInType     string          `json:"invoicesDueInType,omitempty"`
}

type Supplier struct {
	// 		ID                  int      `json:"id"`
	// 		Version             int      `json:"version"`
	// 		Name                string   `json:"name"`
	// 		OrganizationNumber  string   `json:"organizationNumber"`
	// 		SupplierNumber      int      `json:"supplierNumber"`
	// 		CustomerNumber      int      `json:"customerNumber"`
	// 		IsCustomer          bool     `json:"isCustomer"`
	// 		Email               string   `json:"email"`
	// 		BankAccounts        []string `json:"bankAccounts"`
	// 		InvoiceEmail        string   `json:"invoiceEmail"`
	// 		OverdueNoticeEmail  string   `json:"overdueNoticeEmail"`
	// 		PhoneNumber         string   `json:"phoneNumber"`
	// 		PhoneNumberMobile   string   `json:"phoneNumberMobile"`
	// 		Description         string   `json:"description"`
	// 		IsPrivateIndividual bool     `json:"isPrivateIndividual"`
	// 		ShowProducts        bool     `json:"showProducts"`
	// 		AccountManager      struct {
	// 			ID                       int    `json:"id"`
	// 			Version                  int    `json:"version"`
	// 			FirstName                string `json:"firstName"`
	// 			LastName                 string `json:"lastName"`
	// 			EmployeeNumber           string `json:"employeeNumber"`
	// 			DateOfBirth              string `json:"dateOfBirth"`
	// 			Email                    string `json:"email"`
	// 			PhoneNumberMobileCountry struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"phoneNumberMobileCountry"`
	// 			PhoneNumberMobile      string `json:"phoneNumberMobile"`
	// 			PhoneNumberHome        string `json:"phoneNumberHome"`
	// 			PhoneNumberWork        string `json:"phoneNumberWork"`
	// 			NationalIdentityNumber string `json:"nationalIdentityNumber"`
	// 			Dnumber                string `json:"dnumber"`
	// 			InternationalID        struct {
	// 				IntAmeldingType string `json:"intAmeldingType"`
	// 				Country         struct {
	// 					ID      int `json:"id"`
	// 					Version int `json:"version"`
	// 				} `json:"country"`
	// 				Number string `json:"number"`
	// 			} `json:"internationalId"`
	// 			BankAccountNumber     string `json:"bankAccountNumber"`
	// 			Iban                  string `json:"iban"`
	// 			Bic                   string `json:"bic"`
	// 			CreditorBankCountryID int    `json:"creditorBankCountryId"`
	// 			UsesAbroadPayment     bool   `json:"usesAbroadPayment"`
	// 			UserType              string `json:"userType"`
	// 			Comments              string `json:"comments"`
	// 			Address               struct {
	// 				ID           int    `json:"id"`
	// 				Version      int    `json:"version"`
	// 				AddressLine1 string `json:"addressLine1"`
	// 				AddressLine2 string `json:"addressLine2"`
	// 				PostalCode   string `json:"postalCode"`
	// 				City         string `json:"city"`
	// 				Country      struct {
	// 					ID      int `json:"id"`
	// 					Version int `json:"version"`
	// 				} `json:"country"`
	// 			} `json:"address"`
	// 			Department struct {
	// 				ID               int    `json:"id"`
	// 				Version          int    `json:"version"`
	// 				Name             string `json:"name"`
	// 				DepartmentNumber string `json:"departmentNumber"`
	// 			} `json:"department"`
	// 			Employments []struct {
	// 				ID           int    `json:"id"`
	// 				Version      int    `json:"version"`
	// 				EmploymentID string `json:"employmentId"`
	// 				StartDate    string `json:"startDate"`
	// 				EndDate      string `json:"endDate"`
	// 				Division     struct {
	// 					ID                 int    `json:"id"`
	// 					Version            int    `json:"version"`
	// 					Name               string `json:"name"`
	// 					StartDate          string `json:"startDate"`
	// 					EndDate            string `json:"endDate"`
	// 					OrganizationNumber string `json:"organizationNumber"`
	// 					Municipality       struct {
	// 						ID      int `json:"id"`
	// 						Version int `json:"version"`
	// 					} `json:"municipality"`
	// 				} `json:"division"`
	// 				LastSalaryChangeDate     string `json:"lastSalaryChangeDate"`
	// 				NoEmploymentRelationship bool   `json:"noEmploymentRelationship"`
	// 				IsMainEmployer           bool   `json:"isMainEmployer"`
	// 				TaxDeductionCode         string `json:"taxDeductionCode"`
	// 				EmploymentDetails        []struct {
	// 					ID                 int    `json:"id"`
	// 					Version            int    `json:"version"`
	// 					Date               string `json:"date"`
	// 					EmploymentType     string `json:"employmentType"`
	// 					MaritimeEmployment struct {
	// 						ShipRegister string `json:"shipRegister"`
	// 						ShipType     string `json:"shipType"`
	// 						TradeArea    string `json:"tradeArea"`
	// 					} `json:"maritimeEmployment"`
	// 					RemunerationType   string `json:"remunerationType"`
	// 					WorkingHoursScheme string `json:"workingHoursScheme"`
	// 					ShiftDurationHours int    `json:"shiftDurationHours"`
	// 					OccupationCode     struct {
	// 						ID      int    `json:"id"`
	// 						Version int    `json:"version"`
	// 						NameNO  string `json:"nameNO"`
	// 						Code    string `json:"code"`
	// 					} `json:"occupationCode"`
	// 					PercentageOfFullTimeEquivalent int `json:"percentageOfFullTimeEquivalent"`
	// 					AnnualSalary                   int `json:"annualSalary"`
	// 					HourlyWage                     int `json:"hourlyWage"`
	// 					PayrollTaxMunicipalityID       struct {
	// 						ID      int `json:"id"`
	// 						Version int `json:"version"`
	// 					} `json:"payrollTaxMunicipalityId"`
	// 				} `json:"employmentDetails"`
	// 			} `json:"employments"`
	// 			HolidayAllowanceEarned struct {
	// 				Year                   int `json:"year"`
	// 				Amount                 int `json:"amount"`
	// 				Basis                  int `json:"basis"`
	// 				AmountExtraHolidayWeek int `json:"amountExtraHolidayWeek"`
	// 			} `json:"holidayAllowanceEarned"`
	// 		} `json:"accountManager"`
	// 		PostalAddress struct {
	// 			ID           int    `json:"id"`
	// 			Version      int    `json:"version"`
	// 			AddressLine1 string `json:"addressLine1"`
	// 			AddressLine2 string `json:"addressLine2"`
	// 			PostalCode   string `json:"postalCode"`
	// 			City         string `json:"city"`
	// 			Country      struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"country"`
	// 		} `json:"postalAddress"`
	// 		PhysicalAddress struct {
	// 			ID           int    `json:"id"`
	// 			Version      int    `json:"version"`
	// 			AddressLine1 string `json:"addressLine1"`
	// 			AddressLine2 string `json:"addressLine2"`
	// 			PostalCode   string `json:"postalCode"`
	// 			City         string `json:"city"`
	// 			Country      struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"country"`
	// 		} `json:"physicalAddress"`
	// 		DeliveryAddress struct {
	// 			ID       int `json:"id"`
	// 			Version  int `json:"version"`
	// 			Employee struct {
	// 				ID                       int    `json:"id"`
	// 				Version                  int    `json:"version"`
	// 				FirstName                string `json:"firstName"`
	// 				LastName                 string `json:"lastName"`
	// 				EmployeeNumber           string `json:"employeeNumber"`
	// 				DateOfBirth              string `json:"dateOfBirth"`
	// 				Email                    string `json:"email"`
	// 				PhoneNumberMobileCountry struct {
	// 					ID      int `json:"id"`
	// 					Version int `json:"version"`
	// 				} `json:"phoneNumberMobileCountry"`
	// 				PhoneNumberMobile      string `json:"phoneNumberMobile"`
	// 				PhoneNumberHome        string `json:"phoneNumberHome"`
	// 				PhoneNumberWork        string `json:"phoneNumberWork"`
	// 				NationalIdentityNumber string `json:"nationalIdentityNumber"`
	// 				Dnumber                string `json:"dnumber"`
	// 				InternationalID        struct {
	// 					IntAmeldingType string `json:"intAmeldingType"`
	// 					Country         struct {
	// 						ID      int `json:"id"`
	// 						Version int `json:"version"`
	// 					} `json:"country"`
	// 					Number string `json:"number"`
	// 				} `json:"internationalId"`
	// 				BankAccountNumber     string `json:"bankAccountNumber"`
	// 				Iban                  string `json:"iban"`
	// 				Bic                   string `json:"bic"`
	// 				CreditorBankCountryID int    `json:"creditorBankCountryId"`
	// 				UsesAbroadPayment     bool   `json:"usesAbroadPayment"`
	// 				UserType              string `json:"userType"`
	// 				Comments              string `json:"comments"`
	// 				Address               struct {
	// 					ID           int    `json:"id"`
	// 					Version      int    `json:"version"`
	// 					AddressLine1 string `json:"addressLine1"`
	// 					AddressLine2 string `json:"addressLine2"`
	// 					PostalCode   string `json:"postalCode"`
	// 					City         string `json:"city"`
	// 					Country      struct {
	// 						ID      int `json:"id"`
	// 						Version int `json:"version"`
	// 					} `json:"country"`
	// 				} `json:"address"`
	// 				Department struct {
	// 					ID               int    `json:"id"`
	// 					Version          int    `json:"version"`
	// 					Name             string `json:"name"`
	// 					DepartmentNumber string `json:"departmentNumber"`
	// 				} `json:"department"`
	// 				Employments []struct {
	// 					ID           int    `json:"id"`
	// 					Version      int    `json:"version"`
	// 					EmploymentID string `json:"employmentId"`
	// 					StartDate    string `json:"startDate"`
	// 					EndDate      string `json:"endDate"`
	// 					Division     struct {
	// 						ID                 int    `json:"id"`
	// 						Version            int    `json:"version"`
	// 						Name               string `json:"name"`
	// 						StartDate          string `json:"startDate"`
	// 						EndDate            string `json:"endDate"`
	// 						OrganizationNumber string `json:"organizationNumber"`
	// 						Municipality       struct {
	// 							ID      int `json:"id"`
	// 							Version int `json:"version"`
	// 						} `json:"municipality"`
	// 					} `json:"division"`
	// 					LastSalaryChangeDate     string `json:"lastSalaryChangeDate"`
	// 					NoEmploymentRelationship bool   `json:"noEmploymentRelationship"`
	// 					IsMainEmployer           bool   `json:"isMainEmployer"`
	// 					TaxDeductionCode         string `json:"taxDeductionCode"`
	// 					EmploymentDetails        []struct {
	// 						ID                 int    `json:"id"`
	// 						Version            int    `json:"version"`
	// 						Date               string `json:"date"`
	// 						EmploymentType     string `json:"employmentType"`
	// 						MaritimeEmployment struct {
	// 							ShipRegister string `json:"shipRegister"`
	// 							ShipType     string `json:"shipType"`
	// 							TradeArea    string `json:"tradeArea"`
	// 						} `json:"maritimeEmployment"`
	// 						RemunerationType   string `json:"remunerationType"`
	// 						WorkingHoursScheme string `json:"workingHoursScheme"`
	// 						ShiftDurationHours int    `json:"shiftDurationHours"`
	// 						OccupationCode     struct {
	// 							ID      int    `json:"id"`
	// 							Version int    `json:"version"`
	// 							NameNO  string `json:"nameNO"`
	// 							Code    string `json:"code"`
	// 						} `json:"occupationCode"`
	// 						PercentageOfFullTimeEquivalent int `json:"percentageOfFullTimeEquivalent"`
	// 						AnnualSalary                   int `json:"annualSalary"`
	// 						HourlyWage                     int `json:"hourlyWage"`
	// 						PayrollTaxMunicipalityID       struct {
	// 							ID      int `json:"id"`
	// 							Version int `json:"version"`
	// 						} `json:"payrollTaxMunicipalityId"`
	// 					} `json:"employmentDetails"`
	// 				} `json:"employments"`
	// 				HolidayAllowanceEarned struct {
	// 					Year                   int `json:"year"`
	// 					Amount                 int `json:"amount"`
	// 					Basis                  int `json:"basis"`
	// 					AmountExtraHolidayWeek int `json:"amountExtraHolidayWeek"`
	// 				} `json:"holidayAllowanceEarned"`
	// 			} `json:"employee"`
	// 			AddressLine1 string `json:"addressLine1"`
	// 			AddressLine2 string `json:"addressLine2"`
	// 			PostalCode   string `json:"postalCode"`
	// 			City         string `json:"city"`
	// 			Country      struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"country"`
	// 			Name string `json:"name"`
	// 		} `json:"deliveryAddress"`
	// 		Category1 struct {
	// 			ID          int    `json:"id"`
	// 			Version     int    `json:"version"`
	// 			Name        string `json:"name"`
	// 			Number      string `json:"number"`
	// 			Description string `json:"description"`
	// 			Type        int    `json:"type"`
	// 		} `json:"category1"`
	// 		Category2 struct {
	// 			ID          int    `json:"id"`
	// 			Version     int    `json:"version"`
	// 			Name        string `json:"name"`
	// 			Number      string `json:"number"`
	// 			Description string `json:"description"`
	// 			Type        int    `json:"type"`
	// 		} `json:"category2"`
	// 		Category3 struct {
	// 			ID          int    `json:"id"`
	// 			Version     int    `json:"version"`
	// 			Name        string `json:"name"`
	// 			Number      string `json:"number"`
	// 			Description string `json:"description"`
	// 			Type        int    `json:"type"`
	// 		} `json:"category3"`
}

type Employee struct {
	// 	Employee struct {
	// 		ID                       int    `json:"id"`
	// 		Version                  int    `json:"version"`
	// 		FirstName                string `json:"firstName"`
	// 		LastName                 string `json:"lastName"`
	// 		EmployeeNumber           string `json:"employeeNumber"`
	// 		DateOfBirth              string `json:"dateOfBirth"`
	// 		Email                    string `json:"email"`
	// 		PhoneNumberMobileCountry struct {
	// 			ID      int `json:"id"`
	// 			Version int `json:"version"`
	// 		} `json:"phoneNumberMobileCountry"`
	// 		PhoneNumberMobile      string `json:"phoneNumberMobile"`
	// 		PhoneNumberHome        string `json:"phoneNumberHome"`
	// 		PhoneNumberWork        string `json:"phoneNumberWork"`
	// 		NationalIdentityNumber string `json:"nationalIdentityNumber"`
	// 		Dnumber                string `json:"dnumber"`
	// 		InternationalID        struct {
	// 			IntAmeldingType string `json:"intAmeldingType"`
	// 			Country         struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"country"`
	// 			Number string `json:"number"`
	// 		} `json:"internationalId"`
	// 		BankAccountNumber     string `json:"bankAccountNumber"`
	// 		Iban                  string `json:"iban"`
	// 		Bic                   string `json:"bic"`
	// 		CreditorBankCountryID int    `json:"creditorBankCountryId"`
	// 		UsesAbroadPayment     bool   `json:"usesAbroadPayment"`
	// 		UserType              string `json:"userType"`
	// 		Comments              string `json:"comments"`
	// 		Address               struct {
	// 			ID           int    `json:"id"`
	// 			Version      int    `json:"version"`
	// 			AddressLine1 string `json:"addressLine1"`
	// 			AddressLine2 string `json:"addressLine2"`
	// 			PostalCode   string `json:"postalCode"`
	// 			City         string `json:"city"`
	// 			Country      struct {
	// 				ID      int `json:"id"`
	// 				Version int `json:"version"`
	// 			} `json:"country"`
	// 		} `json:"address"`
	// 		Department struct {
	// 			ID               int    `json:"id"`
	// 			Version          int    `json:"version"`
	// 			Name             string `json:"name"`
	// 			DepartmentNumber string `json:"departmentNumber"`
	// 		} `json:"department"`
	// 		Employments []struct {
	// 			ID           int    `json:"id"`
	// 			Version      int    `json:"version"`
	// 			EmploymentID string `json:"employmentId"`
	// 			StartDate    string `json:"startDate"`
	// 			EndDate      string `json:"endDate"`
	// 			Division     struct {
	// 				ID                 int    `json:"id"`
	// 				Version            int    `json:"version"`
	// 				Name               string `json:"name"`
	// 				StartDate          string `json:"startDate"`
	// 				EndDate            string `json:"endDate"`
	// 				OrganizationNumber string `json:"organizationNumber"`
	// 				Municipality       struct {
	// 					ID      int `json:"id"`
	// 					Version int `json:"version"`
	// 				} `json:"municipality"`
	// 			} `json:"division"`
	// 			LastSalaryChangeDate     string `json:"lastSalaryChangeDate"`
	// 			NoEmploymentRelationship bool   `json:"noEmploymentRelationship"`
	// 			IsMainEmployer           bool   `json:"isMainEmployer"`
	// 			TaxDeductionCode         string `json:"taxDeductionCode"`
	// 			EmploymentDetails        []struct {
	// 				ID                 int    `json:"id"`
	// 				Version            int    `json:"version"`
	// 				Date               string `json:"date"`
	// 				EmploymentType     string `json:"employmentType"`
	// 				MaritimeEmployment struct {
	// 					ShipRegister string `json:"shipRegister"`
	// 					ShipType     string `json:"shipType"`
	// 					TradeArea    string `json:"tradeArea"`
	// 				} `json:"maritimeEmployment"`
	// 				RemunerationType   string `json:"remunerationType"`
	// 				WorkingHoursScheme string `json:"workingHoursScheme"`
	// 				ShiftDurationHours int    `json:"shiftDurationHours"`
	// 				OccupationCode     struct {
	// 					ID      int    `json:"id"`
	// 					Version int    `json:"version"`
	// 					NameNO  string `json:"nameNO"`
	// 					Code    string `json:"code"`
	// 				} `json:"occupationCode"`
	// 				PercentageOfFullTimeEquivalent int `json:"percentageOfFullTimeEquivalent"`
	// 				AnnualSalary                   int `json:"annualSalary"`
	// 				HourlyWage                     int `json:"hourlyWage"`
	// 				PayrollTaxMunicipalityID       struct {
	// 					ID      int `json:"id"`
	// 					Version int `json:"version"`
	// 				} `json:"payrollTaxMunicipalityId"`
	// 			} `json:"employmentDetails"`
	// 		} `json:"employments"`
	// 		HolidayAllowanceEarned struct {
	// 			Year                   int `json:"year"`
	// 			Amount                 int `json:"amount"`
	// 			Basis                  int `json:"basis"`
	// 			AmountExtraHolidayWeek int `json:"amountExtraHolidayWeek"`
	// 		} `json:"holidayAllowanceEarned"`
	// 	} `json:"employee"`
}

type Project struct {
}

type Product struct {
}

type Department struct {
	ID               int    `json:"id"`
	Version          int    `json:"version"`
	URL              string `json:"url"`
	Name             string `json:"name"`
	DepartmentNumber string `json:"departmentNumber"`
}

type VATTypes []VATType

type VATType struct {
	ID         int     `json:"id"`
	Version    int     `json:"version"`
	URL        string  `json:"url"`
	Name       string  `json:"name"`
	Number     string  `json:"number"`
	Percentage float64 `json:"percentage"`
}

type Currency struct {
	// ID          int    `json:"id,omitempty"`
	// Version     int    `json:"version,omitempty"`
	// Code        string `json:"code"`
	// Description string `json:"description,omitempty"`
	// Factor      int    `json:"factor,omitempty"`
}

type CloseGroup struct {
	// 		ID      int    `json:"id"`
	// 		Version int    `json:"version"`
	// 		Date    string `json:"date"`
}

type Document struct {
	// 	ID       int    `json:"id"`
	// 	Version  int    `json:"version"`
	// 	FileName string `json:"fileName"`
}

type Attachment struct {
	// 	ID       int    `json:"id"`
	// 	Version  int    `json:"version"`
	// 	FileName string `json:"fileName"`
}

type EDIDocument struct {
	// 	ID       int    `json:"id"`
	// 	Version  int    `json:"version"`
	// 	FileName string `json:"fileName"`
}

type Invoice struct {
	ID             int      `json:"id"`
	Version        int      `json:"version"`
	URL            string   `json:"url"`
	InvoiceNumber  int      `json:"invoiceNumber"`
	InvoiceDate    string   `json:"invoiceDate"`
	Customer       Customer `json:"customer"`
	InvoiceDueDate string   `json:"invoiceDueDate"`
	KID            string   `json:"kid"`
	Comment        string   `json:"comment"`
	Orders         Orders   `json:"orders"`
	Voucher        Voucher  `json:"voucher"`
	// Currency       Currency `json:"currency"`
	InvoiceRemarks string `json:"invoiceRemarks"`
	PaymentTypeID  int    `json:"paymentTypeId"`
	PaidAmount     int    `json:"paidAmount"`
	EhfSendStatus  string `json:"ehfSendStatus,omitempty"`
}

type Voucher struct {
	ID          int          `json:"id,omitempty"`
	Version     int          `json:"version,omitempty"`
	Date        string       `json:"date"`
	Description string       `json:"description"`
	VoucherType *VoucherType `json:"voucherType,omitempty"`
	Postings    Postings     `json:"postings"`
	Document    *Document    `json:"document,omitempty"`
	Attachment  *Attachment  `json:"attachment,omitempty"`
	EDIDocument *EDIDocument `json:"ediDocument,omitempty"`
}

type Orders []Order

type Order struct {
	// ID                 int      `json:"id"`
	// Version            int      `json:"version"`
	// URL                string   `json:"url"`
	Customer Customer `json:"customer"`
	// Contact            Contact  `json:"contact"`
	// Attn               Attn     `json:"attn"`
	ReceiverEmail      string `json:"receiverEmail"`
	OverdueNoticeEmail string `json:"overdueNoticeEmail"`
	Number             string `json:"number"`
	Reference          string `json:"reference"`
	// OurContact         Contact `json:"ourContact"`
	// OurContactEmployee Contact `json:"ourContactEmployee"`
	// Department         struct {
	// 	ID               int    `json:"id"`
	// 	Version          int    `json:"version"`
	// 	Name             string `json:"name"`
	// 	DepartmentNumber string `json:"departmentNumber"`
	// } `json:"department"`
	OrderDate string `json:"orderDate"`
	// Project                                     Project    `json:"project"`
	InvoiceComment string `json:"invoiceComment"`
	// Currency       Currency `json:"currency"`
	// InvoicesDueIn                               int        `json:"invoicesDueIn"`
	// InvoicesDueInType                           string     `json:"invoicesDueInType"`
	IsShowOpenPostsOnInvoices bool `json:"isShowOpenPostsOnInvoices"`
	// IsClosed                                    bool       `json:"isClosed"`
	DeliveryDate string `json:"deliveryDate"`
	// DeliveryAddress                             Address    `json:"deliveryAddress"`
	// DeliveryComment                             string     `json:"deliveryComment"`
	// IsPrioritizeAmountsIncludingVat             bool       `json:"isPrioritizeAmountsIncludingVat"`
	// OrderLineSorting                            string     `json:"orderLineSorting"`
	OrderLines OrderLines `json:"orderLines"`
	// IsSubscription                              bool       `json:"isSubscription"`
	// SubscriptionDuration                        int        `json:"subscriptionDuration"`
	// SubscriptionDurationType                    string     `json:"subscriptionDurationType"`
	// SubscriptionPeriodsOnInvoice                int        `json:"subscriptionPeriodsOnInvoice"`
	// SubscriptionInvoicingTimeInAdvanceOrArrears string     `json:"subscriptionInvoicingTimeInAdvanceOrArrears"`
	// SubscriptionInvoicingTime                   int        `json:"subscriptionInvoicingTime"`
	// SubscriptionInvoicingTimeType               string     `json:"subscriptionInvoicingTimeType"`
	// IsSubscriptionAutoInvoicing                 bool       `json:"isSubscriptionAutoInvoicing"`
}

type Contact struct {
	ID                       int    `json:"id"`
	Version                  int    `json:"version"`
	URL                      string `json:"url"`
	FirstName                string `json:"firstName"`
	LastName                 string `json:"lastName"`
	Email                    string `json:"email"`
	PhoneNumberMobileCountry struct {
		ID      int `json:"id"`
		Version int `json:"version"`
	} `json:"phoneNumberMobileCountry"`
	PhoneNumberMobile string   `json:"phoneNumberMobile"`
	PhoneNumberWork   string   `json:"phoneNumberWork"`
	Customer          Customer `json:"customer"`
}

type Attn struct {
	ID                       int    `json:"id"`
	Version                  int    `json:"version"`
	URL                      string `json:"url"`
	FirstName                string `json:"firstName"`
	LastName                 string `json:"lastName"`
	Email                    string `json:"email"`
	PhoneNumberMobileCountry struct {
		ID      int `json:"id"`
		Version int `json:"version"`
	} `json:"phoneNumberMobileCountry"`
	PhoneNumberMobile string   `json:"phoneNumberMobile"`
	PhoneNumberWork   string   `json:"phoneNumberWork"`
	Customer          Customer `json:"customer"`
}

type Address struct {
	ID           int      `json:"id"`
	Version      int      `json:"version"`
	URL          string   `json:"url"`
	Employee     Employee `json:"employee"`
	AddressLine1 string   `json:"addressLine1"`
	AddressLine2 string   `json:"addressLine2"`
	PostalCode   string   `json:"postalCode"`
	City         string   `json:"city"`
	Country      struct {
		ID      int `json:"id"`
		Version int `json:"version"`
	} `json:"country"`
	Name string `json:"name"`
}

type OrderLines []OrderLine

type OrderLine struct {
	ID int `json:"id"`
	// Version int     `json:"version"`
	// URL     string  `json:"url"`
	Product Product `json:"product"`
	// Inventory struct {
	// 	ID              int    `json:"id"`
	// 	Version         int    `json:"version"`
	// 	Name            string `json:"name"`
	// 	Number          string `json:"number"`
	// 	IsMainInventory bool   `json:"isMainInventory"`
	// 	IsInactive      bool   `json:"isInactive"`
	// } `json:"inventory"`
	Description                   string  `json:"description"`
	Count                         int     `json:"count"`
	UnitCostCurrency              int     `json:"unitCostCurrency"`
	UnitPriceExcludingVATCurrency float64 `json:"unitPriceExcludingVatCurrency"`
	// Currency                      Currency `json:"currency"`
	// Markup                        int     `json:"markup"`
	// Discount                      int     `json:"discount"`
	VATType                       VATType `json:"vatType"`
	UnitPriceIncludingVATCurrency float64 `json:"unitPriceIncludingVatCurrency"`
	AmountExcludingVATCurrency    float64 `json:"amountExcludingVatCurrency"`
	AmountIncludingVATCurrency    float64 `json:"amountIncludingVatCurrency"`
	// IsSubscription                bool    `json:"isSubscription"`
	// SubscriptionPeriodStart       string  `json:"subscriptionPeriodStart"`
	// SubscriptionPeriodEnd         string  `json:"subscriptionPeriodEnd"`
	// OrderGroup                    struct {
	// 	ID        int    `json:"id"`
	// 	Version   int    `json:"version"`
	// 	Title     string `json:"title"`
	// 	Comment   string `json:"comment"`
	// 	SortIndex int    `json:"sortIndex"`
	// } `json:"orderGroup"`
}

type AccountManager struct {
	ID                       int    `json:"id"`
	Version                  int    `json:"version"`
	URL                      string `json:"url"`
	FirstName                string `json:"firstName"`
	LastName                 string `json:"lastName"`
	EmployeeNumber           string `json:"employeeNumber"`
	DateOfBirth              string `json:"dateOfBirth"`
	Email                    string `json:"email"`
	PhoneNumberMobileCountry struct {
		ID      int `json:"id"`
		Version int `json:"version"`
	} `json:"phoneNumberMobileCountry"`
	PhoneNumberMobile      string `json:"phoneNumberMobile"`
	PhoneNumberHome        string `json:"phoneNumberHome"`
	PhoneNumberWork        string `json:"phoneNumberWork"`
	NationalIdentityNumber string `json:"nationalIdentityNumber"`
	Dnumber                string `json:"dnumber"`
	InternationalID        struct {
		IntAmeldingType string `json:"intAmeldingType"`
		Country         struct {
			ID      int `json:"id"`
			Version int `json:"version"`
		} `json:"country"`
		Number string `json:"number"`
	} `json:"internationalId"`
	BankAccountNumber     string     `json:"bankAccountNumber"`
	Iban                  string     `json:"iban"`
	Bic                   string     `json:"bic"`
	CreditorBankCountryID int        `json:"creditorBankCountryId"`
	UsesAbroadPayment     bool       `json:"usesAbroadPayment"`
	UserType              string     `json:"userType"`
	Comments              string     `json:"comments"`
	Address               Address    `json:"address"`
	Department            Department `json:"department"`
	Employments           []struct {
		ID           int    `json:"id"`
		Version      int    `json:"version"`
		EmploymentID string `json:"employmentId"`
		StartDate    string `json:"startDate"`
		EndDate      string `json:"endDate"`
		Division     struct {
			ID                 int    `json:"id"`
			Version            int    `json:"version"`
			Name               string `json:"name"`
			StartDate          string `json:"startDate"`
			EndDate            string `json:"endDate"`
			OrganizationNumber string `json:"organizationNumber"`
			Municipality       struct {
				ID      int `json:"id"`
				Version int `json:"version"`
			} `json:"municipality"`
		} `json:"division"`
		LastSalaryChangeDate     string `json:"lastSalaryChangeDate"`
		NoEmploymentRelationship bool   `json:"noEmploymentRelationship"`
		IsMainEmployer           bool   `json:"isMainEmployer"`
		TaxDeductionCode         string `json:"taxDeductionCode"`
		EmploymentDetails        []struct {
			ID                 int    `json:"id"`
			Version            int    `json:"version"`
			Date               string `json:"date"`
			EmploymentType     string `json:"employmentType"`
			MaritimeEmployment struct {
				ShipRegister string `json:"shipRegister"`
				ShipType     string `json:"shipType"`
				TradeArea    string `json:"tradeArea"`
			} `json:"maritimeEmployment"`
			RemunerationType   string `json:"remunerationType"`
			WorkingHoursScheme string `json:"workingHoursScheme"`
			ShiftDurationHours int    `json:"shiftDurationHours"`
			OccupationCode     struct {
				ID      int    `json:"id"`
				Version int    `json:"version"`
				NameNO  string `json:"nameNO"`
				Code    string `json:"code"`
			} `json:"occupationCode"`
			PercentageOfFullTimeEquivalent int `json:"percentageOfFullTimeEquivalent"`
			AnnualSalary                   int `json:"annualSalary"`
			HourlyWage                     int `json:"hourlyWage"`
			PayrollTaxMunicipalityID       struct {
				ID      int `json:"id"`
				Version int `json:"version"`
			} `json:"payrollTaxMunicipalityId"`
		} `json:"employmentDetails"`
	} `json:"employments"`
	HolidayAllowanceEarned struct {
		Year                   int `json:"year"`
		Amount                 int `json:"amount"`
		Basis                  int `json:"basis"`
		AmountExtraHolidayWeek int `json:"amountExtraHolidayWeek"`
	} `json:"holidayAllowanceEarned"`
}

type Category struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Number      string `json:"number"`
	Description string `json:"description"`
	Type        int    `json:"type"`
}

type Periods []Period

type Period struct {
	StrKey               string   `xml:"strKey"`
	IPeriodTypeID        int      `xml:"iPeriodTypeID"`
	StrPeriodDescription string   `xml:"strPeriodDescription"`
	IPeriodType          string   `xml:"iPeriodType"`
	IPeriodID            int      `xml:"iPeriodID"`
	DTOpenPeriod         DateTime `xml:"dtOpenPeriod"`
	DTClosePeriod        DateTime `xml:"dtClosePeriod"`
	IStatus              string   `xml:"iStatus"`
	StrOperatorCode      string   `xml:"strOperatorCode"`
}

type FinancialReportData struct {
	Document DOCUMENT `xml:"DOCUMENT"`
}

type DOCUMENT struct {
	XMLName           xml.Name `xml:"DOCUMENT"`
	ANALYSISCODESALES struct {
		ANALYSISCODESALESTOTALS struct {
			ANALYSISCODESALE []struct {
				CODE          string  `xml:"CODE,attr"`
				DESCRIPTION   string  `xml:"DESCRIPTION"`
				NETTTOTAL     float64 `xml:"NETTTOTAL"`
				TAXTOTAL      float64 `xml:"TAXTOTAL"`
				GROSSTOTAL    float64 `xml:"GROSSTOTAL"`
				MARKETSEGMENT struct {
					CODE       []string  `xml:"CODE"`
					NETTTOTAL  []float64 `xml:"NETTTOTAL"`
					TAXTOTAL   []float64 `xml:"TAXTOTAL"`
					GROSSTOTAL []float64 `xml:"GROSSTOTAL"`
				} `xml:"MARKETSEGMENT"`
				MEDIASOURCE struct {
					CODE       []string  `xml:"CODE"`
					NETTTOTAL  []float64 `xml:"NETTTOTAL"`
					TAXTOTAL   []float64 `xml:"TAXTOTAL"`
					GROSSTOTAL []float64 `xml:"GROSSTOTAL"`
				} `xml:"MEDIASOURCE"`
				MEDIASOURCEBYMARKETSEGMENT struct {
					CODE          []string  `xml:"CODE"`
					NETTTOTAL     []string  `xml:"NETTTOTAL"`
					TAXTOTAL      []float64 `xml:"TAXTOTAL"`
					GROSSTOTAL    []float64 `xml:"GROSSTOTAL"`
					MARKETSEGMENT []struct {
						CODE       []string  `xml:"CODE"`
						NETTTOTAL  []float64 `xml:"NETTTOTAL"`
						TAXTOTAL   []float64 `xml:"TAXTOTAL"`
						GROSSTOTAL []float64 `xml:"GROSSTOTAL"`
					} `xml:"MARKETSEGMENT"`
				} `xml:"MEDIASOURCEBYMARKETSEGMENT"`
				ANALYSISCODETAXSUMMARY struct {
					ANALYSISCODETAXSUMMARYITEM struct {
						CODE       string  `xml:"CODE,attr"`
						NETTTOTAL  float64 `xml:"NETTTOTAL"`
						TAXTOTAL   float64 `xml:"TAXTOTAL"`
						GROSSTOTAL float64 `xml:"GROSSTOTAL"`
					} `xml:"ANALYSISCODETAXSUMMARYITEM"`
				} `xml:"ANALYSISCODETAXSUMMARY"`
			} `xml:"ANALYSISCODESALE"`
		} `xml:"ANALYSISCODESALESTOTALS"`
		ANALYSISCODESALESTRANS struct {
			ANALYSISCODESALE []struct {
				CODE         string  `xml:"CODE,attr"`
				TRANSID      string  `xml:"TRANSID"`
				DESCRIPTION  string  `xml:"DESCRIPTION"`
				QUANTITY     float64 `xml:"QUANTITY"`
				GROSSPERUNIT float64 `xml:"GROSSPERUNIT"`
				GROSSTOTAL   float64 `xml:"GROSSTOTAL"`
				NETTPERUNIT  float64 `xml:"NETTPERUNIT"`
				NETTTOTAL    float64 `xml:"NETTTOTAL"`
				TAXTOTAL     float64 `xml:"TAXTOTAL"`
				SOURCE       string  `xml:"SOURCE"`
				ENTRYTYPE    string  `xml:"ENTRYTYPE"`
				OPERATOR     string  `xml:"OPERATOR"`
				TAXRULE      string  `xml:"TAXRULE"`
				COMPANYREF   string  `xml:"COMPANYREF"`
			} `xml:"ANALYSISCODESALE"`
		} `xml:"ANALYSISCODESALESTRANS"`
	} `xml:"ANALYSISCODESALES"`
	DEPOSITANALYSIS struct {
		DEPOSITANALYSISTOTALS XMLMapStringStruct[struct {
			DESCRIPTION string  `xml:"DESCRIPTION"`
			NETTTOTAL   float64 `xml:"NETTTOTAL"`
			TAXTOTAL    float64 `xml:"TAXTOTAL"`
			GROSSTOTAL  float64 `xml:"GROSSTOTAL"`
			TAXRULE     string  `xml:"TAXRULE"`
		}] `xml:"DEPOSITANALYSISTOTALS"`
		DEPOSITANALYSISTRANS XMLMapStringStruct[[]struct {
			TRANSID      string  `xml:"TRANSID"`
			DESCRIPTION  string  `xml:"DESCRIPTION"`
			QUANTITY     float64 `xml:"QUANTITY"`
			GROSSPERUNIT float64 `xml:"GROSSPERUNIT"`
			GROSSTOTAL   float64 `xml:"GROSSTOTAL"`
			NETTPERUNIT  float64 `xml:"NETTPERUNIT"`
			NETTTOTAL    float64 `xml:"NETTTOTAL"`
			TAXTOTAL     float64 `xml:"TAXTOTAL"`
			TAXRULE      string  `xml:"TAXRULE"`
			SOURCE       string  `xml:"SOURCE"`
		}] `xml:"DEPOSITANALYSISTRANS"`
	} `xml:"DEPOSITANALYSIS"`
	LEDGERANALYSIS struct {
		LEDGERANALYSISTOTALS XMLMapStringStruct[struct {
			DESCRIPTION string  `xml:"DESCRIPTION"`
			NETTTOTAL   float64 `xml:"NETTTOTAL"`
			TAXTOTAL    float64 `xml:"TAXTOTAL"`
			GROSSTOTAL  float64 `xml:"GROSSTOTAL"`
		}] `xml:"LEDGERANALYSISTOTALS"`
		LEDGERANALYSISTRANS XMLMapStringStruct[[]struct {
			TRANSID      string  `xml:"TRANSID"`
			DESCRIPTION  string  `xml:"DESCRIPTION"`
			QUANTITY     float64 `xml:"QUANTITY"`
			GROSSPERUNIT float64 `xml:"GROSSPERUNIT"`
			GROSSTOTAL   float64 `xml:"GROSSTOTAL"`
			NETTPERUNIT  float64 `xml:"NETTPERUNIT"`
			NETTTOTAL    float64 `xml:"NETTTOTAL"`
			TAXTOTAL     float64 `xml:"TAXTOTAL"`
			COMPANYREf   string  `xml:"COMPANYREf"`
			SOURCE       string  `xml:"SOURCE"`
		}] `xml:"LEDGERANALYSISTRANS"`
		LEDGERANALYSISSUMMARYTRANS XMLMapStringStruct[[]struct {
			INVOICENUMBER     string  `xml:"INVOICENUMBER"`
			EntryType         string  `xml:"EntryType"`
			POREFERENCENUMBER string  `xml:"POREFERENCENUMBER"`
			COMPANYREF        string  `xml:"COMPANYREF"`
			COMPANYNAME       string  `xml:"COMPANYNAME"`
			DESCRIPTION       string  `xml:"DESCRIPTION"`
			GUESTNAME         string  `xml:"GUESTNAME"`
			INVOICEDATE       string  `xml:"INVOICEDATE"`
			GROSSTOTAL        float64 `xml:"GROSSTOTAL"`
			NETTTOTAL         float64 `xml:"NETTTOTAL"`
			TAXTOTAL          float64 `xml:"TAXTOTAL"`
			TAXID             string  `xml:"TAXID"`
			TYPEOFBUSINESS    string  `xml:"TYPEOFBUSINESS"`
			CITY              string  `xml:"CITY"`
			STREET            string  `xml:"STREET"`
			STREETNUMBER      string  `xml:"STREETNUMBER"`
			POSTCODE          string  `xml:"POSTCODE"`
			COUNTRY           string  `xml:"COUNTRY"`
		}] `xml:"LEDGERANALYSISSUMMARYTRANS"`
		LEDGERANALYSISSUMMARYTRANSWB XMLMapStringStruct[[]struct {
			INVOICENUMBER     string  `xml:"INVOICENUMBER"`
			EntryType         string  `xml:"EntryType"`
			POREFERENCENUMBER string  `xml:"POREFERENCENUMBER"`
			COMPANYREF        string  `xml:"COMPANYREF"`
			COMPANYNAME       string  `xml:"COMPANYNAME"`
			DESCRIPTION       string  `xml:"DESCRIPTION"`
			GUESTNAME         string  `xml:"GUESTNAME"`
			INVOICEDATE       string  `xml:"INVOICEDATE"`
			GROSSTOTAL        float64 `xml:"GROSSTOTAL"`
			NETTTOTAL         float64 `xml:"NETTTOTAL"`
			TAXTOTAL          float64 `xml:"TAXTOTAL"`
		}] `xml:"LEDGERANALYSISSUMMARYTRANSWB"`
	} `xml:"LEDGERANALYSIS"`
	PAYMENTSANALYSIS struct {
		PAYMENTSANALYSISTOTALS struct {
			PAYMENTSANALYSISITEM []struct {
				CODE        string  `xml:"CODE,attr"`
				DESCRIPTION string  `xml:"DESCRIPTION"`
				NETTTOTAL   float64 `xml:"NETTTOTAL"`
				TAXTOTAL    float64 `xml:"TAXTOTAL"`
				GROSSTOTAL  float64 `xml:"GROSSTOTAL"`
			} `xml:"PAYMENTSANALYSISITEM"`
		} `xml:"PAYMENTSANALYSISTOTALS"`
		PAYMENTSANALYSISTRANS struct {
			PAYMENTSANALYSISITEM []struct {
				CODE             string  `xml:"CODE,attr"`
				TRANSID          string  `xml:"TRANSID"`
				DESCRIPTION      string  `xml:"DESCRIPTION"`
				QUANTITY         float64 `xml:"QUANTITY"`
				GROSSPERUNIT     float64 `xml:"GROSSPERUNIT"`
				GROSSTOTAL       float64 `xml:"GROSSTOTAL"`
				NETTPERUNIT      float64 `xml:"NETTPERUNIT"`
				NETTTOTAL        float64 `xml:"NETTTOTAL"`
				TAXTOTAL         float64 `xml:"TAXTOTAL"`
				FORDATE          string  `xml:"FORDATE"`
				BOOKINGREFERENCE string  `xml:"BOOKINGREFERENCE"`
				ROOMPICKID       string  `xml:"ROOMPICKID"`
				FOLIOID          string  `xml:"FOLIOID"`
				FOLIOSPLITID     string  `xml:"FOLIOSPLITID"`
				CUSTOMERNAME     string  `xml:"CUSTOMERNAME"`
				CONTACTNAME      string  `xml:"CONTACTNAME"`
			} `xml:"PAYMENTSANALYSISITEM"`
		} `xml:"PAYMENTSANALYSISTRANS"`
	} `xml:"PAYMENTSANALYSIS"`
	TAXANALYSIS struct {
		TAXANALYSISITEM []struct {
			CODE        string  `xml:"CODE,attr"`
			DESCRIPTION string  `xml:"DESCRIPTION"`
			NETT        float64 `xml:"NETT"`
			TAX         float64 `xml:"TAX"`
			GROSS       float64 `xml:"GROSS"`
		} `xml:"TAXANALYSISITEM"`
	} `xml:"TAXANALYSIS"`
	INHOUSEANALYSIS struct {
		INHOUSEANALYSISTOTALS XMLMapStringStruct[struct {
			DESCRIPTION string  `xml:"DESCRIPTION"`
			NETTTOTAL   float64 `xml:"NETTTOTAL"`
			TAXTOTAL    float64 `xml:"TAXTOTAL"`
			GROSSTOTAL  float64 `xml:"GROSSTOTAL"`
		}] `xml:"INHOUSEANALYSISTOTALS"`
	} `xml:"INHOUSEANALYSIS"`
	OCCUPANCY struct {
		ITEM []struct {
			DESCRIPTION string `xml:"DESCRIPTION,attr"`
			GROSS       string `xml:"GROSS,attr"`
			NETT        string `xml:"NETT,attr"`
		} `xml:"ITEM"`
	} `xml:"OCCUPANCY"`
	OCCUPANCYBYMARKETSEGMENT struct {
		MARKETSEGMENT []struct {
			CODE string `xml:"CODE,attr"`
			ITEM []struct {
				DESCRIPTION string  `xml:"DESCRIPTION,attr"`
				GROSS       float64 `xml:"GROSS,attr"`
				NETT        float64 `xml:"NETT,attr"`
			} `xml:"ITEM"`
		} `xml:"MARKETSEGMENT"`
	} `xml:"OCCUPANCYBYMARKETSEGMENT"`
}

func (report *FinancialReportData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	b := []byte{}
	err := d.DecodeElement(&b, &start)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(b, &report.Document)
	if err != nil {
		return err
	}

	return nil
}

type XMLMapStringStruct[T any] map[string]T

func (mp *XMLMapStringStruct[T]) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*mp = XMLMapStringStruct[T]{}
	for {
		token, err := d.Token()
		if err != nil {
			// Quit for-loop when EOF is reached
			if err == io.EOF {
				break
			}

			return err
		}

		sToken, ok := token.(xml.StartElement)
		if !ok {
			continue
		}

		var t T
		err = d.DecodeElement(&t, &sToken)
		if err != nil {
			return err
		}

		(*mp)[(sToken.Name.Local)] = t
	}

	return nil
}

type SelectionCriteria struct {
	REPORTTITLE string `xml:"REPORTTITLE"`
}

type FinancialDocumentSummaryItem struct {
	XMLName        xml.Name `xml:"FinancialDocumentSummaryItem"`
	Text           string   `xml:",chardata"`
	BookRefRoomRef string   `xml:"BookRefRoomRef"`
	FolioId        struct {
		Text string `xml:",chardata"`
		Nil  string `xml:"nil,attr"`
	} `xml:"FolioId"`
	FolioSplitId struct {
		Text string `xml:",chardata"`
		Nil  string `xml:"nil,attr"`
	} `xml:"FolioSplitId"`
	DocumentRef       string   `xml:"DocumentRef"`
	DocumentType      string   `xml:"DocumentType"`
	DocumentTimestamp DateTime `xml:"DocumentTimestamp"`
	Forename          string   `xml:"Forename"`
	Surname           string   `xml:"Surname"`
	LedgerRef         string   `xml:"LedgerRef"`
	LedgerName        string   `xml:"LedgerName"`
	GrossTotal        float64  `xml:"GrossTotal"`
	NettTotal         float64  `xml:"NettTotal"`
	TaxTotal          float64  `xml:"TaxTotal"`
	ArrivalDate       DateTime `xml:"ArrivalDate"`
	DepartureDate     DateTime `xml:"DepartureDate"`
	AddressSource     string   `xml:"AddressSource"`
	AddressRef        string   `xml:"AddressRef"`
}

type DocumentDetail struct {
	XMLName         xml.Name `xml:"DocumentDetail"`
	Text            string   `xml:",chardata"`
	DocumentSummary struct {
		Text              string   `xml:",chardata"`
		BookRefRoomRef    string   `xml:"BookRefRoomRef"`
		FolioId           string   `xml:"FolioId"`
		FolioSplitId      string   `xml:"FolioSplitId"`
		DocumentRef       string   `xml:"DocumentRef"`
		DocumentType      string   `xml:"DocumentType"`
		DocumentTimestamp DateTime `xml:"DocumentTimestamp"`
		LedgerRef         string   `xml:"LedgerRef"`
		LedgerName        string   `xml:"LedgerName"`
		GrossTotal        float64  `xml:"GrossTotal"`
		NettTotal         float64  `xml:"NettTotal"`
		TaxTotal          float64  `xml:"TaxTotal"`
		ArrivalDate       DateTime `xml:"ArrivalDate"`
		DepartureDate     DateTime `xml:"DepartureDate"`
		AddressSource     string   `xml:"AddressSource"`
		AddressRef        string   `xml:"AddressRef"`
	} `xml:"DocumentSummary"`
	HeaderDetails struct {
		Text     string `xml:",chardata"`
		Title    string `xml:"Title"`
		Forename string `xml:"Forename"`
		Surname  string `xml:"Surname"`
		Address1 string `xml:"Address1"`
		Address2 string `xml:"Address2"`
		City     string `xml:"City"`
		County   string `xml:"County"`
		Country  string `xml:"Country"`
		Postcode string `xml:"Postcode"`
	} `xml:"HeaderDetails"`
	TransactionLines struct {
		Text                               string `xml:",chardata"`
		FinancialDocumentDetailTransaction []struct {
			Text                    string   `xml:",chardata"`
			ItemType                string   `xml:"ItemType"`
			EntryType               string   `xml:"EntryType"`
			Status                  string   `xml:"Status"`
			AnalCode                string   `xml:"AnalCode"`
			BookRef                 string   `xml:"BookRef"`
			Comments                string   `xml:"Comments"`
			ConversionId            string   `xml:"ConversionId"`
			ConvertedGrossTotal     string   `xml:"ConvertedGrossTotal"`
			Description             string   `xml:"Description"`
			EndDate                 DateTime `xml:"EndDate"`
			FinTransactionParam1    string   `xml:"FinTransactionParam1"`
			FinTransactionParam2    string   `xml:"FinTransactionParam2"`
			ForDate                 DateTime `xml:"ForDate"`
			GrossPerUnit            float64  `xml:"GrossPerUnit"`
			GrossTotal              float64  `xml:"GrossTotal"`
			IndexOfDate             string   `xml:"IndexOfDate"`
			IsFolioCharge           string   `xml:"IsFolioCharge"`
			NettPerUnit             float64  `xml:"NettPerUnit"`
			NettTotal               float64  `xml:"NettTotal"`
			NoOfNights              int      `xml:"NoOfNights"`
			OrderIndex              string   `xml:"OrderIndex"`
			PackageCode             string   `xml:"PackageCode"`
			PaymentCode             string   `xml:"PaymentCode"`
			Quantity                float64  `xml:"Quantity"`
			RoomPickId              string   `xml:"RoomPickId"`
			SourceTransId           string   `xml:"SourceTransId"`
			StartDate               DateTime `xml:"StartDate"`
			TaxReference            string   `xml:"TaxReference"`
			TaxTotal                string   `xml:"TaxTotal"`
			Timestamp               DateTime `xml:"Timestamp"`
			TransactionWasCategAs   string   `xml:"TransactionWasCategAs"`
			TransactionWasGroupedBy string   `xml:"TransactionWasGroupedBy"`
			TransId                 string   `xml:"TransId"`
			NettSequence            string   `xml:"NettSequence"`
		} `xml:"FinancialDocumentDetailTransaction"`
	} `xml:"TransactionLines"`
	Footer struct {
		Text                           string `xml:",chardata"`
		FinancialDocumentDetailTaxLine struct {
			Text        string  `xml:",chardata"`
			TaxRule     string  `xml:"TaxRule"`
			Description string  `xml:"Description"`
			NettTotal   float64 `xml:"NettTotal"`
			TaxTotal    float64 `xml:"TaxTotal"`
			GrossTotal  float64 `xml:"GrossTotal"`
		} `xml:"FinancialDocumentDetailTaxLine"`
	} `xml:"Footer"`
}

type GetCompanyProfile struct {
	CompanyRef              string `xml:"CompanyRef"`
	RFlag                   string `xml:"RFlag"`
	Name                    string `xml:"Name"`
	ModifiedUTC             string `xml:"ModifiedUTC"`
	CreatedUTC              string `xml:"CreatedUTC"`
	AddressLine1            string `xml:"AddressLine1"`
	AddressLine2            string `xml:"AddressLine2"`
	Street                  string `xml:"Street"`
	Area                    string `xml:"Area"`
	Town                    string `xml:"Town"`
	County                  string `xml:"County"`
	Country                 string `xml:"Country"`
	Postcode                string `xml:"Postcode"`
	TelNo                   string `xml:"TelNo"`
	FaxNo                   string `xml:"FaxNo"`
	FaxNo2                  string `xml:"FaxNo2"`
	EMail                   string `xml:"EMail"`
	ContactProfileRef       string `xml:"ContactProfileRef"`
	CreditFacility          string `xml:"CreditFacility"`
	CreditLimit             string `xml:"CreditLimit"`
	CompanyRegNum           string `xml:"CompanyRegNum"`
	ABTA                    string `xml:"ABTA"`
	ATOL                    string `xml:"ATOL"`
	CIF                     string `xml:"CIF"`
	TaxId                   string `xml:"TaxId"`
	TaxAuthority            string `xml:"TaxAuthority"`
	TypeOfBusiness          string `xml:"TypeOfBusiness"`
	AccountsPayableTitle    string `xml:"AccountsPayableTitle"`
	AccountsPayableForename string `xml:"AccountsPayableForename"`
	AccountsPayableSurname  string `xml:"AccountsPayableSurname"`
	AccountsPayableEmail    string `xml:"AccountsPayableEmail"`
	CompanyHierarchyType    string `xml:"CompanyHierarchyType"`
	CompanyHierarchy        string `xml:"CompanyHierarchy"`
	CustomAttributes        struct {
		ProfileCustomAttributes []struct {
			AttributeCode               string `xml:"AttributeCode"`
			Value                       string `xml:"Value"`
			Param1                      string `xml:"Param1"`
			Param2                      string `xml:"Param2"`
			ProfileAttributeCode        string `xml:"ProfileAttributeCode"`
			ProfileAttributeDescription string `xml:"ProfileAttributeDescription"`
			ProfileAttributeValue       string `xml:"ProfileAttributeValue"`
		} `xml:"ProfileCustomAttributes"`
	} `xml:"CustomAttributes"`
}
