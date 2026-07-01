//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package wiredata

import (
	"github.com/google/wire"
	"github.com/yylego/kratos-examples/demo1kratos/internal/conf"
	"github.com/yylego/kratos-examples/demo1kratos/internal/data"
	"log/slog"
)

// wireApp init kratos application.
func wireApp(*conf.Data, *slog.Logger) (*WireBox, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		newWireBox,
	))
}
