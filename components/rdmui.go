package components

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"goqt-redis/libs/helper"
	"goqt-redis/libs/rdm"
	"goqt-redis/qml"
	"goqt-redis/ui"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/network"
	qml2 "github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/webengine"
	"github.com/therecipe/qt/widgets"
	"github.com/xiusin/logger"
)

// RdmUi 窗体实体
type RdmUi struct {
	*ui.RdmUi
	ConnectionServerIsShow bool
	isMousePressed         bool
	connectedServer        map[string]bool
	mousePoint             *core.QPoint
	Js                     *Js
	DbTitles               map[string]string
}

var rdmUi *RdmUi
var rdmOnce sync.Once

// InitRdmUi 初始化UI
func InitRdmUI() {
	rdmOnce.Do(func() {
		network.QNetworkProxyFactory_SetUseSystemConfiguration(false) // 屏蔽自动寻找代理
		rdmUi = &RdmUi{
			RdmUi:           ui.NewRdmUi(nil),
			Js:              &Js{},
			DbTitles:        map[string]string{},
			connectedServer: map[string]bool{},
		}
		rdmUi.SetWindowFlag(core.Qt__FramelessWindowHint, true) // 屏蔽webview动画闪烁窗体.
		rdmUi.ConnectMousePressEvent(func(event *gui.QMouseEvent) {
			rdmUi.mousePoint = event.Pos() // 记录鼠标相对于窗体的位置
			rdmUi.isMousePressed = true
			event.Accept()
		})

		rdmUi.ConnectMouseMoveEvent(func(event *gui.QMouseEvent) {
			if rdmUi.isMousePressed {
				curMousePoint := event.GlobalPos()
				rdmUi.Move2(curMousePoint.X()-rdmUi.mousePoint.X(), curMousePoint.Y()-rdmUi.mousePoint.Y())
				event.Accept()
			}
		})

		rdmUi.ConnectMouseReleaseEvent(func(event *gui.QMouseEvent) {
			rdmUi.isMousePressed = false
			rdmUi.mousePoint = nil
			event.Accept()
		})

		rdmUi.CloseBtn.SetToolTip("关闭RDM")
		rdmUi.CloseBtn.ConnectClicked(func(checked bool) {
			rdmUi.Close()
		})
		rdmUi.WebEngineView.Settings().SetAttribute(webengine.QWebEngineSettings__PluginsEnabled, true)
		rdmUi.WebEngineView.Settings().SetAttribute(webengine.QWebEngineSettings__SpatialNavigationEnabled, true)
		url := fmt.Sprintf("http://127.0.0.1:%d", serverPort)
		rdmUi.WebEngineView.SetUrl(core.NewQUrl3(url+"/#/?qt_web_port="+strconv.Itoa(serverPort), 0))
		rdmUi.VerticalLayout_2.SetSpacing(0)
		rdmUi.HorizontalLayout.SetSpacing(0)
		qss, err := helper.GetFileContent("qss/rdm.css")
		if err == nil {
			rdmUi.SetStyleSheet(string(qss))
		}

		rdmUi.HorizontalLayout.SetSizeConstraint(widgets.QLayout__SetFixedSize)
		NewTreeComp(rdmUi.TreeWidget, rdmUi).RegisterEvent()
		rdmUi.ConnectionServer.ConnectClicked(func(checked bool) {
			if !rdmUi.ConnectionServerIsShow {
				rdmUi.ConnectionServerIsShow = true
				view := qml2.NewQQmlApplicationEngine(nil)
				ctxObject := qml.NewCtxConnectionObject(nil)
				ctxObject.ExitCh = make(chan struct{})
				ctxObject.TreeWidget = rdmUi.TreeWidget
				ctxObject.SaveConn = rdm.RedisManagerConfigSaveForQt
				ctxObject.TestConn = rdm.RedisManagerConnectionTestForQt
				ctxObject.AttachIcon = rdmUi.AttachIcon

				view.RootContext().SetContextProperty("ctxObject", ctxObject)
				view.Load(core.NewQUrl3("qrc:/qml/AddServerDialog.qml", 0))

				go func() {
					<-ctxObject.ExitCh
					rdmUi.ConnectionServerIsShow = false
					ctxObject.ExitCh = nil
					ctxObject.DestroyQObject()
				}()
			}

		})
		rdmUi.TreeWidget.ConnectItemDoubleClicked(func(item *widgets.QTreeWidgetItem, column int) {
			level, err := strconv.Atoi(item.Data(0, GetRole(RoleLevel)).ToString())
			if err == nil && level == 3 {
				info := strings.Split(item.Parent().Text(0), " ")
				titlePre := fmt.Sprintf("%s::%s", item.Parent().Parent().Text(0), strings.ToUpper(info[0]))
				rdmUi.Js.RunGetValue(item, titlePre)
			}
		})
		rdmUi.SetTreeData()
	})
	rdmUi.TreeWidget.SetVisible(true)
	rdmUi.TreeWidget.Show()
	rdmUi.Show()
}

