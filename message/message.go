package message

import (
	"bytes"
	"log"
	"net"
	"time"
)

// Builds the version message
func BuildVersion(recIP string) ([]byte, error) {
	t := uint64(time.Now().Unix())
	payload := VersionMsg{
		Version:     ProtocolVersion,
		Timestamp:   t,
		Services:    uint64(0),
		RecAddr:     NetAddress{time.Now(), 0, []byte(recIP), 8333},
		SendAddr:    NetAddress{time.Now(), 0, []byte(getLocalIP()), 8333},
		Nounce:      0,
		UserAgent:   "AnkurHandshakerV1.0",
		StartHeight: uint32(0),
		Relay:       false,
	}
	pBytes, err := payload.ToBytes()
	if err != nil {
		return nil, err
	}
	header := BuildHeaderFor(CmdVersion, pBytes)

	hBytes, err := header.ToBytes()
	if err != nil {
		return nil, err
	}
	msg := new(bytes.Buffer)
	msg.Write(hBytes)
	msg.Write(pBytes)

	return msg.Bytes(), nil
}

// Builds the VerAck message
func BuildVerAck(recIP string) ([]byte, error) {
	payload := VerAck{}
	pBytes, err := payload.ToBytes()
	if err != nil {
		return nil, err
	}
	header := BuildHeaderFor(CmdVerAck, pBytes)

	hBytes, err := header.ToBytes()
	if err != nil {
		return nil, err
	}
	msg := new(bytes.Buffer)
	msg.Write(hBytes)
	msg.Write(pBytes)

	return msg.Bytes(), nil
}

func getLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}
