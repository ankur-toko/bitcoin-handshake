package message

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

/*
Field Size	Description	Data type	Comments
4	version	int32_t	Identifies protocol version being used by the node
8	services	uint64_t	bitfield of features to be enabled for this connection
8	timestamp	int64_t	standard UNIX timestamp in seconds
26	addr_recv	net_addr	The network address of the node receiving this message
Fields below require version ≥ 106
26	addr_from	net_addr	Field can be ignored. This used to be the network address of the node emitting this message, but most P2P implementations send 26 dummy bytes. The "services" field of the address would also be redundant with the second field of the version message.
8	nonce	uint64_t	Node random nonce, randomly generated every time a version packet is sent. This nonce is used to detect connections to self.

	?	user_agent	var_str	User Agent (0x00 if string is 0 bytes long)

4	start_height	int32_t	The last block received by the emitting node
Fields below require version ≥ 70001
1	relay	bool	Whether the remote peer should announce relayed transactions or not, see BIP 0037
*/

type NetAddress struct {
	Timestamp time.Time

	// Bitfield which identifies the services supported by the address.
	Services uint64

	// IP address of the peer.
	IP   net.IP
	Port uint16
}

/*
Field Size	Description	Data type	Comments
4	time	uint32	the Time (version >= 31402). Not present in version message.
8	services	uint64_t	same service(s) listed in version
16	IPv6/4	char[16]	IPv6 address. Network byte order. The original client only supported IPv4 and only read the last 4 bytes to get the IPv4 address. However, the IPv4 address is written into the message as a 16 byte IPv4-mapped IPv6 address
(12 bytes 00 00 00 00 00 00 00 00 00 00 FF FF, followed by the 4 bytes of the IPv4 address).

2	port	uint16_t	port number, network byte order
*/

// todo : Add all possible validations as well
func (a NetAddress) IsValid() bool {
	if len(a.IP) > 0 && a.Port > 0 {
		return true
	}
	return false
}

type VersionMsg struct {
	Version     uint32
	Timestamp   uint64
	Services    uint64
	RecAddr     NetAddress
	SendAddr    NetAddress
	Nounce      uint64
	UserAgent   string
	StartHeight uint32
	Relay       bool
}

// Converts the message if everything if all required fields are available
// Validates the message as well
// Ensures architecture validations as well
func (v VersionMsg) ToBytes() ([]byte, error) {
	if isValid, err := v.IsValid(); !isValid {
		return nil, err
	}
	payload := new(bytes.Buffer)
	binary.Write(payload, binary.LittleEndian, int32(v.Version))
	binary.Write(payload, binary.LittleEndian, uint64(v.Services))
	binary.Write(payload, binary.LittleEndian, int64(v.Timestamp))
	CopyAddressToBuffer(payload, v.RecAddr)
	if len(v.SendAddr.IP) > 0 {
		CopyAddressToBuffer(payload, v.SendAddr)
	}
	if v.Nounce == 0 {
		// 0 will generate a new random number, otherwise send random in the struct itself.
		nonce := make([]byte, 8)
		binary.LittleEndian.PutUint64(nonce, uint64(time.Now().UnixNano()))
		payload.Write(nonce)
	}
	userAgent := v.UserAgent
	binary.Write(payload, binary.LittleEndian, uint8(len(userAgent))) // Length of the user agent string
	payload.WriteString(userAgent)                                    // User agent string

	// Last four bytes, specifying the height of the sender's blockchain
	binary.Write(payload, binary.LittleEndian, int32(v.StartHeight)) // Starting height
	binary.Write(payload, binary.LittleEndian, v.Relay)
	return payload.Bytes(), nil
}

func CopyAddressToBuffer(payload *bytes.Buffer, addr NetAddress) {
	binary.Write(payload, binary.LittleEndian, uint64(addr.Services)) // Services
	payload.Write(make([]byte, 12))
	payload.Write(net.ParseIP("0.0.0.0").To4())
	binary.Write(payload, binary.BigEndian, uint16(8333)) // Port 8333 (standard Bitcoin port)
}

// Version message structure can be different depending upon the version of the message
// Add all complex validations here
func (v VersionMsg) IsValid() (bool, error) {
	if v.Version < 70015 {
		return false, fmt.Errorf("version must be greater than 70015")
	}
	if v.Timestamp <= 0 {
		return false, fmt.Errorf("timestamp must not be nil")
	}
	if !v.RecAddr.IsValid() {
		return false, fmt.Errorf("reciever address must not be nil")
	}
	return true, nil
}
