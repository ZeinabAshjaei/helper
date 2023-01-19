package nameProducer

import "fmt"

func GetGreeting(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func GetHelloWorld() {
	fmt.Print("Hello World!")
}
