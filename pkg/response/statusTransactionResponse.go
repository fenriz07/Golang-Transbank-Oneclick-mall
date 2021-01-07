package response

import (
	"encoding/json"
	"log"
	"time"
)

type StatusTransactionResponse struct {
	BuyOrder   string `json:"buy_order"`
	CardDetail struct {
		CardNumber string `json:"card_number"`
	} `json:"card_detail"`
	AccountingDate  string    `json:"accounting_date"`
	TransactionDate time.Time `json:"transaction_date"`
	Details         []struct {
		Amount             int    `json:"amount"`
		Status             string `json:"status"`
		AuthorizationCode  string `json:"authorization_code"`
		PaymentTypeCode    string `json:"payment_type_code"`
		ResponseCode       int    `json:"response_code"`
		InstallmentsNumber int    `json:"installments_number"`
		CommerceCode       string `json:"commerce_code"`
		BuyOrder           string `json:"buy_order"`
	} `json:"details"`
}

/*GetStatusTransactionResponse fascade with return struct CreateInscriptionRespons*/
func GetStatusTransactionResponse(response []byte) (StatusTransactionResponse, error) {
	var str StatusTransactionResponse

	err := json.Unmarshal(response, &str)

	if err != nil {
		log.Printf("Error in GetStatusTransactionResponse  %v", err)

		return str, err
	}

	return str, nil
}
