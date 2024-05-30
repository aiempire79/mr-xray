package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/aiempire79/mr-xray/app/dispatcher"
	_ "github.com/aiempire79/mr-xray/app/proxyman/inbound"
	_ "github.com/aiempire79/mr-xray/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/aiempire79/mr-xray/app/commander"
	_ "github.com/aiempire79/mr-xray/app/log/command"
	_ "github.com/aiempire79/mr-xray/app/proxyman/command"
	_ "github.com/aiempire79/mr-xray/app/stats/command"

	// Developer preview services
	_ "github.com/aiempire79/mr-xray/app/observatory/command"

	// Other optional features.
	_ "github.com/aiempire79/mr-xray/app/dns"
	_ "github.com/aiempire79/mr-xray/app/dns/fakedns"
	_ "github.com/aiempire79/mr-xray/app/log"
	_ "github.com/aiempire79/mr-xray/app/metrics"
	_ "github.com/aiempire79/mr-xray/app/policy"
	_ "github.com/aiempire79/mr-xray/app/reverse"
	_ "github.com/aiempire79/mr-xray/app/router"
	_ "github.com/aiempire79/mr-xray/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/aiempire79/mr-xray/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/aiempire79/mr-xray/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/aiempire79/mr-xray/proxy/blackhole"
	_ "github.com/aiempire79/mr-xray/proxy/dns"
	_ "github.com/aiempire79/mr-xray/proxy/dokodemo"
	_ "github.com/aiempire79/mr-xray/proxy/freedom"
	_ "github.com/aiempire79/mr-xray/proxy/http"
	_ "github.com/aiempire79/mr-xray/proxy/loopback"
	_ "github.com/aiempire79/mr-xray/proxy/shadowsocks"
	_ "github.com/aiempire79/mr-xray/proxy/socks"
	_ "github.com/aiempire79/mr-xray/proxy/trojan"
	_ "github.com/aiempire79/mr-xray/proxy/vless/inbound"
	_ "github.com/aiempire79/mr-xray/proxy/vless/outbound"
	_ "github.com/aiempire79/mr-xray/proxy/vmess/inbound"
	_ "github.com/aiempire79/mr-xray/proxy/vmess/outbound"
	_ "github.com/aiempire79/mr-xray/proxy/wireguard"

	// Transports
	_ "github.com/aiempire79/mr-xray/transport/internet/domainsocket"
	_ "github.com/aiempire79/mr-xray/transport/internet/grpc"
	_ "github.com/aiempire79/mr-xray/transport/internet/http"
	_ "github.com/aiempire79/mr-xray/transport/internet/httpupgrade"
	_ "github.com/aiempire79/mr-xray/transport/internet/kcp"
	_ "github.com/aiempire79/mr-xray/transport/internet/quic"
	_ "github.com/aiempire79/mr-xray/transport/internet/reality"
	_ "github.com/aiempire79/mr-xray/transport/internet/tcp"
	_ "github.com/aiempire79/mr-xray/transport/internet/tls"
	_ "github.com/aiempire79/mr-xray/transport/internet/udp"
	_ "github.com/aiempire79/mr-xray/transport/internet/websocket"

	// Transport headers
	_ "github.com/aiempire79/mr-xray/transport/internet/headers/http"
	_ "github.com/aiempire79/mr-xray/transport/internet/headers/noop"
	_ "github.com/aiempire79/mr-xray/transport/internet/headers/srtp"
	_ "github.com/aiempire79/mr-xray/transport/internet/headers/tls"
	_ "github.com/aiempire79/mr-xray/transport/internet/headers/utp"
	_ "github.com/aiempire79/mr-xray/transport/internet/headers/wechat"
	_ "github.com/aiempire79/mr-xray/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/aiempire79/mr-xray/main/json"
	_ "github.com/aiempire79/mr-xray/main/toml"
	_ "github.com/aiempire79/mr-xray/main/yaml"

	// Load config from file or http(s)
	_ "github.com/aiempire79/mr-xray/main/confloader/external"

	// Commands
	_ "github.com/aiempire79/mr-xray/main/commands/all"
)
