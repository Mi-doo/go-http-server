package main

import (
	"fmt"
	"strings"
	"net"
	"os"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	// Status line:
	// HTTP/1.1  // HTTP version
	// 200       // Status code
	// OK        // Optional reason phrase
	// \r\n      // CRLF that marks the end of the status line

	// Headers (empty)
  // \r\n      // CRLF that marks the end of the headers
	
	// Response body (empty)


	req :=make([]byte,1024)

	if _,err :=conn.Read(data) ; err != nil {
		fmt.Println("Failed to read data:",err.Error())
	}
	
	data := string(req)
	idx := strings.Index(data,"/")

	if strings.HasPrefix(string(data), "GET / HTTP/1.1") { 
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	}else{
		conn.Write([]byte("HTTP/1.1 404 NOT FOUND\r\n\r\n"))
	}
}
