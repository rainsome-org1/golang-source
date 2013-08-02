package main

import (
	"net"
	"os"
	"io"
	"bytes"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}
	host := os.Args[1]

	conn, err := net.Dial("ip4:icmp", host)
	checkError("net.Dial()->", err)

	var msg [512]byte
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 0
	msg[7] = 37

	var length int = 8
	check := checkSum(msg[0:length])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[0:length])
	checkError("conn.Write()->", err)

	result, err := readFully(conn)
	checkError("readFully()->", err)

	fmt.Println("Got responsed.")
	fmt.Println(string(result))
	if msg[5] == 13 {
		fmt.Println("The identifier match.")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence match.")
	}

	os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	sum := 0

	for i := 1; i < len(msg) - 1; i += 2 {
		sum += int(msg[i]) * 256 + int(msg[i + 1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)

	return answer
}

func checkError(str string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s error: %s\n", str, err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
