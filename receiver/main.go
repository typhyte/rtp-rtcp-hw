package main

import (
	"fmt"
	"log"
	//"math/rand"
	"net"
	"strings"
	"time"
)

func main() {
	fmt.Println("starting server...")

	pc, err := net.ListenPacket("udp", ":7030")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	fmt.Println("initialized!")

	fmt.Println("waiting for client...")

	for {
		buf := make([]byte, 2048)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			log.Println("read:", err)
			continue
		}
		go respond(pc, addr, buf[:n])
		fmt.Println("message sent to client!")
	}
}

func respond(pc net.PacketConn, addr net.Addr, payload []byte) {
	// vars
	// receiving info
	ts := time.Now().Format(time.ANSIC)                                   // timestamp
	msg := strings.SplitN(string(payload), ",", 4)                        // base message
	send := fmt.Sprintf("time recieved: %s. The message is %s!", ts, msg) // message 2 be sent

	// sending info
	// recvId := rand.Intn(256)
	// var recvPackets int
	// var lostPackets int

	// debug prints
	fmt.Println(strings.SplitN(string(payload), ",", 4))

	// logic
	if _, err := pc.WriteTo([]byte(send), addr); err != nil {
		log.Println("write: ", err)
	}
}
