package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	// parts of the packet
	var seq int      // sequence of the packet being sent
	var size int     // how long the message is in characters (no more than 64)
	var msg string   // the actual contents
	var dType string // two types, data and eof.

	raddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:7030") // TODO: remember to softcode addr
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	send_str := chunkStrings("Hi everybody, what's up? This should be a slightly smaller message, to test.")

	for i := range send_str {
		seq = i
		size = len(send_str[i])
		msg = send_str[i]
		if send_str[i] == "" {
			dType = "eof"
		} else {
			dType = "data"
		}

		pckt := []string{strconv.Itoa(seq), strconv.Itoa(size), msg, dType}
		fmt.Println(strings.Join(pckt, ","))
		send(strings.Join(pckt, ","), conn)
	}

	// read from server
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))
}

func send(s string, conn net.Conn) {
	if _, err := conn.Write([]byte(s)); err != nil {
		log.Fatal(err)
	}
}

// figure out how to send each chunk.

func chunkStrings(s string) (brokeStrings []string) {
	var rang int
	var count int

	for i := range s {
		if i != 0 && (i%64 == 0 || i == len(s)-1) {
			brokeStrings = append(brokeStrings, s[rang:i])
			rang = i
			count++
		}
	}
	// so we know where to put the eof signifier
	brokeStrings = append(brokeStrings, "")

	fmt.Println("Strings chunked!")

	return
}