const (
	RoleData = iota
	RoleLevel
	RoleKey
	RolePattern
	RoleConnFlag
)

func GetRole(va int) int {
	return int(core.Qt__UserRole) + va
}

func (ui *RdmUi) AttachIcon(item *widgets.QTreeWidgetItem, normalPngName string) {
	iconRedis := gui.NewQIcon()
	mp, mp1 := gui.NewQPixmap(), gui.NewQPixmap()
	mp.Load(":/qml/qrc/rdm/"+normalPngName+".png", "", core.Qt__AutoColor)
	iconRedis.AddPixmap(mp, gui.QIcon__Normal, gui.QIcon__Off)
	mp1.Load(":/qml/qrc/rdm/"+normalPngName+"_select.png", "", core.Qt__AutoColor)
	iconRedis.AddPixmap(mp1, gui.QIcon__Selected, gui.QIcon__On)
	iconRedis.AddPixmap(mp1, gui.QIcon__Active, gui.QIcon__On)
	item.SetIcon(0, iconRedis)
}

func (ui *RdmUi) SetTreeData() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(err)
		}
	}()

	ui.TreeWidget.SetFixedWidth(320)

	cones := rdm.RedisManagerConnectionListForQt()
	for k, conn := range cones {
		ui.connectedServer[conn.Flag] = false
		item := widgets.NewQTreeWidgetItem(k)
		item.SetData(0, GetRole(RoleData), core.NewQVariant1(fmt.Sprintf("%d", conn.ID)))
		item.SetData(0, GetRole(RoleLevel), core.NewQVariant1("1"))
		item.SetData(0, GetRole(RoleConnFlag), core.NewQVariant1(conn.Flag))
		ui.AttachIcon(item, "redis")
		item.SetText(0, conn.Title)
		ui.TreeWidget.AddTopLevelItem(item)
	}

	ui.TreeWidget.ConnectItemDoubleClicked(func(item *widgets.QTreeWidgetItem, column int) {
		level, _ := strconv.Atoi(item.Data(0, GetRole(RoleLevel)).ToString())
		switch level {
		case 1:
			if item.ChildCount() > 0 {
				return
			}
			flag := item.Data(0, GetRole(RoleConnFlag)).ToString()
			ui.connectedServer[flag] = true
			serverIdx := item.Data(0, GetRole(RoleData)).ToString()
			dbs := rdm.RedisManagerDbListForQt(serverIdx, true)
			if len(dbs) > 0 {
				for k, db := range dbs {
					newItem := widgets.NewQTreeWidgetItem(k)
					newItem.SetData(0, GetRole(RoleData), core.NewQVariant1(fmt.Sprintf("%s-%d", serverIdx, k)))
					newItem.SetData(0, GetRole(RoleLevel), core.NewQVariant1("2"))
					newItem.SetData(0, GetRole(RolePattern), core.NewQVariant1("*")) // 默认正则规则为*
					ui.AttachIcon(newItem, "db")
					newItem.SetText(0, fmt.Sprintf("db%2d  (%d)", k, db))
					item.AddChild(newItem)
				}
			} else {
				ui.connectedServer[flag] = false
			}
		case 2:
			if item.ChildCount() > 0 {
				return
			}
			serverInfo := strings.Split(item.Data(0, GetRole(RoleData)).ToString(), "-")
			pattern := item.Data(0, GetRole(RolePattern)).ToString()
			if pattern == "" {
				pattern = "*"
			}
			keys := rdm.RedisManagerConnectionSelectDbForQt(rdm.RequestData{
				"id":      serverInfo[0],
				"index":   serverInfo[1],
				"pattern": pattern,
			})
			for idx, k := range keys {
				newItem := widgets.NewQTreeWidgetItem(idx)
				newItem.SetData(0, GetRole(RoleData), core.NewQVariant1(fmt.Sprintf("%s-%s", serverInfo[0], serverInfo[1])))
				newItem.SetData(0, GetRole(RoleLevel), core.NewQVariant1("3"))
				newItem.SetData(0, GetRole(RoleKey), core.NewQVariant1(k))
				newItem.SetText(0, k)
				ui.AttachIcon(newItem, "key")
				item.AddChild(newItem)
			}
		}
	})

	ui.TreeWidget.Show()
}
