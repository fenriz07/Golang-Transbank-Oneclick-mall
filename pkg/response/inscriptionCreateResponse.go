package response

import (
	"encoding/json"
	"log"
)

/*CreateInscriptionResponse struct with contain skeleton json to createInscriptionResponse*/
type CreateInscriptionResponse struct {
	Token     string `json:"token"`
	URLWebpay string `json:"url_webpay"`
}

/*GetCreateInscriptionResponse fascade with return struct CreateInscriptionRespons*/
func GetCreateInscriptionResponse(response []byte) (CreateInscriptionResponse, error) {
	var cir CreateInscriptionResponse

	err := json.Unmarshal(response, &cir)

	if err != nil {
		log.Printf("Error in GetCreateInscriptionResponse  %v", err)

		return cir, err
	}

	return cir, nil
}
