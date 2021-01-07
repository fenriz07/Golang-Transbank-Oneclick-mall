package response

import (
	"encoding/json"
	"log"
	"time"
)

type RefundTransactionResponse struct {
	Type              string    `json:"type"`
	AuthorizationCode string    `json:"authorization_code"`
	AuthorizationDate time.Time `json:"authorization_date"`
	NullifiedAmount   float64   `json:"nullified_amount"`
	Balance           float64   `json:"balance"`
	ResponseCode      int       `json:"response_code"`
}

/*GetRefundTransactionResponse fascade with return struct CreateInscriptionRespons*/
func GetRefundTransactionResponse(response []byte) (RefundTransactionResponse, error) {
	var rtr RefundTransactionResponse

	err := json.Unmarshal(response, &rtr)

	if err != nil {
		log.Printf("Error in GetRefundTransactionResponse  %v", err)

		return rtr, err
	}

	return rtr, nil
}
