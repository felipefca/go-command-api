package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// @title Command API
// @version 1.0
// @description API para processamento de comandos de moedas e stocks
// @host localhost:5000
// @BasePath /api/v1
// @schemes http https
// @produces json
// @consumes json
// @contact.name Felipe Fernandes
// @contact.email felipe.fca1987@gmail.com
// @license.name MIT
// @license.url http://opensource.org/licenses/MIT
func main() {
	config.Load()

	fmt.Println(config.Port)

	fmt.Println("Rodando API")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
