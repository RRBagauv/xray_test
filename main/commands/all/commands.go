package all

import (
	"xray-core/main/commands/all/api"
	"xray-core/main/commands/all/tls"
	"xray-core/main/commands/base"
)

// go:generate go run xray-core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
	)
}
