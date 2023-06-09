
package main

import (
	"fmt"
	_ "io"
	"log"
	"net"
)

func main() {
	Server()
}

func Server() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("访问客户端信息: con=%v 客户端: ip=%v\n", conn, conn.RemoteAddr().String())

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	for {
		fmt.Printf("服务器在等待客户端 %s 发送信息\n", c.RemoteAddr().String())
		buf := make([]byte, 1024)
		n, err := c.Read(buf)
		if err != nil {
			log.Fatal(err)
			break
		}

		fmt.Print(string(buf[:n]))
	}
}