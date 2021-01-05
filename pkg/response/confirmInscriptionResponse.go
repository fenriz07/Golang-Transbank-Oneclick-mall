package response

import (
	"encoding/json"
	"log"
)

/*ConfirmInscriptionResponse struct with contain skeleton json to confirmInscriptionResponse*/
type ConfirmInscriptionResponse struct {
	ResponseCode      int    `json:"response_code"`
	TbkUser           string `json:"tbk_user"`
	AuthorizationCode string `json:"authorization_code"`
	CardType          string `json:"card_type"`
	CardNumber        string `json:"card_number"`
}

/*GetConfirmInscriptionResponse fascade with return struct ConfirmInscriptionResponse*/
func GetConfirmInscriptionResponse(response []byte) (ConfirmInscriptionResponse, error) {
	var cir ConfirmInscriptionResponse

	err := json.Unmarshal(response, &cir)

	if err != nil {
		log.Printf("Error in GetConfirmInscriptionResponse  %v", err)

		return cir, err
	}

	return cir, nil
}
