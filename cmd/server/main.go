package main

import (
	"fmt"
	"io"
	"net"

	"github.com/altenhofen/rengechat/pkg/message"
)

var users []string

func main() {

	l, err := net.Listen("tcp", "127.0.0.1:3333")
	if err != nil {
		panic(err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		go handleRequest(conn)
	}

}
func handleRequest(conn net.Conn) {
	buf, err := io.ReadAll(conn)
	if err != nil {
		fmt.Println("failed to read buffer")
		return
	}

	mes, err := message.ParseMessage(string(buf))
	if err != nil {
		fmt.Println("failed to parse message")
		return
	}

	if []byte(*mes.Content)[0] == '/' {
		mesRes := mes.ParseCommands()
		_, err = conn.Write([]byte(mesRes))
		if err != nil {
			fmt.Println("failed to send message to the clients")
			return
		}
		conn.(*net.TCPConn).CloseWrite()
		return
	}

	// check if the user is registered in the current chat

	var mesRes string

	if mes.Action == byte('S') {
		mesRes = fmt.Sprintf("%s says %s [%s]", mes.Sender, *mes.Content, mes.Timestamp.Format("02/01/2006: 15:04:05"))
	}
	_, err = conn.Write([]byte(mesRes))
	if err != nil {
		fmt.Println("failed to send message to the clients")
		return
	}
	conn.(*net.TCPConn).CloseWrite()
}
