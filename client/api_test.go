package sms_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/hrygo/gosms/auth"
	"github.com/hrygo/gosms/bootstrap"
	sms "github.com/hrygo/gosms/client"
)

var _ = bootstrap.BasePath

func init() {
	// 启动记录数据库的程序
	if sms.Conf.GetString("Mongo.URI") != "" {
		sms.PersistenceSmsJournal()
	} else {
		sms.StartCacheExpireTicker(nil)
	}
	auth.Cache = auth.New(bootstrap.ConfigYml)
}

func TestSend(t *testing.T) {
	i := 100
	for i > 0 {
		queryId := sms.Send("hello world", "13800001111")
		assert.True(t, len(sms.Query(queryId)) > 0)
		i--
	}
	i = 100
	for i > 0 {
		queryId := sms.Send("hello world", "13300001111")
		assert.True(t, len(sms.Query(queryId)) > 0)
		i--
	}
	time.Sleep(time.Minute)
}
