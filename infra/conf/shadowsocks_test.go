package conf_test

import (
	"testing"

	"github.com/aiempire79/mr-xray/common/net"
	"github.com/aiempire79/mr-xray/common/protocol"
	"github.com/aiempire79/mr-xray/common/serial"
	. "github.com/aiempire79/mr-xray/infra/conf"
	"github.com/aiempire79/mr-xray/proxy/shadowsocks"
)

func TestShadowsocksServerConfigParsing(t *testing.T) {
	creator := func() Buildable {
		return new(ShadowsocksServerConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"method": "aes-256-GCM",
				"password": "xray-password"
			}`,
			Parser: loadJSON(creator),
			Output: &shadowsocks.ServerConfig{
				Users: []*protocol.User{{
					Account: serial.ToTypedMessage(&shadowsocks.Account{
						CipherType: shadowsocks.CipherType_AES_256_GCM,
						Password:   "xray-password",
					}),
				}},
				Network: []net.Network{net.Network_TCP},
			},
		},
	})
}
