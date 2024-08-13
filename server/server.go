package server

import (
	"fmt"
	"net"
)

func StartServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при создании сокета: ", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии соединения: ", err)
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte("OK"))
	if err != nil {
		fmt.Println("Ошибка при отправке ответа: ", err)
	}
}
