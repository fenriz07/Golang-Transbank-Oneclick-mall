package main

import (
	"fmt"

	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/inscription"
	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/oneclickmall"
)

func main() {

	oneclickmall.SetEnvironmentIntegration()

	response, _ := inscription.CreateInscription("1212", "servio.za@gmail.com", "https://www.transbankdevelopers.cl")

	fmt.Println(response)

}
