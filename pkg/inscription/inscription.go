package inscription

import (
	"fmt"
	"log"

	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/client"
	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/response"
)

const createInscriptionEndpoint = "rswebpaytransaction/api/oneclick/v1.0/inscriptions"
const confirmInscriptionEndpoint = "rswebpaytransaction/api/oneclick/v1.0/inscriptions"

/*CreateInscription Lets get started with the registration process.*/
func CreateInscription(username string, email string, responseURL string) (response.CreateInscriptionResponse, error) {

	body := map[string]interface{}{
		"username":     username,
		"email":        email,
		"response_url": responseURL,
	}

	httpClient := client.GetInstance()

	resp, err := httpClient.Post(createInscriptionEndpoint, body)

	if err != nil {
		log.Printf("Inscription fail in create func.  \n%v\n", err)
	}

	inscriptionCreateResponse, err := response.GetCreateInscriptionResponse(resp.Body())

	return inscriptionCreateResponse, err

}

/*ConfirmInscription  It allows to finish the registration process obtaining the registration status and the tbk_user that is used to make payments.*/
func ConfirmInscription(token string) (response.ConfirmInscriptionResponse, error) {

	httpClient := client.GetInstance()

	endpoint := fmt.Sprintf("%s/%s", confirmInscriptionEndpoint, token)

	resp, err := httpClient.Put(endpoint)

	if err != nil {
		log.Printf("Transaction fail in confirmInscription.  \n%v\n", err)
	}

	confirmInscriptionResponse, err := response.GetConfirmInscriptionResponse(resp.Body())

	return confirmInscriptionResponse, err
}
