package main

import (
	"flag"
	"fmt"
	"net/http"
	"encoding/json"	
	"bytes"
)

type Message struct {
	Username string `json:"username"`
    Content string `json:"content"`
}

func main() {
	var api string
	flag.StringVar(&api, "url", "",  "URL do webhook do Discord do servidor para o qual enviar mensagem")
	
	var message string
	flag.StringVar(&message, "mensagem", "", "Mensagem para enviar")
	
	var user string
	flag.StringVar(&user, "user", "", "Nome do Webhook")
	
	flag.Parse()

	myMsg := Message{Username:user,Content:message}
	b, err := json.Marshal(myMsg)
	
	if err != nil {
		fmt.Println("Erro:", err)	
	}
	enviarMensagem(api, string(b))
}

func enviarMensagem(api string, jsonData string) {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("POST", api, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}