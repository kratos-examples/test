package wireservice

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/yylego/kratos-examples/demo1kratos/internal/conf"
	"github.com/yylego/kratos-examples/demo1kratos/internal/service"
)

type WireBox struct {
	StudentService *service.StudentService
}

func newWireBox(studentService *service.StudentService) *WireBox {
	return &WireBox{
		StudentService: studentService,
	}
}

func NewWireBox(confData *conf.Data, logger log.Logger) (*WireBox, func(), error) {
	return wireApp(confData, logger)
}
