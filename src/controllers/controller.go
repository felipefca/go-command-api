package controllers

import (
	"api/src/models"
	"api/src/notification"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// SaveStock          godoc
// @Summary 		  Salva stock para processamento
// @Description 	  Salva a stock no MongoDB e envia para a fila de processamento
// @Tags              tags
// @Success 		  200 {object} NomeDoObjetoDeRetorno
// @Failure 		  400 {object} Erro
// @Router			  /tags [post]
func SaveStock(w http.ResponseWriter, r *http.Request) {

	request, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var stock models.Stock

	if erro = json.Unmarshal(request, &stock); erro != nil {
		log.Fatal(erro)
	}

	stock.CreationDate = time.Now()

	str, erro := repository.AddStock(stock)
	if erro != nil {
		log.Fatal(erro)
	}

	jsonBody, err := json.Marshal(stock)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	err = notification.SendStockMessage(jsonBody)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	w.Write([]byte(fmt.Sprintf("Stock inserida: %s", str)))
}

// SaveStock          godoc
// @Summary 		  Salva moeda para processamento
// @Description 	  Salva a moeda no MongoDB e envia para a fila de processamento
// @Tags              tags
// @Success 		  200 {object} NomeDoObjetoDeRetorno
// @Failure 		  400 {object} Erro
// @Router			  /tags [post]
func SaveFiat(w http.ResponseWriter, r *http.Request) {

	request, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var fiat models.Fiat

	if erro = json.Unmarshal(request, &fiat); erro != nil {
		log.Fatal(erro)
	}

	fiat.CreationDate = time.Now()

	response, err := repository.GetFiat(fiat.Code)
	if err != nil {
		log.Fatal(err)
	}

	if response.Code != "" {
		w.Write([]byte("Já existe moeda inserida com esses valores"))
		return
	}

	str, erro := repository.AddFiat(fiat)
	if erro != nil {
		log.Fatal(erro)
	}

	jsonBody, err := json.Marshal(fiat)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	err = notification.SendFiatMessage(jsonBody)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	w.Write([]byte(fmt.Sprintf("Moeda inserida: %s", str)))
}
