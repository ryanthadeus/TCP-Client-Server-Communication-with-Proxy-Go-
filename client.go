package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	menu()
}

func menu() {
	for {
		var choice int

		fmt.Println("1. Input a Message")
		fmt.Println("2. Exit")
		fmt.Scanln(&choice)

		if choice == 1 {
			userInput()
		} else if choice == 2 {
			fmt.Println("Exited")
			break
		} else {
			fmt.Println("Input a either 1 or 2")
		}
	}
}

func userInput() {
	input := bufio.NewScanner(os.Stdin)
	var message string

	for {
		fmt.Print("Input Your Message (Message must be more than 6 characters long and ends with '!'): ")
		input.Scan()
		message = input.Text()

		if len(message) < 7 || message[len(message)-1] != '!' {
			fmt.Println("Message must be more than 6 characters long and ends with '!'")
		} else {
			break
		}
	}

	sendMessage(message)
}

func sendMessage(msg string) {
	dial, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		panic(err)
	}

	binary.Write(dial, binary.BigEndian, uint32(len(msg)))
	_, err = dial.Write([]byte(msg))
	if err != nil {
		panic(err)
	}

	dial.SetReadDeadline(time.Now().Add(3 * time.Second))

	var size uint32
	err = binary.Read(dial, binary.BigEndian, &size)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, size)
	dial.Read([]byte(buffer))
	receive := string(buffer)

	fmt.Println(receive)

	defer dial.Close()
}
