package services

import (
	"bytes"
	"encoding/xml"
	template2 "html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type SendBillResponse struct {
	Text                string `xml:",chardata"`
	Br                  string `xml:"br,attr"`
	ApplicationResponse string `xml:"applicationResponse"`
}
type Fault struct {
	Text        string `xml:",chardata"`
	Faultcode   string `xml:"faultcode"`
	Faultstring string `xml:"faultstring"`
}
type Response struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	SoapEnv string   `xml:"soap-env,attr"`
	Xsi     string   `xml:"xsi,attr"`
	Wsse    string   `xml:"wsse,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text  string `xml:",chardata"`
		SendBillResponse SendBillResponse `xml:"sendBillResponse"`
		Fault Fault `xml:"Fault"`
	} `xml:"Body"`
}


func SendBill(request interface{}) (Response, error) {
	var response Response
	template, err := template2.New("").Parse(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.sunat.gob.pe" xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"><soapenv:Header><wsse:Security><wsse:UsernameToken><wsse:Username>{{.Username}}</wsse:Username><wsse:Password>{{.Password}}</wsse:Password></wsse:UsernameToken></wsse:Security></soapenv:Header><soapenv:Body><ser:sendBill><fileName>{{.FileName}}</fileName><contentFile>{{.ContentFile}}</contentFile></ser:sendBill></soapenv:Body></soapenv:Envelope>`)
	if err != nil {
		log.Print(err)
		return response, nil
	}
	doc := &bytes.Buffer{}
	err = template.Execute(doc, request)
	if err != nil {
		return response, err
	}
	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	if err != nil {
		return response, err
	}
	req, err := http.NewRequest("POST","https://e-beta.sunat.gob.pe/ol-ti-itcpfegem-beta/billService", bytes.NewBuffer([]byte(doc.String())))
	if err != nil {
		return response, err
	}
	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", "sendBill")
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return response, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}