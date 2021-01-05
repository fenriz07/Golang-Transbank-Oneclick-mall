package oneclickmall

import (
	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/client"
	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/environment"
)

/*SetEnvironmentIntegration fascade for instance integration evironment*/
func SetEnvironmentIntegration() {
	environmentIntegration := environment.IntegrationEnviroment{}

	setInstanceAndClient(environmentIntegration)
}

/*SetEnvironmentProduction fascade for instance production evironment*/
func SetEnvironmentProduction(APIKeyID string, APIKeySecret string) {

	environmentProduction := environment.ProductionEnviroment{}
	environmentProduction.SetCredentials(APIKeyID, APIKeySecret)

	setInstanceAndClient(environmentProduction)

}

func setInstanceAndClient(env environment.TypeEnvironment) {

	environment.SetInstance(env)
	client.SetInstance("http")
}
