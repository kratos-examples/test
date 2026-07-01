package wireservice

import (
	"github.com/yylego/kratos-examples/demo1kratos/internal/conf"
	"github.com/yylego/kratos-examples/demo1kratos/internal/service"
	"log/slog"
)

type WireBox struct {
	StudentService *service.StudentService
}

func newWireBox(studentService *service.StudentService) *WireBox {
	return &WireBox{
		StudentService: studentService,
	}
}

func NewWireBox(confData *conf.Data, logger *slog.Logger) (*WireBox, func(), error) {
	return wireApp(confData, logger)
}
