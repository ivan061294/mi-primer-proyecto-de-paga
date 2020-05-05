package entities

import (
	"encoding/xml"
	"time"
	"log"
)

type InvoiceDetail struct {
	Id           int64      `json:"id"`
	InvoiceId    int64      `json:"invoiceid"`
	ProductId    int64      `json:"product"`
	Description  string     `json:"description"`
	Amount       int64      `json:"quantity"`
	Price        float64    `json:"unitprice"`
}

type Invoice struct {
	Id           int64 `json:"id"`
	EmployeeId  int64 `json:"employeeid"`
	CustomerId  int64 `json:"customerid"`
	QuotationId int64 `json:"quotationid"`
	Contact      string `json:"contact"`
	Status       string  `json:"status"`
	Currency     string `json:"currency"`
	Total        float64 `json:"total"`
	Regdate      time.Time `json:"regdate"`
	Observation  string `json:"observation"`
	Detail  []InvoiceDetail `json:"detail"`
}

type InvoiceView struct {
	Id          int64   `json:"id"`
	Seller      string  `json:"seller"`
	Customer    string  `json:"customer"`
	Doctype     string  `json:"doctype"`
	Docnum      string  `json:"docnum"`
	Issue       time.Time `json:"issue"`
	Contact     string  `json:"contact"`
	Status      string  `json:"status"`
	Currency    string  `json:"currency"`
	Total       float64 `json:"total"`
	Observation string  `json:"observation"`
	Xmlsign     string  `json:"xmlsign"`
	Xmlsunat    string  `json:"xmlsunat"`
}

