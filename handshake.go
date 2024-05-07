package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ankur-toko/bitcoin-handshake/message"
)

func main() {
	urls := []string{"127.0.0.1:8333"}

	for i := 0; i < len(urls); i++ {
		go handshake(urls[i])
	}

	// preventing process exit
	time.Sleep(1 * time.Hour)
}

/*
main test function
1. Sends out a version message
2. Gets the version msg and verack message in response
3. Sends out the verack in response
4. Keeps waiting for any messages from the target node, but ignores them.
*/
func handshake(url string) {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		fmt.Println("error connecting to bitcoin node:", err)
		return
	}
	defer conn.Close()

	// Build version message payload
	msg, err := message.BuildVersion(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	if e := write(conn, msg); e != nil {
		return
	}
	// Waiting for response from the bitcoin node
	if e := read(conn); e != nil {
		return
	}

	verackPayload, err := message.BuildVerAck(url)
	if err != nil {
		fmt.Println("error building verack message:", err)
		return
	}

	if e := write(conn, verackPayload); e != nil {
		fmt.Print(e)
		return
	}

	for {
		if e := read(conn); e != nil {
			fmt.Println("connection lost with :", url)
			return
		}
	}

}

func read(conn net.Conn) error {
	response := make([]byte, 200)
	_, err := conn.Read(response)
	parseResponse(response)
	if err != nil {
		fmt.Println("error reading response:", err)
		return err
	} else {
		fmt.Printf("response from bitcoin node: %v\n", string(response))
		return nil
	}
}

func parseResponse(header []byte) {
	command := string(bytes.TrimRight(header[4:16], string(0)))

	// Check if it's a version message
	if command == "version" {
		fmt.Println("Received version message")
		payload := header[24:]
		version := binary.LittleEndian.Uint32(payload[:4])
		log.Println("bitcoin node version recieved:", version)
	} else if command == "verack" {
		// Process the verack message
		fmt.Println("Received verack message")
	} else if command == "ping" {
		fmt.Println("Received ping message")
	} else {
		// Handle unknown message types
		fmt.Println("Received unknown message:", command)
	}
}

func write(conn net.Conn, msg []byte) error {
	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println("error sending message:", err)
		return err
	}
	return nil
}
