package wirebiz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/yylego/kratos-examples/demo1kratos/internal/biz"
	"github.com/yylego/kratos-examples/demo1kratos/internal/conf"
)

type WireBox struct {
	StudentUsecase *biz.StudentUsecase
}

func newWireBox(studentUsecase *biz.StudentUsecase) *WireBox {
	return &WireBox{
		StudentUsecase: studentUsecase,
	}
}

func NewWireBox(confData *conf.Data, logger log.Logger) (*WireBox, func(), error) {
	return wireApp(confData, logger)
}
