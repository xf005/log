package log

import (
	"context"
	"go.uber.org/zap"
	"testing"
)

func TestCodeTime(t *testing.T) {
	NewLogger(SetDev(true), SetLevel("info"), SetName("test"))
	ctx, _ := AddCtx(context.Background(), zap.String("traceId", "1234567890"))
	log := WithContext(ctx)
	log.Info("想测试什么呐")
	Infos(ctx, "另类测试。", "有没有啊！！！")
}
