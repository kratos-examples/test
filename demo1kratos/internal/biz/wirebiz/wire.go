//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package wirebiz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/yylego/kratos-examples/demo1kratos/internal/biz"
	"github.com/yylego/kratos-examples/demo1kratos/internal/conf"
	"github.com/yylego/kratos-examples/demo1kratos/internal/data"
)

// wireApp init kratos application.
func wireApp(*conf.Data, log.Logger) (*WireBox, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		biz.ProviderSet,
		newWireBox,
	))
}
