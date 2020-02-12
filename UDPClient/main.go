package main

import (
	"bufio"
	"fmt"
	"net"
)

const (
	HOST = "192.168.100.240"
	PORT = "41234"
)

func GetRespose(conn *net.Conn) (resp string, err error) {
	buffer := make([]byte, 8096)
	_, err = bufio.NewReader(*conn).Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

func main() {
	conn, err := net.Dial("udp4",  fmt.Sprintf("%s:%s", HOST, PORT))
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
	}

}
