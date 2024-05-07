package message

const (
	ProtocolVersion uint32 = 70015
	MagicWord       uint32 = 0xD9B4BEF9
)

type Command string

const (
	CmdVersion Command = "version"
	CmdVerAck  Command = "verack"
)
