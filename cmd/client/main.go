package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const (
	CONN_HOST = "127.0.0.1"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	// altenhofen|S|Hello World!
	var sb strings.Builder
	for scanner.Scan() {
		sb.WriteString(string(scanner.Bytes()))
	}

	_, err = conn.Write([]byte(sb.String()))
	if err != nil {
		panic(err)
	}
	conn.(*net.TCPConn).CloseWrite()

	buf, err := io.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf))

}
