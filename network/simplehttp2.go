package main

import (
	"net"
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError("net.ResolveTCPAddr()->", err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError("net.DialTCP()->", err)

	_, err = conn.Write([]byte("HEAD HTTP/1.1\r\n\r\n"))
	checkError("conn.Write()->", err)

	result, err := ioutil.ReadAll(conn)
	checkError("ioutil.ReadAll()->", err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(str string, err error) {
	if err != nil {
		fmt.Println(os.Stderr, "%s error %s\n", str, err.Error())
		os.Exit(1)
	}
}
