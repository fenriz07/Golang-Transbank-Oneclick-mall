package inscription

import (
	"fmt"
	"log"

	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/client"
	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/response"
)

const inscriptionEndpoint = "rswebpaytransaction/api/oneclick/v1.0/inscriptions"

/*CreateInscription Lets get started with the registration process.*/
func CreateInscription(username string, email string, responseURL string) (response.CreateInscriptionResponse, error) {

	body := map[string]interface{}{
		"username":     username,
		"email":        email,
		"response_url": responseURL,
	}

	httpClient := client.GetInstance()

	resp, err := httpClient.Post(inscriptionEndpoint, body)

	if err != nil {
		log.Printf("Inscription fail in create func.  \n%v\n", err)

		return response.CreateInscriptionResponse{}, err
	}

	inscriptionCreateResponse, err := response.GetCreateInscriptionResponse(resp.Body())

	return inscriptionCreateResponse, err

}

/*ConfirmInscription  It allows to finish the registration process obtaining the registration status and the tbk_user that is used to make payments.*/
func ConfirmInscription(token string) (response.ConfirmInscriptionResponse, error) {

	httpClient := client.GetInstance()

	endpoint := fmt.Sprintf("%s/%s", inscriptionEndpoint, token)

	resp, err := httpClient.Put(endpoint)

	if err != nil {
		log.Printf("Confirm inscription fail in confirmInscription.  \n%v\n", err)

		return response.ConfirmInscriptionResponse{}, err
	}

	confirmInscriptionResponse, err := response.GetConfirmInscriptionResponse(resp.Body())

	return confirmInscriptionResponse, err
}

/*DeleteInscription It allows you to delete a user enrolled in Oneclick Mall.
  204 = inscription delete
  404 = inscription not delete
*/
func DeleteInscription(tbkUSER string, username string) (int, error) {
	httpClient := client.GetInstance()

	body := map[string]interface{}{
		"tbk_user": tbkUSER,
		"username": username,
	}

	resp, err := httpClient.Delete(inscriptionEndpoint, body)

	return resp.StatusCode(), err

}
