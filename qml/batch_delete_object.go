package qml

import (
	"goqt-redis/libs/helper"
	"goqt-redis/libs/rdm"
	"strings"

	"github.com/therecipe/qt/core"
)

type CtxObject struct {
	core.QObject

	ExitCh chan struct{}

	_ string `property:"patternKey"`

	ServerIdx string
	DbIdx     string

	_ func(string) `signal:"onEditingFinished,auto"`
	_ func(string) `signal:"onClicked,auto"`
}

func (ptr *CtxObject) onClicked(keyPattern string) {
	go func() {
		rdm.RedisManagerBatchDeleteForQt(rdm.RequestData{
			"id":      ptr.ServerIdx,
			"index":   ptr.DbIdx,
			"pattern": keyPattern,
		})
		ptr.ExitCh <- struct{}{}
		helper.ShowInfoMessage("提醒", "指定表达式"+keyPattern+"的key已删除")
	}()
}

func (ptr *CtxObject) onEditingFinished(keyPattern string) {
	keys := rdm.RedisManagerConnectionSelectDbForQt(rdm.RequestData{
		"id":      ptr.ServerIdx,
		"index":   ptr.DbIdx,
		"pattern": keyPattern,
	})

	// 将key渲染成label
	ptr.SetPatternKey(strings.Join(keys, "\n"))
}
