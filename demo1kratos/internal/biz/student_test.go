package biz_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/require"
	"github.com/yylego/kratos-examples/demo1kratos"
	"github.com/yylego/kratos-examples/demo1kratos/internal/biz"
	"github.com/yylego/kratos-examples/demo1kratos/internal/biz/wirebiz"
	"github.com/yylego/kratos-examples/demo1kratos/internal/pkg/cfgdata"
	"github.com/yylego/must"
	"github.com/yylego/neatjson/neatjsons"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

var caseWireBox *wirebiz.WireBox

func TestMain(m *testing.M) {
	configPath := osmustexist.ROOT(filepath.Join(demo1kratos.SourceRoot(), "configs"))
	zaplog.LOG.Debug("run-test-main", zap.String("config-path", configPath))

	cfg := cfgdata.ParseConfig(configPath)
	wireBox, cleanup, err := wirebiz.NewWireBox(cfg.Data, log.NewStdLogger(os.Stdout))
	must.Done(err)
	defer cleanup()

	caseWireBox = wireBox
	zaplog.LOG.Debug("run-test-main")
	m.Run()
}

func TestStudentUsecase_CreateStudent(t *testing.T) {
	res, ebz := caseWireBox.StudentUsecase.CreateStudent(context.Background(), &biz.Student{
		Name: "abc",
	})
	require.Nil(t, ebz)
	t.Log(neatjsons.S(res))
	require.NotEmpty(t, res.Name)
}
