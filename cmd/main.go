package main

import (
	"fmt"

	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/inscription"
	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/oneclickmall"
)

func main() {

	oneclickmall.SetEnvironmentIntegration()

	response, err := inscription.DeleteInscription("b1c2abdb-8df8-439b-be7a-53280627e070", "1212")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)

}
