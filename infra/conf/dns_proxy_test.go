package conf_test

import (
	"testing"

	"github.com/aiempire79/mr-xray/common/net"
	. "github.com/aiempire79/mr-xray/infra/conf"
	"github.com/aiempire79/mr-xray/proxy/dns"
)

func TestDnsProxyConfig(t *testing.T) {
	creator := func() Buildable {
		return new(DNSOutboundConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"address": "8.8.8.8",
				"port": 53,
				"network": "tcp"
			}`,
			Parser: loadJSON(creator),
			Output: &dns.Config{
				Server: &net.Endpoint{
					Network: net.Network_TCP,
					Address: net.NewIPOrDomain(net.IPAddress([]byte{8, 8, 8, 8})),
					Port:    53,
				},
				Non_IPQuery: "drop",
			},
		},
	})
}
