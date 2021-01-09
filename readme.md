![WebPay-One-Click-Mall](https://user-images.githubusercontent.com/9199380/104020581-33c18280-519c-11eb-9a4d-2a3a19b25fdd.png)

<p align="center">
<h1 align="center">Golang - Webpay OneClick Mall</h1>
<p align="center">Biblioteca para la integración de Webpay OneClick Mall en golang </p>

</p>

<p align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/fenriz07/Golang-Transbank-Oneclick-mall)](https://goreportcard.com/report/github.com/fenriz07/Golang-Transbank-Oneclick-mall)
<a href="https://pkg.go.dev/github.com/fenriz07/Golang-Transbank-Oneclick-mall"><img src="https://godoc.org/github.com/fenriz07/Golang-Transbank-Oneclick-mall?status.svg" alt="GoDoc"></a>


</p>

## Características


- Soporte para ambiente de integración y producción
- Crear una inscripción
- Confirmar una inscripción
- Eliminar una inscripción
- Autorizar un pago
- Obtener estado de una transacción
- Reversar o anular una transacción
- Consolidación de respuestas en structs.
- Manejo de errores http.

## Ejemplos de uso:

- https://github.com/fenriz07/oneclick-golang-example todas las caracteristicas implementadas bajo Iris Framework y Docker

## Instalación

```bash
go get github.com/fenriz07/Golang-Transbank-Oneclick-mall
```

## Uso

### Inicializar ambiente

Hay 2 ambientes Integración y producción

```go
//Importar el package webpayplus
import (
	"github.com/fenriz07/Golang-Transbank-WebPay-Rest/pkg/webpayplus"
)
```

#### Integración

```go
/*
    Usar la función SetEnvironmentIntegration para el ambiente de desarrollo.
  Automaticamente configura las credenciales del comercio.
  Configura el cliente y todas las transacciones seran ejecutadas bajo este ambiente automaticamente

*/
oneclickmall.SetEnvironmentIntegration()
```

#### Producción

```go
/*
      Usar la función SetEnvironmentProduction para el ambiente de desarrollo.
  Automaticamente configura las credenciales del comercio.
  Configura el cliente y todas las transacciones seran ejecutadas bajo este ambiente automaticamente

  Dicha función acepta 2 parametros
  1 - APIKeyID     (string)
  2 - APIKeySecret (string)
*/

APIKeyID := "Código de comercio"
APIKeySecret := "Llave secreta"

oneclickmall.SetEnvironmentProduction(APIKeyID,APIKeySecret)
```

### Crear una inscripción
```go
response, err := inscription.CreateInscription(username, email, "http://localhost:8080/inscription/confirm")
```

### Confirmar una inscripción
```go
response, err := inscription.ConfirmInscription(token)
```

### Borrar una inscripción
```go
status, err := inscription.DeleteInscription(userToken, username)
```

### Autorizar un pago
```go
//Primero creamos el detalle de la transacción
detail := transaction.CreateDetailTransaction("597055555542", order, amount, 1)

//Llamamos a transaction.AuthorizeTransaction y pasamos: Usuario, token de usuario, numero de orden y detalle (detail)
response, err := transaction.AuthorizeTransaction(username, userToken, order, detail)
```

### Estado de un pago 
```go
response, err := transaction.StatusTransaction(order)
```


### Reversar o anular un pago
```go
response, err := transaction.RefundTransaction(order, "597055555542", order, amount)
```

## Otras bibliotecas para los servicios de transbank:
- [WebPay Plus](https://github.com/fenriz07/Golang-Transbank-WebPay-Rest)

## Creador

[Fenriz07](https://github.com/fenriz07)
