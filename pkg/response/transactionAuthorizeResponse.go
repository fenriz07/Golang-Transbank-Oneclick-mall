package response

import (
	"encoding/json"
	"log"
	"time"
)

/*TransactionAuthorizeResponse struct with contain skeleton json to transactionAuthorizeResponse*/
type TransactionAuthorizeResponse struct {
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

/*GetTransactionAuthorizeResponse fascade with return struct  TransactionAuthorizeResponse*/
func GetTransactionAuthorizeResponse(response []byte) (TransactionAuthorizeResponse, error) {
	var tar TransactionAuthorizeResponse

	err := json.Unmarshal(response, &tar)

	if err != nil {
		log.Printf("GetTransactionAuthorizeResponse  %v", err)

		return tar, err
	}

	return tar, nil
}
