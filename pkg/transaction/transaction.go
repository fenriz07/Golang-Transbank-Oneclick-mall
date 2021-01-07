package transaction

import (
	"fmt"
	"log"

	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/client"
	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/response"
)

const transactionEndpoint = "rswebpaytransaction/api/oneclick/v1.0/transactions"
const transactonRefundEndpoint = "/rswebpaytransaction/api/oneclick/v1.0/transactions/%s/refunds"

/*AuthorizeTransaction Allows you to authorize a payment.*/
func AuthorizeTransaction(username string, tbkUSER string, parentBuyOrder string, details detailsTransaction) (response.TransactionAuthorizeResponse, error) {

	detailBody := make(map[string]interface{}, 4)

	detailBody["commerce_code"] = details.CommerceCode
	detailBody["buy_order"] = details.BuyOrder
	detailBody["amount"] = details.Amount
	detailBody["installments_number"] = details.InstallmentsNumber

	var detailsBody [1]map[string]interface{}

	detailsBody[0] = detailBody

	body := map[string]interface{}{
		"username":  username,
		"tbk_user":  tbkUSER,
		"buy_order": parentBuyOrder,
		"details":   detailsBody,
	}

	httpClient := client.GetInstance()

	resp, err := httpClient.Post(transactionEndpoint, body)

	if err != nil {
		log.Printf("Transaction fail in authorize func.  \n%v\n", err)

		return response.TransactionAuthorizeResponse{}, err
	}

	transactionAuthorizeResponse, err := response.GetTransactionAuthorizeResponse(resp.Body())

	return transactionAuthorizeResponse, err
}

/*StatusTransaction This allows the operation to obtain the status of the transaction at any time*/
func StatusTransaction(buyOrder string) (response.StatusTransactionResponse, error) {

	httpClient := client.GetInstance()

	endpoint := fmt.Sprintf("%s/%s", transactionEndpoint, buyOrder)

	resp, err := httpClient.Get(endpoint)

	if err != nil {
		log.Printf("Status Transaction fail in StatusTransaction func.  \n%v\n", err)

		return response.StatusTransactionResponse{}, err
	}

	refundTransactionResponse, err := response.GetStatusTransactionResponse(resp.Body())

	return refundTransactionResponse, err
}

/*RefundTransaction Allows you to reverse or void a previously authorized sales
transaction. This method returns a unique identifier of the reverse / abort transaction
as a response.*/
func RefundTransaction(buyOrder string, commerceCode string, detailBuyOrder string, amount int) (response.RefundTransactionResponse, error) {

	httpClient := client.GetInstance()

	endpoint := fmt.Sprintf(transactonRefundEndpoint, buyOrder)

	body := map[string]interface{}{
		"commerce_code":    commerceCode,
		"detail_buy_order": detailBuyOrder,
		"amount":           amount,
	}

	resp, err := httpClient.Post(endpoint, body)

	if err != nil {
		log.Printf("Refund Transaction fail in RefundTransaction func.  \n%v\n", err)

		return response.RefundTransactionResponse{}, err
	}

	refundTransactionResponse, err := response.GetRefundTransactionResponse(resp.Body())

	return refundTransactionResponse, err

}
