package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	defer l.Close()

	// var wg sync.WaitGroup

	// for i := 1; i <= 3; i++ {
	// wg.Add(1)
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	// go func() {
	// defer wg.Done()
	req := make([]byte, 1024)

	if _, err := conn.Read(req); err != nil {
		fmt.Println("Failed to read data:", err.Error())
	}

	data := string(req)
	path := strings.Split(data, " ")[1]
	response := ""
	fmt.Println(data)
	if path == "/" {
		response = "HTTP/1.1 200 OK\r\n\r\n"
	} else if strings.HasPrefix(path, "/echo") {
		qry := path[6:]
		response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(qry), qry)
	} else if strings.HasPrefix(path, "/user-agent") {
		agent := strings.Split(data, "\r\n")[3]
		agent = strings.Split(agent, " ")[1]
		response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(agent), agent)
	} else if strings.HasPrefix(path, "/files/") {
		if strings.HasPrefix(data, "GET") {
			text, err := os.ReadFile("./tmp/" + path[6:])
			if err != nil {
				response = "HTTP/1.1 404 NOT FOUND\r\n\r\n"
			} else {
				response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(text), text)
			}
		} else if strings.HasPrefix(data, "POST") {
			body := strings.Split(data, "\r\n\r\n")[1]

			cmd := exec.Command("echo", "-n", body, ">", "./tmp/"+path[6:])
			_, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	} else {
		response = "HTTP/1.1 404 NOT FOUND\r\n\r\n"
	}

	conn.Write([]byte(response))
	// }()
	// }

	// wg.Wait()
}
