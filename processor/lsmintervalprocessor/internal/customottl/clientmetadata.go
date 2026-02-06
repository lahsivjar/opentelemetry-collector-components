package customottl

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"go.opentelemetry.io/collector/client"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func NewClientMetadataFactory() ottl.Factory[*ottldatapoint.TransformContext] {
	return ottl.NewFactory("ClientMetadata", nil, createClientMetadataFunc)
}

func createClientMetadataFunc(ottl.FunctionContext, ottl.Arguments) (ottl.ExprFunc[*ottldatapoint.TransformContext], error) {
	return clientMetadata()
}

func clientMetadata() (ottl.ExprFunc[*ottldatapoint.TransformContext], error) {
	return func(ctx context.Context, tCtx *ottldatapoint.TransformContext) (any, error) {
		return convertClientMetadataToMap(client.FromContext(ctx).Metadata), nil
	}, nil
}

func convertClientMetadataToMap(md client.Metadata) pcommon.Map {
	mdMap := pcommon.NewMap()
	for k := range md.Keys() {
		convertStringArrToValueSlice(md.Get(k)).MoveTo(mdMap.PutEmpty(k))
	}
	return mdMap
}

func convertStringArrToValueSlice(vals []string) pcommon.Value {
	val := pcommon.NewValueSlice()
	sl := val.Slice()
	sl.EnsureCapacity(len(vals))
	for _, val := range vals {
		sl.AppendEmpty().SetStr(val)
	}
	return val
}
