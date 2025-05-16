package soap

import (
	"fmt"
	"github.com/aronipurwanto/go-api-northwind/config"
	"github.com/tiaguinho/gosoap"
	"net/http"
)

type SOAPClient struct {
	client *gosoap.Client
}

func NewSOAPClient(cfg config.Config) (*SOAPClient, error) {
	httpClient := &http.Client{}
	client, err := gosoap.SoapClient(cfg.SoapWSDLURL, httpClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create SOAP client: %w", err)
	}
	return &SOAPClient{client: client}, nil
}

func (s *SOAPClient) NumberToWords(number int) (string, error) {
	params := gosoap.Params{
		"ubiNum": number,
	}
	res, err := s.client.Call("NumberToWords", params)
	if err != nil {
		return "", err
	}

	// Ambil hasil dari response
	var result string
	if err := res.Unmarshal(&result); err != nil {
		return "", err
	}
	return result, nil
}
