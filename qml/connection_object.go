package qml

import (
	"fmt"
	"goqt-redis/libs/helper"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type CtxConnectionObject struct {
	core.QObject
	ExitCh chan struct{}

	TreeWidget *widgets.QTreeWidget

	AttachIcon func(*widgets.QTreeWidgetItem, string)

	TestConn func(data map[string]interface{}) (bool, error)

	SaveConn func(map[string]interface{}) int

	_ func()                               `signal:"onClosing,auto"`
	_ func(string, string, string, string) `signal:"testServer,auto"`
	_ func(string, string, string, string) `signal:"saveServer,auto"`
}

func (receiver *CtxConnectionObject) onClosing() {
	go func() {
		receiver.ExitCh <- struct{}{}
	}()
}

func (receiver *CtxConnectionObject) testServer(name string, address string, port string, password string) {
	ok, err := receiver.TestConn(map[string]interface{}{
		"ip":    address,
		"port":  port,
		"title": name,
		"auth":  password,
	})
	if ok {
		helper.ShowInfoMessage("连接成功", "连接服务器成功")
	} else {
		helper.ShowWarningMessage("连接错误", err.Error())
	}
}

func (receiver *CtxConnectionObject) saveServer(name string, address string, port string, password string) {
	id := receiver.SaveConn(map[string]interface{}{
		"ip":    address,
		"port":  port,
		"title": name,
		"auth":  password,
	})
	item := widgets.NewQTreeWidgetItem(id)
	item.SetData(0, int(core.Qt__UserRole+0), core.NewQVariant1(fmt.Sprintf("%d", id)))
	item.SetData(0, int(core.Qt__UserRole+1), core.NewQVariant1("1"))
	receiver.AttachIcon(item, "redis")
	item.SetText(0, name)
	receiver.TreeWidget.AddTopLevelItem(item)

	receiver.onClosing()
}
