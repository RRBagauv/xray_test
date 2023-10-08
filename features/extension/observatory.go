package extension

import (
	"context"

	"google.golang.org/protobuf/proto"
	"xray-core/features"
)

type Observatory interface {
	features.Feature

	GetObservation(ctx context.Context) (proto.Message, error)
}

func ObservatoryType() interface{} {
	return (*Observatory)(nil)
}
