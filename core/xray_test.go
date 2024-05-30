package core_test

import (
	"testing"

	"github.com/aiempire79/mr-xray/app/dispatcher"
	"github.com/aiempire79/mr-xray/app/proxyman"
	"github.com/aiempire79/mr-xray/common"
	"github.com/aiempire79/mr-xray/common/net"
	"github.com/aiempire79/mr-xray/common/protocol"
	"github.com/aiempire79/mr-xray/common/serial"
	"github.com/aiempire79/mr-xray/common/uuid"
	. "github.com/aiempire79/mr-xray/core"
	"github.com/aiempire79/mr-xray/features/dns"
	"github.com/aiempire79/mr-xray/features/dns/localdns"
	_ "github.com/aiempire79/mr-xray/main/distro/all"
	"github.com/aiempire79/mr-xray/proxy/dokodemo"
	"github.com/aiempire79/mr-xray/proxy/vmess"
	"github.com/aiempire79/mr-xray/proxy/vmess/outbound"
	"github.com/aiempire79/mr-xray/testing/servers/tcp"
	"google.golang.org/protobuf/proto"
)

func TestXrayDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	})
	instance.AddFeature(localdns.New())
	<-wait
}

func TestXrayClose(t *testing.T) {
	port := tcp.PickPort()

	userID := uuid.New()
	config := &Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{
						Range: []*net.PortRange{net.SinglePortRange(port)},
					},
					Listen: net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userID.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}
