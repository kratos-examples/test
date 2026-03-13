package data_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/require"
	"github.com/yylego/kratos-examples/demo1kratos"
	"github.com/yylego/kratos-examples/demo1kratos/internal/data/wiredata"
	"github.com/yylego/kratos-examples/demo1kratos/internal/pkg/cfgdata"
	"github.com/yylego/must"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

var caseWireBox *wiredata.WireBox

func TestMain(m *testing.M) {
	configPath := osmustexist.ROOT(filepath.Join(demo1kratos.SourceRoot(), "configs"))
	zaplog.LOG.Debug("run-test-main", zap.String("config-path", configPath))

	cfg := cfgdata.ParseConfig(configPath)
	wireBox, cleanup, err := wiredata.NewWireBox(cfg.Data, log.NewStdLogger(os.Stdout))
	must.Done(err)
	defer cleanup()

	caseWireBox = wireBox
	zaplog.LOG.Debug("run-test-main")
	m.Run()
}

func TestData_NotNil(t *testing.T) {
	require.NotNil(t, caseWireBox.Data)
}
