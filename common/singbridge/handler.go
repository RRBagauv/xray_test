package singbridge

import (
	"context"
	"io"

	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
	"xray-core/common/buf"
	"xray-core/common/errors"
	"xray-core/common/net"
	"xray-core/common/session"
	"xray-core/features/routing"
	"xray-core/transport"
)

var (
	_ N.TCPConnectionHandler = (*Dispatcher)(nil)
	_ N.UDPConnectionHandler = (*Dispatcher)(nil)
)

type Dispatcher struct {
	upstream     routing.Dispatcher
	newErrorFunc func(values ...any) *errors.Error
}

func NewDispatcher(dispatcher routing.Dispatcher, newErrorFunc func(values ...any) *errors.Error) *Dispatcher {
	return &Dispatcher{
		upstream:     dispatcher,
		newErrorFunc: newErrorFunc,
	}
}

func (d *Dispatcher) NewConnection(ctx context.Context, conn net.Conn, metadata M.Metadata) error {
	xConn := NewConn(conn)
	return d.upstream.DispatchLink(ctx, ToDestination(metadata.Destination, net.Network_TCP), &transport.Link{
		Reader: xConn,
		Writer: xConn,
	})
}

func (d *Dispatcher) NewPacketConnection(ctx context.Context, conn N.PacketConn, metadata M.Metadata) error {
	return d.upstream.DispatchLink(ctx, ToDestination(metadata.Destination, net.Network_UDP), &transport.Link{
		Reader: buf.NewPacketReader(conn.(io.Reader)),
		Writer: buf.NewWriter(conn.(io.Writer)),
	})
}

func (d *Dispatcher) NewError(ctx context.Context, err error) {
	d.newErrorFunc(err).WriteToLog(session.ExportIDToError(ctx))
}
