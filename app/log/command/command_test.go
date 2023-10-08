package command_test

import (
	"context"
	"testing"

	"xray-core/app/dispatcher"
	"xray-core/app/log"
	. "xray-core/app/log/command"
	"xray-core/app/proxyman"
	_ "xray-core/app/proxyman/inbound"
	_ "xray-core/app/proxyman/outbound"
	"xray-core/common"
	"xray-core/common/serial"
	"xray-core/core"
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
