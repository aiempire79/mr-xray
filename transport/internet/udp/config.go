package udp

import (
	"github.com/aiempire79/mr-xray/common"
	"github.com/aiempire79/mr-xray/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
