package main

import (
	"fmt"
	"sync"
	"time"
)

func senderGoroutine(name string, message string, outChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s: Sending the message: \"%s\"\n", name, message)
	time.Sleep(500 * time.Millisecond)
	outChan <- message
}

func receiverGoroutine(name string, inChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s: Waiting for the message...\n", name)
	msg := <-inChan
	fmt.Printf("%s: Message received: \"%s\"\n", name, msg)
}

func sender(out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- "Hello"
}

func receiver(in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-in
	fmt.Println("Additional receiver got:", msg)
}

func main() {
	fmt.Println("Beginning of messaging...")

	messageChannel := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)
	go receiverGoroutine("Receiver 1", messageChannel, &wg)

	wg.Add(1)
	go senderGoroutine("Sender 1", "Hello! Where are you?", messageChannel, &wg)

	additionalChannel := make(chan string, 1)
	wg.Add(2)
	go sender(additionalChannel, &wg)
	go receiver(additionalChannel, &wg)

	wg.Wait()

	close(messageChannel)
	close(additionalChannel)
	fmt.Println("All channels closed.")

	fmt.Println("Messaging finished.")
}
