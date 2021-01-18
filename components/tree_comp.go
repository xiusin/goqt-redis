package components

import (
	"fmt"
	"strconv"
	"strings"

	"goqt-redis/libs/channel"
	"goqt-redis/libs/helper"
	"goqt-redis/libs/rdm"
	"goqt-redis/qml"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	qml2 "github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

// TreeComp 树组件
type TreeComp struct {
	rdm *RdmUi
	*widgets.QTreeWidget
}

// NewTreeComp 初始化
// https://github.com/gabrielchristo/custom-treewidget/blob/master/MainWindow.cpp
func NewTreeComp(tw *widgets.QTreeWidget, rdmUi *RdmUi) *TreeComp {
	return &TreeComp{rdmUi, tw}
}

func (receiver *TreeComp) RegisterEvent() {
	receiver.ConnectCustomContextMenuRequested(receiver.slotContextMenu)
	go receiver.keyChangeEvent()
}

func (receiver *TreeComp) keyChangeEvent() {
	for {
		select {
		case updateInfo := <-channel.UpdateKeyCh:
			dbItem := receiver.QTreeWidget.TopLevelItem(updateInfo.ServeIdx).Child(updateInfo.DbIdx)
			var childCount int
			childCount = dbItem.ChildCount()
			switch updateInfo.Action {
			case channel.ADD:
				newItem := widgets.NewQTreeWidgetItem(childCount)
				newItem.SetData(0, GetRole(RoleData), core.NewQVariant1(fmt.Sprintf("%d-%d", updateInfo.ServeId, updateInfo.DbIdx)))
				newItem.SetData(0, GetRole(RoleLevel), core.NewQVariant1("3"))
				newItem.SetData(0, GetRole(RoleKey), core.NewQVariant1(updateInfo.Key))
				newItem.SetText(0, updateInfo.Key)
				receiver.rdm.AttachIcon(newItem, "key")
				dbItem.AddChild(newItem)
				childCount++
			}
			dbItem.SetText(0, fmt.Sprintf("db%2d  (%d)", updateInfo.DbIdx, childCount))
		}
	}
}

func (receiver *TreeComp) slotContextMenu(point *core.QPoint) {
	currentItem := receiver.QTreeWidget.ItemAt(point)
	if currentItem == nil {
		return
	}
	// 鼠标右键
	level, _ := strconv.Atoi(currentItem.Data(0, GetRole(RoleLevel)).ToString())
	receiver.QTreeWidget.ClearSelection()
	menu := widgets.NewQMenu(receiver.QTreeWidget)
	switch level {
	case 1:
		flag := currentItem.Data(0, GetRole(RoleConnFlag)).ToString()
		var exists bool
		var ok bool
		if exists, ok = receiver.rdm.connectedServer[flag]; !ok {
			receiver.rdm.connectedServer[flag] = ok
		}
		if exists {
			receiver.attachContextMenuToConn(menu, currentItem)
		}
	case 2:
		receiver.attachContextMenuToDb(menu, currentItem)
	case 3:
		receiver.attachContextMenuToKey(menu, currentItem)
	default:
		return
	}

	m, err := helper.GetFileContent("qss/tree_menu.css")
	if err == nil {
		menu.SetStyleSheet(string(m))
	}
	menu.Exec2(receiver.QTreeWidget.Viewport().MapToGlobal(point), nil)
}

func (receiver *TreeComp) RemoveAllSubItem(item *widgets.QTreeWidgetItem) func(bool) {
	return func(bool) {
		item.SetExpanded(false)
		childNum := item.ChildCount() // 移除下级节点
		for childNum > 0 {
			childNum--
			c := item.Child(childNum)
			ccNum := c.ChildCount()
			for ccNum > 0 {
				ccNum--
				c.RemoveChild(c.Child(ccNum))
			}
			item.RemoveChild(c)
		}
	}
}

func (receiver *TreeComp) attachContextMenuToDb(menu *widgets.QMenu, item *widgets.QTreeWidgetItem) {

	actionReload := menu.AddAction("刷新")
	receiver.attachMenuIcon(actionReload, "refresh")
	actionReload.SetParent(receiver.QTreeWidget)
	actionReload.ConnectTriggered(func(checked bool) {
		receiver.RemoveAllSubItem(item)(checked)
		item.TreeWidget().ItemDoubleClicked(item, 0)
		item.SetExpanded(true)
	})

	actionAddNewKey := menu.AddAction("添加")
	receiver.attachMenuIcon(actionAddNewKey, "add")
	actionAddNewKey.SetParent(receiver.QTreeWidget)
	actionAddNewKey.ConnectTriggered(func(checked bool) {
		receiver.rdm.Js.ShowAddNowKeyModal(item)
	})

	actionPattenFilter := menu.AddAction("筛选")
	receiver.attachMenuIcon(actionPattenFilter, "filter")
	actionPattenFilter.SetParent(receiver.QTreeWidget)
	actionPattenFilter.ConnectTriggered(func(checked bool) {
		txt := widgets.NewQLineEdit(receiver.rdm)
		pattern := item.Data(0, GetRole(RolePattern)).ToString()
		if pattern == "" {
			pattern = "*"
		}
		txt.SetText(pattern)
		txt.ConnectKeyReleaseEvent(func(event *gui.QKeyEvent) { // keydown 删除
			if event.Key() == int(core.Qt__Key_Return) || int(core.Qt__Key_Enter) == event.Key() {
				txtContent := txt.Text()
				if txtContent == "" {
					txtContent = "*"
				}
				item.SetData(0, GetRole(RolePattern), core.NewQVariant1(txtContent))
				receiver.rdm.TreeWidget.RemoveItemWidget(item, 0)
				txt.DestroyQLineEdit()
			}
		})
		receiver.rdm.TreeWidget.SetItemWidget(item, 0, txt)
	})

	actionOpenCli := menu.AddAction("命令行")
	receiver.attachMenuIcon(actionOpenCli, "cli")
	actionOpenCli.SetParent(receiver.QTreeWidget)
	actionOpenCli.ConnectTriggered(func(checked bool) {
		receiver.rdm.Js.ShowCliModal(item)
	})

	actionFlush := menu.AddAction("清空")
	receiver.attachMenuIcon(actionFlush, "clear")
	actionFlush.SetParent(receiver.QTreeWidget)
	actionFlush.ConnectTriggered(func(checked bool) {
		if widgets.QMessageBox_Question(nil, "清空提醒", "确定要清空: "+item.Text(0)+"吗, 该操作不可恢复!",
			widgets.QMessageBox__Yes|widgets.QMessageBox__No, widgets.QMessageBox__No) == widgets.QMessageBox__Yes {
			receiver.RemoveAllSubItem(item)(checked)
			receiver.rdm.Js.FlushDB(item)
			//item.DestroyQTreeWidgetItem()
		}
	})

	actionPattenDel := menu.AddAction("批量删除")
	receiver.attachMenuIcon(actionPattenDel, "batchdel")
	actionPattenDel.SetParent(receiver.QTreeWidget)
	actionPattenDel.ConnectTriggered(func(checked bool) {
		view := qml2.NewQQmlApplicationEngine(nil)
		ctxObject := qml.NewCtxObject(nil)
		ctxObject.ExitCh = make(chan struct{})
		serverInfo := strings.Split(item.Data(0, GetRole(RoleData)).ToString(), "-")
		ctxObject.ServerIdx = serverInfo[0]
		ctxObject.DbIdx = serverInfo[1]
		view.RootContext().SetContextProperty("ctxObject", ctxObject)
		view.Load(core.NewQUrl3("qrc:/qml/BatchDelete.qml", 0))

		go func() {
			<-ctxObject.ExitCh
			// 刷新菜单节点
			receiver.RemoveAllSubItem(item)(checked)
			item.TreeWidget().ItemDoubleClicked(item, 0)
			item.SetExpanded(true)

			ctxObject.ExitCh = nil
			ctxObject.DestroyQObject()
		}()
	})
}

func (receiver *TreeComp) attachContextMenuToConn(menu *widgets.QMenu, item *widgets.QTreeWidgetItem) {
	disConn := menu.AddAction("断开")
	receiver.attachMenuIcon(disConn, "discon")
	disConn.SetParent(receiver.QTreeWidget)
	disConn.ConnectTriggered(func(checked bool) {
		receiver.RemoveAllSubItem(item)(checked)
		flag := item.Data(0, GetRole(RoleConnFlag)).ToString()
		delete(receiver.rdm.connectedServer, flag)
	})
	serveInfo := menu.AddAction("信息")
	receiver.attachMenuIcon(serveInfo, "info")
	serveInfo.SetParent(receiver.QTreeWidget)
	serveInfo.ConnectTriggered(func(checked bool) {
		receiver.rdm.Js.ShowSlowLog(item)
	})
	actionDel := menu.AddAction("删除")
	receiver.attachMenuIcon(actionDel, "delete")
	actionDel.SetParent(receiver.QTreeWidget)
	actionDel.ConnectTriggered(func(checked bool) {
		disConn.Triggered(true)
		ok := rdm.RedisManagerRemoveConnectionForQt(map[string]interface{}{
			"id": item.Data(0, GetRole(RoleData)).ToString(),
		})
		flag := item.Data(0, GetRole(RoleConnFlag)).ToString()
		delete(receiver.rdm.connectedServer, flag)
		if ok {
			item.DestroyQTreeWidgetItem()
		}
	})
}

func (receiver *TreeComp) attachContextMenuToKey(menu *widgets.QMenu, item *widgets.QTreeWidgetItem) {
	actionDel := menu.AddAction("删除")
	receiver.attachMenuIcon(actionDel, "delete")
	actionDel.SetParent(receiver.QTreeWidget)
	actionDel.ConnectTriggered(func(checked bool) { // 移除所有的节点
		if widgets.QMessageBox_Question(nil, "删除提醒", "确定要删除: "+item.Text(0),
			widgets.QMessageBox__Yes|widgets.QMessageBox__No, widgets.QMessageBox__No) == widgets.QMessageBox__Yes {
			receiver.RemoveAllSubItem(item)(checked)
			receiver.rdm.Js.RemoveKey(item)
			item.DestroyQTreeWidgetItem()
		}
	})
}

func (receiver *TreeComp) attachMenuIcon(action *widgets.QAction, normalPngName string) {
	action.SetIcon(gui.NewQIcon5(fmt.Sprintf(":/qml/qrc/icon/%s.png", normalPngName)))
}
