package udp

import (
	"github.com/aiempire79/mr-xray/common/buf"
	"github.com/aiempire79/mr-xray/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
