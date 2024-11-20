package main

import (
	"fmt"

	"simasware.com.br/email-microservice/api"
)

func main() {
	fmt.Println("⏳ Serviço iniciado na porta 8080...")
	api.StartServer()
}
