/*
Utilitário responsável por postar mensagens em uma fila no ActiveMQ para testes
Autor: Fábio Sartori
Copyright: 202401
*/
package main

import (
	"github.com/go-stomp/stomp/v3"
	utils "github.com/kiosanim/simple-activemq-client/util"
	"log"
	"os"
	"time"
)

func sendMessage(connection *stomp.Conn, message string) {
	// Enviando mensagem para a fila
	err := connection.Send(os.Getenv("MQ_QUEUE"), "text/plain", []byte(message))
	if err != nil {
		log.Fatalln("Erro ao postar mensagem na fila:", err)
	}
	log.Printf("Enviando mensagem: %s", message)
}

func main() {
	log.Println("Inicializando o producer")
	utils.LoadEnv()
	// Conectando ao message broker
	conn, err := stomp.Dial("tcp", os.Getenv("MQ_URL"),
		stomp.ConnOpt.Login(os.Getenv("MQ_USERNAME"), os.Getenv("MQ_PASSWORD")))
	if err != nil {
		log.Fatalln("Erro ao conectar com o broker:", err)
	}
	defer conn.Disconnect()

	if len(os.Args) > 1 {
		log.Println(os.Args[1])
		sendMessage(conn, os.Args[1])
	} else {
		for idx := 0; idx < 1000; idx++ {
			message := time.Now().Format("2006-01-02 15:04:05")
			sendMessage(conn, message)
		}
	}
}
