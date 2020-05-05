package controllers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
	"gitlab.com/kevynestrada/bodywoks-ws/models"
	"gitlab.com/kevynestrada/bodywoks-ws/services"
	"gitlab.com/kevynestrada/bodywoks-ws/utils"
	"log"
	"net/http"
	"strconv"
)

type Request struct {
	Username    string
	Password    string
	FileName    string
	ContentFile string
}

func GetAllInvoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var invoiceView models.InvoiceView
	invoices, err := invoiceView.FindAll()
	if err != nil {
		log.Print(err)
		http.Error(w, utils.ConvertObjToJson(err), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(invoices)
}

func GetInvoice(w http.ResponseWriter, r *http.Request)  {

}

func CreateInvoice(w http.ResponseWriter, r *http.Request)  {
	//w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var quotation models.Invoice
	if e := json.NewDecoder(r.Body).Decode(&quotation); e != nil {
		http.Error(w, e.Error(), http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()
	id, err := quotation.Create()
	if err != nil {
		log.Print(err)
		http.Error(w, utils.ConvertObjToJson(err), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(id)
}
func GetViewinvoice(w http.ResponseWriter, r *http.Request)  {

}

func UpdateInvoice(w http.ResponseWriter, r *http.Request)  {

}

func DeleteInvoice(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	var order models.Invoice
	response := order.Delete(id)
	http.Error(w, utils.ConvertObjToJson(response), http.StatusOK)
}

func AllowInvoice(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, DELETE, PUT")
	log.Print("AllowQuotation")
	if r.Method == "OPTIONS" {
		return
	}
}

func CreateInvoiceSunat(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var response entities.Response
	response.Code = 0
	response.Message = "OK"

	// DATA
	ruc := "20394008398"
	invoiceId := "F001-00000099"
	doctype := "01"
	username := "MREPTO01"//"MODDATOS"
	password := "12345678"//"moddatos"

	//var invoice entities.InvoiceSunat
	var invoice entities.Invoice2
	invoice.Xmlns = "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2"
	invoice.Cac = "urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2"
	invoice.Cbc = "urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2"
	invoice.Ccts = "urn:un:unece:uncefact:documentation:2"
	invoice.Ds = "http://www.w3.org/2000/09/xmldsig#"
	invoice.Ext = "urn:oasis:names:specification:ubl:schema:xsd:CommonExtensionComponents-2"
	invoice.Qdt = "urn:oasis:names:specification:ubl:schema:xsd:QualifiedDatatypes-2"
	invoice.Sac = "urn:sunat:names:specification:ubl:peru:schema:xsd:SunatAggregateComponents-1"
	invoice.Udt = "urn:un:unece:uncefact:data:specification:UnqualifiedDataTypesSchemaModule:2"
	invoice.Xsi = "http://www.w3.org/2001/XMLSchema-instance"

	invoice.UBLVersionID = "2.1"
	invoice.CustomizationID = "2.0"
	invoice.ID = invoiceId
	invoice.IssueDate = "2019-09-25"
	invoice.IssueTime = "07:30:40"
	invoice.DueDate = "2019-09-25"

	invoice.InvoiceTypeCode.ListID = "0101"
	invoice.InvoiceTypeCode.ListAgencyName = "PE:SUNAT"
	invoice.InvoiceTypeCode.ListName = "PE:SUNAT"
	invoice.InvoiceTypeCode.ListURI = "urn:pe:gob:sunat:cpe:see:gem:catalogos:catalogo01"
	invoice.InvoiceTypeCode.Text = "01"

	invoice.Note.LanguageLocaleID = "1000"
	invoice.Note.Text = "TREINTA Y TRES CON 00/100 SOLES"

	invoice.LineCountNumeric = "2"

	invoice.DocumentCurrencyCode.ListID = "ISO 4217 Alpha"
	invoice.DocumentCurrencyCode.ListName = "Currency"
	invoice.DocumentCurrencyCode.ListAgencyName = "United Nations Economic Commission for Europe"
	invoice.DocumentCurrencyCode.Text = "PEN"

	invoice.Signature.ID = invoiceId
	invoice.Signature.SignatoryParty.PartyIdentification.ID = ruc
	invoice.Signature.SignatoryParty.PartyName.Name = "LLAMA.PE SA"
	invoice.Signature.DigitalSignatureAttachment.ExternalReference.URI = ruc + "-" + invoiceId

	invoice.AccountingSupplierParty.Party.PartyIdentification.ID.SchemeID = "6"
	invoice.AccountingSupplierParty.Party.PartyIdentification.ID.SchemeName = "SUNAT:Identificador de Documento de Identidad"
	invoice.AccountingSupplierParty.Party.PartyIdentification.ID.SchemeAgencyName = "PE:SUNAT"
	invoice.AccountingSupplierParty.Party.PartyIdentification.ID.SchemeURI = "urn:pe:gob:sunat:cpe:see:gem:catalogos:catalogo06"

	invoice.AccountingCustomerParty.Party.PartyIdentification.ID.SchemeID = "6"
	invoice.AccountingCustomerParty.Party.PartyIdentification.ID.SchemeName = "SUNAT:Identificador de Documento de Identidad"
	invoice.AccountingCustomerParty.Party.PartyIdentification.ID.SchemeAgencyName = "PE:SUNAT"
	invoice.AccountingCustomerParty.Party.PartyIdentification.ID.SchemeURI = "urn:pe:gob:sunat:cpe:see:gem:catalogos:catalogo06"

	err := utils.CreateSignature(&invoice)
	if err != nil {
		response.Code = 1
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	var file utils.File
	file.Name = ruc + "-" + doctype + "-" + invoiceId + ".xml"
	//file.Body = invoice.String()
	file.Body = `<?xml version="1.0" encoding="utf-8"?><Invoice xmlns="urn:oasis:names:specification:ubl:schema:xsd:Invoice-2" xmlns:cac="urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2" xmlns:cbc="urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2" xmlns:ccts="urn:un:unece:uncefact:documentation:2" xmlns:ds="http://www.w3.org/2000/09/xmldsig#" xmlns:ext="urn:oasis:names:specification:ubl:schema:xsd:CommonExtensionComponents-2" xmlns:qdt="urn:oasis:names:specification:ubl:schema:xsd:QualifiedDatatypes-2" xmlns:sac="urn:sunat:names:specification:ubl:peru:schema:xsd:SunatAggregateComponents-1" xmlns:udt="urn:un:unece:uncefact:data:specification:UnqualifiedDataTypesSchemaModule:2" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><ext:UBLExtensions><ext:UBLExtension><ext:ExtensionContent><Signature Id="SignOpenInvoice" xmlns="http://www.w3.org/2000/09/xmldsig#"><SignedInfo><CanonicalizationMethod Algorithm="http://www.w3.org/TR/2001/REC-xml-c14n-20010315" /><SignatureMethod Algorithm="http://www.w3.org/2000/09/xmldsig#rsa-sha1" /><Reference URI=""><Transforms><Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature" /></Transforms><DigestMethod Algorithm="http://www.w3.org/2000/09/xmldsig#sha1" /><DigestValue>GSUbZ1VhsKP0+AJ+etObAqJJJfQ=</DigestValue></Reference></SignedInfo><SignatureValue>Goy68TagnzEzXkem8X5EKOczTAAvFx/Hj291OLcLCHzHnfVCEMoUS7wfcC5O82ySb9Blgip5LuaKdSIftYOg6poQ29EXllKYy5kIcApogsmewi2f7573cy2VXF3xPELbgPQRHi+Uprg0CYy2GlpKDCCn9UeBIxnnLpHfmIM2XVdAY2i2/JqB8jKZZaEWMTOsu9wioGcb2i+AJJWqv5/axFeYvH8HWomVwpMdlz4sd72uiVbrcoR2F3oCpZDujB33sRP18305udVZizWXcyXq1aRBswofsFzNsy560eJvtoUs0wwEqycX9RaY3snuJcpdy9MX/kK5m3sftDx29PTKvA==</SignatureValue><KeyInfo><X509Data><X509SubjectName>CN=RAQUEL CASTILLO ZUTA + SERIALNUMBER="RUC: 20394008398, DNI: 40726239" + O=UCAYALI MULTIREPUESTOS E.I.R.L., L=Coronel Portillo, S=Ucayali, C=PE</X509SubjectName><X509Certificate>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</X509Certificate></X509Data></KeyInfo></Signature></ext:ExtensionContent></ext:UBLExtension></ext:UBLExtensions><cbc:UBLVersionID>2.1</cbc:UBLVersionID><cbc:CustomizationID>2.0</cbc:CustomizationID><cbc:ID>F001-00000099</cbc:ID><cbc:IssueDate>2019-09-25</cbc:IssueDate><cbc:IssueTime>07:30:40</cbc:IssueTime><cbc:DueDate>2019-09-25</cbc:DueDate><cbc:InvoiceTypeCode listID="0101" listAgencyName="PE:SUNAT" listName="PE:SUNAT" listURI="urn:pe:gob:sunat:cpe:see:gem:catalogos:catalogo01">01</cbc:InvoiceTypeCode><cbc:Note languageLocaleID="1000">TREINTA Y TRES CON 00/100 SOLES</cbc:Note><cbc:DocumentCurrencyCode listID="ISO 4217 Alpha" listName="Currency" listAgencyName="United Nations Economic Commission for Europe">PEN</cbc:DocumentCurrencyCode><cbc:LineCountNumeric>2</cbc:LineCountNumeric><cac:Signature><cbc:ID>F001-00000099</cbc:ID><cac:SignatoryParty><cac:PartyIdentification><cbc:ID>20394008398</cbc:ID></cac:PartyIdentification><cac:PartyName><cbc:Name>UCAYALI MULTIREPUESTOS E.I.R.L.</cbc:Name></cac:PartyName></cac:SignatoryParty><cac:DigitalSignatureAttachment><cac:ExternalReference><cbc:URI>20394008398-F001-00000099</cbc:URI></cac:ExternalReference></cac:DigitalSignatureAttachment></cac:Signature><cac:AccountingSupplierParty><cac:Party><cac:PartyIdentification><cbc:ID schemeID="6" schemeName="SUNAT:Identificador de Documento de Identidad" schemeAgencyName="PE:SUNAT" schemeURI="urn:pe:gob:sunat:cpe:see:gem:catalogos:catalogo06">20394008398</cbc:ID></cac:PartyIdentification><cac:PartyName><cbc:Name><![CDATA[UCAYALI MULTIREPUESTOS E.I.R.L.]]></cbc:Name></cac:PartyName><cac:PartyLegalEntity><cbc:RegistrationName><![CDATA[UCAYALI MULTIREPUESTOS E.I.R.L.]]></cbc:RegistrationName><cac:RegistrationAddress><cbc:AddressTypeCode>0000</cbc:AddressTypeCode></cac:RegistrationAddress></cac:PartyLegalEntity></cac:Party></cac:AccountingSupplierParty><cac:AccountingCustomerParty><cac:Party><cac:PartyIdentification><cbc:ID schemeID="6" schemeName="SUNAT:Identificador de Documento de Identidad" schemeAgencyName="PE:SUNAT" schemeURI="urn:pe:gob:sunat:cpe:see:gem:catalogos:catalogo06">20393867397</cbc:ID></cac:PartyIdentification><cac:PartyLegalEntity><cbc:RegistrationName><![CDATA[AGUA DE MESA SAN MARTINENSE E.I.R.L]]></cbc:RegistrationName></cac:PartyLegalEntity></cac:Party></cac:AccountingCustomerParty><cac:TaxTotal><cbc:TaxAmount currencyID="PEN">0.00</cbc:TaxAmount><cac:TaxSubtotal><cbc:TaxableAmount currencyID="PEN">33.00</cbc:TaxableAmount><cbc:TaxAmount currencyID="PEN">0.00</cbc:TaxAmount><cac:TaxCategory><cbc:ID schemeID="UN/ECE 5305" schemeName="Tax Category Identifier" schemeAgencyName="United Nations Economic Commission for Europe">E</cbc:ID><cac:TaxScheme><cbc:ID schemeID="UN/ECE 5305" schemeAgencyID="6">9997</cbc:ID><cbc:Name>EXO</cbc:Name><cbc:TaxTypeCode>VAT</cbc:TaxTypeCode></cac:TaxScheme></cac:TaxCategory></cac:TaxSubtotal></cac:TaxTotal><cac:LegalMonetaryTotal><cbc:PayableAmount currencyID="PEN">33.00</cbc:PayableAmount></cac:LegalMonetaryTotal><cac:InvoiceLine><cbc:ID>1</cbc:ID><cbc:InvoicedQuantity unitCode="NIU" unitCodeListID="UN/ECE rec 20" unitCodeListAgencyName="United Nations Economic Commission forEurope">1.00</cbc:InvoicedQuantity><cbc:LineExtensionAmount currencyID="PEN">11.00</cbc:LineExtensionAmount><cac:PricingReference><cac:AlternativeConditionPrice><cbc:PriceAmount currencyID="PEN">11.00</cbc:PriceAmount><cbc:PriceTypeCode listName="SUNAT:Indicador de Tipo de Precio" listAgencyName="PE:SUNAT" listURI="urn:pe:gob:sunat:cpe:see:gem:catalogos:catalogo16">01</cbc:PriceTypeCode></cac:AlternativeConditionPrice></cac:PricingReference><cac:TaxTotal><cbc:TaxAmount currencyID="PEN">0.00</cbc:TaxAmount><cac:TaxSubtotal><cbc:TaxableAmount currencyID="PEN">11.00</cbc:TaxableAmount><cbc:TaxAmount currencyID="PEN">0.00</cbc:TaxAmount><cac:TaxCategory><cbc:ID schemeID="UN/ECE 5305" schemeName="Tax Category Identifier" schemeAgencyName="United Nations Economic Commission for Europe">E</cbc:ID><cbc:Percent>18.00</cbc:Percent><cbc:TaxExemptionReasonCode listAgencyName="PE:SUNAT" listName="SUNAT:Codigo de Tipo de Afectación del IGV" listURI="urn:pe:gob:sunat:cpe:see:gem:catalogos:catalogo07">20</cbc:TaxExemptionReasonCode><cac:TaxScheme><cbc:ID schemeID="UN/ECE 5153" schemeName="Tax Scheme Identifier" schemeAgencyName="United Nations Economic Commission for Europe">9997</cbc:ID><cbc:Name>EXO</cbc:Name><cbc:TaxTypeCode>VAT</cbc:TaxTypeCode></cac:TaxScheme></cac:TaxCategory></cac:TaxSubtotal></cac:TaxTotal><cac:Item><cbc:Description>ACEITE CAM2 MONOGRADO *</cbc:Description><cac:SellersItemIdentification><cbc:ID>00000216</cbc:ID></cac:SellersItemIdentification></cac:Item><cac:Price><cbc:PriceAmount currencyID="PEN">11.00</cbc:PriceAmount></cac:Price></cac:InvoiceLine><cac:InvoiceLine><cbc:ID>2</cbc:ID><cbc:InvoicedQuantity unitCode="NIU" unitCodeListID="UN/ECE rec 20" unitCodeListAgencyName="United Nations Economic Commission forEurope">1.00</cbc:InvoicedQuantity><cbc:LineExtensionAmount currencyID="PEN">22.00</cbc:LineExtensionAmount><cac:PricingReference><cac:AlternativeConditionPrice><cbc:PriceAmount currencyID="PEN">22.00</cbc:PriceAmount><cbc:PriceTypeCode listName="SUNAT:Indicador de Tipo de Precio" listAgencyName="PE:SUNAT" listURI="urn:pe:gob:sunat:cpe:see:gem:catalogos:catalogo16">01</cbc:PriceTypeCode></cac:AlternativeConditionPrice></cac:PricingReference><cac:TaxTotal><cbc:TaxAmount currencyID="PEN">0.00</cbc:TaxAmount><cac:TaxSubtotal><cbc:TaxableAmount currencyID="PEN">22.00</cbc:TaxableAmount><cbc:TaxAmount currencyID="PEN">0.00</cbc:TaxAmount><cac:TaxCategory><cbc:ID schemeID="UN/ECE 5305" schemeName="Tax Category Identifier" schemeAgencyName="United Nations Economic Commission for Europe">E</cbc:ID><cbc:Percent>18.00</cbc:Percent><cbc:TaxExemptionReasonCode listAgencyName="PE:SUNAT" listName="SUNAT:Codigo de Tipo de Afectación del IGV" listURI="urn:pe:gob:sunat:cpe:see:gem:catalogos:catalogo07">20</cbc:TaxExemptionReasonCode><cac:TaxScheme><cbc:ID schemeID="UN/ECE 5153" schemeName="Tax Scheme Identifier" schemeAgencyName="United Nations Economic Commission for Europe">9997</cbc:ID><cbc:Name>EXO</cbc:Name><cbc:TaxTypeCode>VAT</cbc:TaxTypeCode></cac:TaxScheme></cac:TaxCategory></cac:TaxSubtotal></cac:TaxTotal><cac:Item><cbc:Description>ACEITE ACTEVO CASTROL 20W 50 4T *</cbc:Description><cac:SellersItemIdentification><cbc:ID>00003264</cbc:ID></cac:SellersItemIdentification></cac:Item><cac:Price><cbc:PriceAmount currencyID="PEN">22.00</cbc:PriceAmount></cac:Price></cac:InvoiceLine></Invoice>`
	var request Request
	request.Username = ruc + username
	request.Password = password
	request.FileName = ruc + "-" + doctype + "-" + invoiceId + ".zip"
	request.ContentFile = base64.StdEncoding.EncodeToString(utils.StreamZip(file))
	responseBill, err := services.SendBill(request)
	if err != nil {
		response.Code = 1
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	if responseBill.Body.Fault != (services.Fault{}) {
		response.Code = 1
		response.Message = responseBill.Body.Fault.Faultstring
		json.NewEncoder(w).Encode(response)
		return
	}
	msg, err := base64.StdEncoding.DecodeString(responseBill.Body.SendBillResponse.ApplicationResponse)
	if err != nil {
		log.Print(err)
	}
	sunatXmlResponse, err := utils.StreamUnzip(msg)
	if err != nil {
		response.Code = 1
		response.Message = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Code = 0
	response.Message = sunatXmlResponse.Body
	json.NewEncoder(w).Encode(response)
}