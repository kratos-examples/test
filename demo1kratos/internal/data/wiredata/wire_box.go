package wiredata

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/yylego/kratos-examples/demo1kratos/internal/conf"
	"github.com/yylego/kratos-examples/demo1kratos/internal/data"
)

type WireBox struct {
	Data *data.Data
}

func newWireBox(d *data.Data) *WireBox {
	return &WireBox{
		Data: d,
	}
}

func NewWireBox(confData *conf.Data, logger log.Logger) (*WireBox, func(), error) {
	return wireApp(confData, logger)
}
