package all

import (
	"github.com/aiempire79/mr-xray/main/commands/all/api"
	"github.com/aiempire79/mr-xray/main/commands/all/tls"
	"github.com/aiempire79/mr-xray/main/commands/base"
)

// go:generate go run github.com/aiempire79/mr-xray/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
	)
}