type InvoiceLine struct {
	Text             string `xml:",chardata"`
	ID               string `xml:"ID"`
	InvoicedQuantity struct {
		Text                   string `xml:",chardata"`
		UnitCode               string `xml:"unitCode,attr"`
		UnitCodeListID         string `xml:"unitCodeListID,attr"`
		UnitCodeListAgencyName string `xml:"unitCodeListAgencyName,attr"`
	} `xml:"InvoicedQuantity"`
	LineExtensionAmount struct {
		Text       string `xml:",chardata"`
		CurrencyID string `xml:"currencyID,attr"`
	} `xml:"LineExtensionAmount"`
	PricingReference struct {
		Text                      string `xml:",chardata"`
		AlternativeConditionPrice struct {
			Text        string `xml:",chardata"`
			PriceAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"PriceAmount"`
			PriceTypeCode struct {
				Text           string `xml:",chardata"`
				ListName       string `xml:"listName,attr"`
				ListAgencyName string `xml:"listAgencyName,attr"`
				ListURI        string `xml:"listURI,attr"`
			} `xml:"PriceTypeCode"`
		} `xml:"AlternativeConditionPrice"`
	} `xml:"PricingReference"`
	TaxTotal struct {
		Text      string `xml:",chardata"`
		TaxAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"TaxAmount"`
		TaxSubtotal struct {
			Text          string `xml:",chardata"`
			TaxableAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"TaxableAmount"`
			TaxAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"TaxAmount"`
			TaxCategory struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text             string `xml:",chardata"`
					SchemeID         string `xml:"schemeID,attr"`
					SchemeName       string `xml:"schemeName,attr"`
					SchemeAgencyName string `xml:"schemeAgencyName,attr"`
				} `xml:"ID"`
				Percent                string `xml:"Percent"`
				TaxExemptionReasonCode struct {
					Text           string `xml:",chardata"`
					ListAgencyName string `xml:"listAgencyName,attr"`
					ListName       string `xml:"listName,attr"`
					ListURI        string `xml:"listURI,attr"`
				} `xml:"TaxExemptionReasonCode"`
				TaxScheme struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text             string `xml:",chardata"`
						SchemeID         string `xml:"schemeID,attr"`
						SchemeName       string `xml:"schemeName,attr"`
						SchemeAgencyName string `xml:"schemeAgencyName,attr"`
					} `xml:"ID"`
					Name        string `xml:"Name"`
					TaxTypeCode string `xml:"TaxTypeCode"`
				} `xml:"TaxScheme"`
			} `xml:"TaxCategory"`
		} `xml:"TaxSubtotal"`
	} `xml:"TaxTotal"`
	Item struct {
		Text                      string `xml:",chardata"`
		Description               string `xml:"Description"`
		SellersItemIdentification struct {
			Text string `xml:",chardata"`
			ID   string `xml:"ID"`
		} `xml:"SellersItemIdentification"`
	} `xml:"Item"`
	Price struct {
		Text        string `xml:",chardata"`
		PriceAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"PriceAmount"`
	} `xml:"Price"`
}

type SignedInfo struct {
	Text                   string `xml:",chardata"`
	CanonicalizationMethod struct {
		Text      string `xml:",chardata"`
		Algorithm string `xml:"Algorithm,attr"`
	} `xml:"CanonicalizationMethod"`
	SignatureMethod struct {
		Text      string `xml:",chardata"`
		Algorithm string `xml:"Algorithm,attr"`
	} `xml:"SignatureMethod"`
	Reference struct {
		Text       string `xml:",chardata"`
		URI        string `xml:"URI,attr"`
		Transform struct {
			Text      string `xml:",chardata"`
			Algorithm string `xml:"Algorithm,attr"`
		} `xml:"Transforms>Transform"`
		DigestMethod struct {
			Text      string `xml:",chardata"`
			Algorithm string `xml:"Algorithm,attr"`
		} `xml:"DigestMethod"`
		DigestValue string `xml:"DigestValue"`
	} `xml:"Reference"`
}

type Signature struct {
	Text       string `xml:",chardata"`
	ID         string `xml:"Id,attr"`
	Xmlns      string `xml:"xmlns,attr"`
	SignedInfo SignedInfo `xml:"SignedInfo"`
	SignatureValue string `xml:"SignatureValue"`
	KeyInfo        struct {
		Text     string `xml:",chardata"`
		X509Data struct {
			Text            string `xml:",chardata"`
			X509SubjectName string `xml:"X509SubjectName"`
			X509Certificate string `xml:"X509Certificate"`
		} `xml:"X509Data"`
	} `xml:"KeyInfo"`
}

type UBLExtensions struct {
	Text         string `xml:",chardata"`
	UBLExtension struct {
		Text      string `xml:",chardata"`
		Signature *Signature `xml:"Signature"`
	} `xml:"ext:UBLExtension>ext:ExtensionContent"`
}

type Invoice2 struct {
	XMLName       xml.Name `xml:"Invoice"`
	Text          string   `xml:",chardata"`
	Xmlns         string   `xml:"xmlns,attr"`
	Cac           string   `xml:"xmlns:cac,attr"`
	Cbc           string   `xml:"xmlns:cbc,attr"`
	Ccts          string   `xml:"xmlns:ccts,attr"`
	Ds            string   `xml:"xmlns:ds,attr"`
	Ext           string   `xml:"xmlns:ext,attr"`
	Qdt           string   `xml:"xmlns:qdt,attr"`
	Sac           string   `xml:"xmlns:sac,attr"`
	Udt           string   `xml:"xmlns:udt,attr"`
	Xsi           string   `xml:"xmlns:xsi,attr"`
	UBLExtensions UBLExtensions `xml:"ext:UBLExtensions,omitempty"`
	/*UBLExtensions struct {
		Signature *Signature `xml:"Signature"`
	} `xml:"ext:UBLExtensions>ext:UBLExtension>ext:ExtensionContent"`*/
	UBLVersionID    string `xml:"cbc:UBLVersionID"`
	CustomizationID string `xml:"cbc:CustomizationID"`
	ID              string `xml:"cbc:ID"`
	IssueDate       string `xml:"cbc:IssueDate"`
	IssueTime       string `xml:"cbc:IssueTime"`
	DueDate         string `xml:"cbc:DueDate"`
	InvoiceTypeCode struct {
		Text           string `xml:",chardata"`
		ListID         string `xml:"listID,attr"`
		ListAgencyName string `xml:"listAgencyName,attr"`
		ListName       string `xml:"listName,attr"`
		ListURI        string `xml:"listURI,attr"`
	} `xml:"cbc:InvoiceTypeCode"`
	Note struct {
		Text             string `xml:",chardata"`
		LanguageLocaleID string `xml:"languageLocaleID,attr"`
	} `xml:"cbc:Note"`
	DocumentCurrencyCode struct {
		Text           string `xml:",chardata"`
		ListID         string `xml:"listID,attr"`
		ListName       string `xml:"listName,attr"`
		ListAgencyName string `xml:"listAgencyName,attr"`
	} `xml:"cbc:DocumentCurrencyCode"`
	LineCountNumeric string `xml:"cbc:LineCountNumeric"`
	Signature        struct {
		Text           string `xml:",chardata"`
		ID             string `xml:"cbc:ID"`
		SignatoryParty struct {
			Text                string `xml:",chardata"`
			PartyIdentification struct {
				Text string `xml:",chardata"`
				ID   string `xml:"cbc:ID"`
			} `xml:"cac:PartyIdentification"`
			PartyName struct {
				Text string `xml:",chardata"`
				Name string `xml:"cbc:Name"`
			} `xml:"cac:PartyName"`
		} `xml:"cac:SignatoryParty"`
		DigitalSignatureAttachment struct {
			Text              string `xml:",chardata"`
			ExternalReference struct {
				Text string `xml:",chardata"`
				URI  string `xml:"cbc:URI"`
			} `xml:"cac:ExternalReference"`
		} `xml:"cac:DigitalSignatureAttachment"`
	} `xml:"cac:Signature"`
	AccountingSupplierParty struct {
		Text  string `xml:",chardata"`
		Party struct {
			Text                string `xml:",chardata"`
			PartyIdentification struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text             string `xml:",chardata"`
					SchemeID         string `xml:"schemeID,attr"`
					SchemeName       string `xml:"schemeName,attr"`
					SchemeAgencyName string `xml:"schemeAgencyName,attr"`
					SchemeURI        string `xml:"schemeURI,attr"`
				} `xml:"cbc:ID"`
			} `xml:"cac:PartyIdentification"`
			PartyName struct {
				Text string `xml:",chardata"`
				Name string `xml:"cbc:Name"`
			} `xml:"cac:PartyName"`
			PartyLegalEntity struct {
				Text                string `xml:",chardata"`
				RegistrationName    string `xml:"cbc:RegistrationName"`
				RegistrationAddress struct {
					Text            string `xml:",chardata"`
					AddressTypeCode string `xml:"cbc:AddressTypeCode"`
				} `xml:"cac:RegistrationAddress"`
			} `xml:"cac:PartyLegalEntity"`
		} `xml:"cac:Party"`
	} `xml:"cac:AccountingSupplierParty"`
	AccountingCustomerParty struct {
		Text  string `xml:",chardata"`
		Party struct {
			Text                string `xml:",chardata"`
			PartyIdentification struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text             string `xml:",chardata"`
					SchemeID         string `xml:"schemeID,attr"`
					SchemeName       string `xml:"schemeName,attr"`
					SchemeAgencyName string `xml:"schemeAgencyName,attr"`
					SchemeURI        string `xml:"schemeURI,attr"`
				} `xml:"cbc:ID"`
			} `xml:"cac:PartyIdentification"`
			PartyLegalEntity struct {
				Text             string `xml:",chardata"`
				RegistrationName string `xml:"cbc:RegistrationName"`
			} `xml:"cac:PartyLegalEntity"`
		} `xml:"cac:Party"`
	} `xml:"cac:AccountingCustomerParty"`
	TaxTotal struct {
		Text      string `xml:",chardata"`
		TaxAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"cbc:TaxAmount"`
		TaxSubtotal struct {
			Text          string `xml:",chardata"`
			TaxableAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"cbc:TaxableAmount"`
			TaxAmount struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"cbc:TaxAmount"`
			TaxCategory struct {
				Text string `xml:",chardata"`
				ID   struct {
					Text             string `xml:",chardata"`
					SchemeID         string `xml:"schemeID,attr"`
					SchemeName       string `xml:"schemeName,attr"`
					SchemeAgencyName string `xml:"schemeAgencyName,attr"`
				} `xml:"cbc:ID"`
				TaxScheme struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text           string `xml:",chardata"`
						SchemeID       string `xml:"schemeID,attr"`
						SchemeAgencyID string `xml:"schemeAgencyID,attr"`
					} `xml:"cbc:ID"`
					Name        string `xml:"cbc:Name"`
					TaxTypeCode string `xml:"cbc:TaxTypeCode"`
				} `xml:"cac:TaxScheme"`
			} `xml:"cac:TaxCategory"`
		} `xml:"cac:TaxSubtotal"`
	} `xml:"cac:TaxTotal"`
	LegalMonetaryTotal struct {
		Text          string `xml:",chardata"`
		PayableAmount struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"cbc:PayableAmount"`
	} `xml:"cac:LegalMonetaryTotal"`
	InvoiceLine []InvoiceLine `xml:"cac:InvoiceLine"`
}
func (invoice Invoice2) String() string {
	var strXml []byte
	strXml, err := xml.Marshal(invoice)
	if err != nil {
		log.Print(err)
		return string(strXml)
	}
	return "<?xml version=\"1.0\" encoding=\"utf-8\"?>" + string(strXml)
}