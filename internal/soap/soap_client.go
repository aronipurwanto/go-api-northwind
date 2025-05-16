package soap

import (
	"encoding/xml"
	"fmt"
	"github.com/aronipurwanto/go-api-northwind/config"
	"github.com/tiaguinho/gosoap"
	"log"
	"net/http"
)

type SOAPClient struct {
	client *gosoap.Client
}

func NewSOAPClient(cfg config.Config) *SOAPClient {
	httpClient := &http.Client{}
	client, err := gosoap.SoapClient(cfg.SoapWSDLURL, httpClient)
	if err != nil {
		log.Fatalf("Failed to create SOAP client: %v", err)
	}
	return &SOAPClient{client: client}
}

func (s *SOAPClient) CallNumberToWords(number string) (string, error) {
	params := gosoap.Params{
		"ubiNum": number,
	}
	resp, err := s.client.Call("NumberToWords", params)
	if err != nil {
		return "", err
	}

	var result struct {
		XMLName             xml.Name `xml:"NumberToWordsResponse"`
		NumberToWordsResult string   `xml:"NumberToWordsResult"`
	}

	if err := resp.Unmarshal(&result); err != nil {
		return "", fmt.Errorf("failed to parse SOAP response: %v", err)
	}

	return result.NumberToWordsResult, nil
}

func (s *SOAPClient) CallNumberToDollars(number string) (string, error) {
	params := gosoap.Params{
		"dNum": number,
	}
	resp, err := s.client.Call("NumberToDollars", params)
	if err != nil {
		return "", err
	}

	var result struct {
		XMLName               xml.Name `xml:"NumberToDollarsResponse"`
		NumberToDollarsResult string   `xml:"NumberToDollarsResult"`
	}

	if err := resp.Unmarshal(&result); err != nil {
		return "", fmt.Errorf("failed to parse SOAP response: %v", err)
	}

	return result.NumberToDollarsResult, nil
}
