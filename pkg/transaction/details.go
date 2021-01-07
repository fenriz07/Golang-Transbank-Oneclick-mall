package transaction

type detailsTransaction struct {
	CommerceCode       string `json:"commerce_code"`
	BuyOrder           string `json:"buy_order"`
	Amount             int    `json:"amount"`
	InstallmentsNumber int    `json:"installments_number"`
}

/*CreateDetailTransaction return details struct*/
func CreateDetailTransaction(commerceCode string, buyOrder string, amount int, installmentsNumber int) detailsTransaction {

	return detailsTransaction{
		CommerceCode:       commerceCode,
		BuyOrder:           buyOrder,
		Amount:             amount,
		InstallmentsNumber: installmentsNumber,
	}
}
