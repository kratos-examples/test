package service_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/require"
	"github.com/yylego/kratos-examples/demo1kratos"
	pb "github.com/yylego/kratos-examples/demo1kratos/api/student"
	"github.com/yylego/kratos-examples/demo1kratos/internal/pkg/cfgdata"
	"github.com/yylego/kratos-examples/demo1kratos/internal/service/wireservice"
	"github.com/yylego/must"
	"github.com/yylego/neatjson/neatjsons"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/zaplog"
	"go.uber.org/zap"
)

var caseWireBox *wireservice.WireBox

func TestMain(m *testing.M) {
	configPath := osmustexist.ROOT(filepath.Join(demo1kratos.SourceRoot(), "configs"))
	zaplog.LOG.Debug("run-test-main", zap.String("config-path", configPath))

	cfg := cfgdata.ParseConfig(configPath)
	wireBox, cleanup, err := wireservice.NewWireBox(cfg.Data, log.NewStdLogger(os.Stdout))
	must.Done(err)
	defer cleanup()

	caseWireBox = wireBox
	zaplog.LOG.Debug("run-test-main")
	m.Run()
}

func TestStudentService_CreateStudent(t *testing.T) {
	res, err := caseWireBox.StudentService.CreateStudent(context.Background(), &pb.CreateStudentRequest{
		Name: "test-abc",
	})
	require.NoError(t, err)
	t.Log(neatjsons.S(res))
	require.NotNil(t, res.Student)
}
