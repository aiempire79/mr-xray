package udp

import (
	"context"

	"github.com/aiempire79/mr-xray/common"
	"github.com/aiempire79/mr-xray/common/net"
	"github.com/aiempire79/mr-xray/transport/internet"
	"github.com/aiempire79/mr-xray/transport/internet/stat"
)

func init() {
	common.Must(internet.RegisterTransportDialer(protocolName,
		func(ctx context.Context, dest net.Destination, streamSettings *internet.MemoryStreamConfig) (stat.Connection, error) {
			var sockopt *internet.SocketConfig
			if streamSettings != nil {
				sockopt = streamSettings.SocketSettings
			}
			conn, err := internet.DialSystem(ctx, dest, sockopt)
			if err != nil {
				return nil, err
			}
			// TODO: handle dialer options
			return stat.Connection(conn), nil
		}))
}
