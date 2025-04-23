package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	req := make([]byte, 1024)

	if _, err := conn.Read(req); err != nil {
		fmt.Println("Failed to read data:", err.Error())
	}

	data := string(req)
	path := strings.Split(data, " ")[1]
	response := ""

	if path == "/" {
		response = "HTTP/1.1 200 OK\r\n\r\n"
	} else if strings.HasPrefix(path, "/echo") {
		qry := path[6:]
		response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(qry), qry)
	} else if strings.HasPrefix(path, "/user-agent") {
		agent := strings.Split(data, "\r\n")[3]
		agent = strings.Split(agent, " ")[1]
		response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(agent), agent)
	} else {
		response = "HTTP/1.1 404 NOT FOUND\r\n\r\n"
	}

	conn.Write([]byte(response))
}
