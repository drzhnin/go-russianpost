package russianpost

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://tracking.russianpost.ru/rtm34?wsdl%s"
)

//Envelope struct contain request params
type Envelope struct {
	XMLName     xml.Name `xml:"soap:Envelope"`
	SoapAttr    string   `xml:"xmlns:soap,attr"`
	OperAttr    string   `xml:"xmlns:oper,attr"`
	DataAttr    string   `xml:"xmlns:data,attr"`
	SoapEnvAttr string   `xml:"xmlns:soapenv,attr"`
	Header      string   `xml:"soap:Header"`
	Body        EnvelopeBody
}

type EnvelopeBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Oper    EnvelopeOper
}

type EnvelopeOper struct {
	XMLName xml.Name `xml:"oper:getOperationHistory"`
	Data    OperationHistoryRequest
	Auth    AuthorizationHeader
}

type OperationHistoryRequest struct {
	XMLName     xml.Name `xml:"data:OperationHistoryRequest"`
	Barcode     string   `xml:"data:Barcode"`
	MessageType string   `xml:"data:MessageType"`
	Language    string   `xml:"data:Language"`
}
type AuthorizationHeader struct {
	XMLName        xml.Name `xml:"data:AuthorizationHeader"`
	MustUnderstand string   `xml:"soapenv:mustUnderstand,attr"`
	Login          string   `xml:"data:login"`
	Password       string   `xml:"data:password"`
}

type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	baseURL *url.URL

	//Login to access the API Service Tracking
	login string
	//Password to access the API Service Tracking
	password string
}

// NewClient returns a new API client.
// If a nil httpClient is provided, http.DefaultClient will be used.
func NewClient(serviceLogin, servicePassword string) *Client {
	httpClient := http.DefaultClient

	bURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, baseURL: bURL, login: serviceLogin, password: servicePassword}
	return c
}

func (c *Client) GetOperationHistory(barcode, messegeType, language string) {
	operHistReq := OperationHistoryRequest{Barcode: barcode, MessageType: messegeType, Language: language}
	authHeader := AuthorizationHeader{MustUnderstand: "1", Login: c.login, Password: c.password}
	envelopeOper := EnvelopeOper{Data: operHistReq, Auth: authHeader}
	envelopeBody := EnvelopeBody{Oper: envelopeOper}
	d := Envelope{
		SoapAttr:    "http://www.w3.org/2003/05/soap-envelope",
		OperAttr:    "http://russianpost.org/operationhistory",
		DataAttr:    "http://russianpost.org/operationhistory/data",
		SoapEnvAttr: "http://schemas.xmlsoap.org/soap/envelope/",
		Body:        envelopeBody,
	}
	b, err := xml.MarshalIndent(d, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
