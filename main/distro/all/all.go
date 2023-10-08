package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "xray-core/app/dispatcher"
	_ "xray-core/app/proxyman/inbound"
	_ "xray-core/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "xray-core/app/log/command"
	_ "xray-core/app/proxyman/command"
	_ "xray-core/app/stats/command"

	// Other optional features.
	_ "xray-core/app/dns"
	_ "xray-core/app/dns/fakedns"
	_ "xray-core/app/log"
	_ "xray-core/app/policy"
	_ "xray-core/app/reverse"
	_ "xray-core/app/router"
	_ "xray-core/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "xray-core/transport/internet/tagged/taggedimpl"

	// Inbound and outbound proxies.
	_ "xray-core/proxy/http"
	_ "xray-core/proxy/loopback"
	_ "xray-core/proxy/socks"
	_ "xray-core/proxy/vless/inbound"
	_ "xray-core/proxy/vless/outbound"

	// Transports
	_ "xray-core/transport/internet/quic"
	_ "xray-core/transport/internet/reality"
	_ "xray-core/transport/internet/tcp"
	_ "xray-core/transport/internet/tls"
	_ "xray-core/transport/internet/udp"

	// Transport headers
	_ "xray-core/transport/internet/headers/http"
	_ "xray-core/transport/internet/headers/noop"
	_ "xray-core/transport/internet/headers/srtp"
	_ "xray-core/transport/internet/headers/tls"

	// JSON & TOML & YAML
	_ "xray-core/main/json"

	// Load config from file or http(s)
	_ "xray-core/main/confloader/external"

	// Commands
	_ "xray-core/main/commands/all"
)
