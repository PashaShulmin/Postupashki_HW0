package client

import (
	"fmt"
	"net"
	"os"
)

func Connect() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		os.Exit(1)
	}
	defer conn.Close()

	buffer := make([]byte, 2)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка при чтении данных от сервера:", err)
		return
	}

	if string(buffer) == "OK" {
		fmt.Println("Сервер отправил: OK")
	} else {
		fmt.Println("Неверный ответ от сервера:", string(buffer))
	}
}
