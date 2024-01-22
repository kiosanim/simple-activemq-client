/*
Utilitário responsável por consumir mensagens de uma fila no ActiveMQ para testes
Autor: Fábio Sartori
Copyright: 202401
*/
package main

import (
	"fmt"
	"github.com/go-stomp/stomp/v3"
	utils "github.com/kiosanim/simple-activemq-client/util"
	"log"
	"os"
	"time"
)

func timeoutSecs() time.Duration {
	timeout, err := time.ParseDuration(os.Getenv("MQ_TIMEOUT_SECS") + "s")
	if err != nil {
		log.Fatalln("Erro ao converter timeout:", err)
	}
	return timeout
}

func reconnectTimeoutSecs() time.Duration {
	timeout, err := time.ParseDuration(os.Getenv("MQ_RECONNECT_TIMEOUT_SECS") + "s")
	if err != nil {
		log.Fatalln("Erro ao converter timeout:", err)
	}
	return timeout
}

func connectToActiveMQ() (*stomp.Conn, error) {
	// Connectando ao STOMP message broker
	return stomp.Dial("tcp", os.Getenv("MQ_URL"),
		stomp.ConnOpt.Login(os.Getenv("MQ_USERNAME"), os.Getenv("MQ_PASSWORD")),
		stomp.ConnOpt.HeartBeat(timeoutSecs(), timeoutSecs()))
}

func processMessage(message *stomp.Message) {
	log.Println("Message received:", string(message.Body))
}

func main() {
	var connection *stomp.Conn
	var err error
	log.Println("Inicializando o consumer")
	utils.LoadEnv()
	// Connectando ao STOMP message broker
	connection, err = connectToActiveMQ()
	if err != nil {
		log.Fatalln("Erro ao conectar com o broker:", err)
	}
	defer connection.Disconnect()
	// Subscrevendo a fila
	subscription, err := connection.Subscribe(os.Getenv("MQ_QUEUE"), stomp.AckAuto)
	if err != nil {
		log.Fatalln("Erro ao subscrever a fila:", err)
	}
	defer subscription.Unsubscribe()
	fmt.Println("Aguardando mensagens. Pressione <CRTL> + C, para sair")

	//MainLoop:
	for {
		select {
		case message := <-subscription.C:
			if message.Err != nil {
				log.Println("Erro ao receber mensagem:", message.Err)
				continue
			}
			processMessage(message)
		case <-time.After(timeoutSecs()):
			log.Println("Timeout. Reconectando...")
			connection, err = connectToActiveMQ()
			if err != nil {
				log.Println("Erro ao reconectar no ActiveMQ:", err)
				time.Sleep(reconnectTimeoutSecs())
			}
			subscription, err = connection.Subscribe(os.Getenv("MQ_QUEUE"), stomp.AckAuto)
			if err != nil {
				log.Fatalln("Erro ao subscrever a fila:", err)
				time.Sleep(reconnectTimeoutSecs())
				continue
			}
			log.Println("Reconectado com sucesso !!!")
		}
	}
	log.Println("Encerrando consumer")
	log.Println(time.DateTime)
}
