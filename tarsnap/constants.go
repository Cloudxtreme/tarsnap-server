package tarsnap

// Default server proto and addr
const (
	DefaultProto      = "tcp"
	DefaultListenAddr = ":9279"
)

// Defaults for validating input
const (
	// Max length of proto version
	MaxUserAgentLen = 255

	// Max user length
	MaxUserLen = 255
)

// Packet headers
const (
	// A new machine wants to register with the server.
	NetPacketRegisterRequest = 0x00

	// As part of the machine registration process, the server needs to verify that the machine has the user password.
	NetPacketRegisterChallenge = 0x80
)
