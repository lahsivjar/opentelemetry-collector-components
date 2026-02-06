package customottl

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
)

func DatapointFuncs() map[string]ottl.Factory[*ottldatapoint.TransformContext] {
	common := commonFuncs[*ottlspan.TransformContext]()
	clientMetaFactory := NewClientMetadataFactory()
	common[clientMetaFactory.Name()] = clientMetaFactory
	return common
}
