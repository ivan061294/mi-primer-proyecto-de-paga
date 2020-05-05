package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"gitlab.com/kevynestrada/bodywoks-ws/entities"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"log"
)

func PFXToCER(pfxfilename string, password string) (interface{}, *x509.Certificate, error) {
	pfxData, err := ioutil.ReadFile(pfxfilename)
	if err != nil {
		return nil, nil, err
	}
	priv, cert, err := pkcs12.Decode(pfxData, password)
	if err != nil {
		return nil, nil, err
	}
	if err := priv.(*rsa.PrivateKey).Validate(); err != nil {
		return nil, nil, err
	}
	return priv, cert, nil
}

func GenerateHash(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}
func Digest(doc []byte) string {
	sum := GenerateHash(doc)
	return base64.StdEncoding.EncodeToString(sum)
}
func Sign(data []byte, cert *rsa.PrivateKey) (string, error) {
	sum := GenerateHash(data)
	sig, err := cert.Sign(rand.Reader, sum, crypto.SHA256)
	if err != nil {
		return "", err
	}
	err = cert.Validate()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}
func CreateSignature(invoice *entities.Invoice2) error {
	var signature entities.Signature
	var signedInfo entities.SignedInfo

	//log.Print(invoice.String())

	pfxfile := "/home/kevyn/Descargas/certificado/LLAMA-PE-CERTIFICADO-DEMO-20454870591.pfx"
	keyPem, certPem, err := PFXToCER(pfxfile, "QERpY2llbWJyZTIwMTkrK1NBQkFETzI4Kw==")
	if err != nil {
		log.Print(err)
		return err
	}

	log.Print("Generacion hash del xml")

	signedInfo.Reference.DigestValue = Digest([]byte(invoice.String()))//Digest(byteDigest) //OK

	//Firma SignedInfo
	signedInfo.CanonicalizationMethod.Algorithm = "http://www.w3.org/TR/2001/REC-xml-c14n-20010315"
	signedInfo.SignatureMethod.Algorithm = "http://www.w3.org/2000/09/xmldsig#rsa-sha1"
	signedInfo.Reference.Transform.Algorithm = "http://www.w3.org/2000/09/xmldsig#enveloped-signature"
	signedInfo.Reference.DigestMethod.Algorithm = "http://www.w3.org/2001/04/xmlenc#sha256"//"http://www.w3.org/2000/09/xmldsig#sha1"

	byteSignedInfo, err := xml.Marshal(signedInfo)
	if err != nil {
		return err
	}

	byteSignedInfo = []byte(hex.EncodeToString(byteSignedInfo))

	log.Print("Firma del elemento SignedInfo")
	signatureValue, err := Sign(byteSignedInfo, keyPem.(*rsa.PrivateKey)) //OK
	if err != nil {
		log.Print(err)
		return err
	}

	log.Print("Conversion a base64 del certificado")
	certificate := base64.StdEncoding.EncodeToString(certPem.Raw)

	signature.Xmlns = "http://www.w3.org/2000/09/xmldsig#"
	//signature.ID = "SignOpenInvoice"
	signature.SignedInfo = signedInfo
	signature.SignatureValue = signatureValue
	signature.KeyInfo.X509Data.X509SubjectName = certPem.Subject.String()
	signature.KeyInfo.X509Data.X509Certificate = certificate

	invoice.UBLExtensions.UBLExtension.Signature = &signature
	//invoice.UBLExtensions.Signature = &signature

	//log.Print(invoice.String())

	return nil
}