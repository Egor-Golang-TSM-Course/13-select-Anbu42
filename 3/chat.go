package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type Message struct {
	SenderID int
	Text     string
}

func client(clientID int, messages chan<- Message, quit <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-quit:
			log.Printf("Клиент %d завершает работу\n", clientID)
			return
		default:
			message := Message{
				SenderID: clientID,
				Text:     fmt.Sprintf("Привет, я клиент %d!", clientID),
			}
			messages <- message
			time.Sleep(time.Second * 2)
		}
	}
}

func server(messages <-chan Message, quit <-chan struct{}, wg *sync.WaitGroup, activeClients int) {
	defer wg.Done()
	for {
		select {
		case <-quit:
			log.Println("Сервер завершает работу")
			return
		case message := <-messages:
			for i := 0; i < activeClients; i++ {
				if i != message.SenderID {
					log.Printf("Клиент %d получил сообщение: %s\n", i, message.Text)
					time.Sleep(1)
				}
			}
		}
	}
}

func main() {
	log.SetOutput(os.Stdout)

	messages := make(chan Message)
	quit := make(chan struct{})

	activeClients := 3

	var wg sync.WaitGroup

	wg.Add(1)
	go server(messages, quit, &wg, activeClients)

	for i := 0; i < activeClients; i++ {
		wg.Add(1)
		go client(i, messages, quit, &wg)
	}

	time.Sleep(time.Second * 4)

	close(quit)

	close(messages)

	wg.Wait()
}
