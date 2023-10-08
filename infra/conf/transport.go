package conf

import (
	"xray-core/common/serial"
	"xray-core/transport/global"
	"xray-core/transport/internet"
)

type TransportConfig struct {
	TCPConfig  *TCPConfig          `json:"tcpSettings"`
	WSConfig   *WebSocketConfig    `json:"wsSettings"`
	DSConfig   *DomainSocketConfig `json:"dsSettings"`
	QUICConfig *QUICConfig         `json:"quicSettings"`
}

// Build implements Buildable.
func (c *TransportConfig) Build() (*global.Config, error) {
	config := new(global.Config)

	if c.TCPConfig != nil {
		ts, err := c.TCPConfig.Build()
		if err != nil {
			return nil, newError("failed to build TCP config").Base(err).AtError()
		}
		config.TransportSettings = append(config.TransportSettings, &internet.TransportConfig{
			ProtocolName: "tcp",
			Settings:     serial.ToTypedMessage(ts),
		})
	}

	if c.QUICConfig != nil {
		qs, err := c.QUICConfig.Build()
		if err != nil {
			return nil, newError("Failed to build QUIC config.").Base(err)
		}
		config.TransportSettings = append(config.TransportSettings, &internet.TransportConfig{
			ProtocolName: "quic",
			Settings:     serial.ToTypedMessage(qs),
		})
	}

	return config, nil
}
