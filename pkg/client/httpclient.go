package client

import (
	"log"

	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/environment"
	"github.com/go-resty/resty/v2"
)

/*HttpClient client*/
type HttpClient struct {
	Environment *environment.Environment
}

/*Post http post for client*/
func (client HttpClient) Post(endpoint string, body map[string]interface{}) (*resty.Response, error) {
	clientResty := client.init()

	url := client.generateURL(endpoint)

	resp, err := clientResty.R().
		SetBody(body).
		Post(url)

	if err != nil {
		log.Printf("Resty fail in post method client :  %v\n", err)
		return resp, err
	}

	err = validateResponse(resp)

	if err != nil {
		return resp, err
	}

	return resp, err
}

/*Put http put method for client*/
func (client HttpClient) Put(endpoint string) (*resty.Response, error) {
	clientResty := client.init()

	url := client.generateURL(endpoint)

	resp, err := clientResty.R().
		Put(url)

	if err != nil {
		log.Printf("Resty fail in post method client :  %v\n", err)
		return resp, err
	}

	err = validateResponse(resp)

	if err != nil {
		return resp, err
	}

	return resp, err
}

/*Get http get method for client*/
func (client HttpClient) Get(endpoint string) (*resty.Response, error) {

	clientResty := client.init()

	url := client.generateURL(endpoint)

	resp, err := clientResty.R().
		Get(url)

	if err != nil {
		log.Printf("Resty fail in get method client :  %v\n", err)
		return resp, err
	}

	err = validateResponse(resp)

	if err != nil {
		return resp, err
	}

	return resp, err

}

/*Delete http delete method for client*/
func (client HttpClient) Delete(endpoint string, body map[string]interface{}) (*resty.Response, error) {
	clientResty := client.init()

	url := client.generateURL(endpoint)

	resp, err := clientResty.R().
		SetBody(body).
		Delete(url)

	if err != nil {
		log.Printf("Resty fail in delete method client :  %v\n", err)
		return resp, err
	}

	err = validateResponse(resp)

	if err != nil {
		return resp, err
	}

	return resp, err
}
