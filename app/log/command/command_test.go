package command_test

import (
	"context"
	"testing"

	"github.com/aiempire79/mr-xray/app/dispatcher"
	"github.com/aiempire79/mr-xray/app/log"
	. "github.com/aiempire79/mr-xray/app/log/command"
	"github.com/aiempire79/mr-xray/app/proxyman"
	_ "github.com/aiempire79/mr-xray/app/proxyman/inbound"
	_ "github.com/aiempire79/mr-xray/app/proxyman/outbound"
	"github.com/aiempire79/mr-xray/common"
	"github.com/aiempire79/mr-xray/common/serial"
	"github.com/aiempire79/mr-xray/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
