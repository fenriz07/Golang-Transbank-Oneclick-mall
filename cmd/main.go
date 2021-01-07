package main

import (
	"fmt"

	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/oneclickmall"
	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/transaction"
)

func main() {

	oneclickmall.SetEnvironmentIntegration()

	detail := transaction.CreateDetailTransaction("597055555542", "7oneclickgolang", 4000, 1)

	response, err := transaction.AuthorizeTransaction("1212", "6130952c-52db-4aed-b791-530a0fd1f94d", "7oneclickgolang", detail)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)

	statusTransaction, err := transaction.StatusTransaction("7oneclickgolang")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(statusTransaction)

	refundTransaction, err := transaction.RefundTransaction("7oneclickgolang", "597055555542", "7oneclickgolang", 2000)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(refundTransaction)

}
