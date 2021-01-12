package components

import (
	"encoding/json"
	"strings"

	"github.com/therecipe/qt/widgets"
)

// Js 执行webView的js方法
type Js struct{}

// RunGetValue 点击事件
func (r *Js) RunGetValue(item *widgets.QTreeWidgetItem, titlePrefix string) {
	conns := strings.Split(item.Data(0, GetRole(RoleData)).ToString(), "-")
	key := item.Data(0, GetRole(RoleKey)).ToString()
	nodes := []map[string]string{
		{
			"action":   "get_value",
			"redis_id": conns[0],
			"index":    conns[1],
			"title":    key,
			"prefix":   titlePrefix,
		},
	}
	data, _ := json.Marshal(nodes)
	rdmUi.WebEngineView.Page().RunJavaScript("window.changeValue(" + string(data) + ")")
}

// RemoveKey 移除key
func (r *Js) RemoveKey(item *widgets.QTreeWidgetItem) {
	key := item.Data(0, GetRole(RoleKey)).ToString()
	rdmUi.WebEngineView.Page().RunJavaScript("window.removeKey(\"" + key + "\")")
}

// ShowAddNowKeyModal 展示添加key的模层
func (r *Js) ShowAddNowKeyModal(item *widgets.QTreeWidgetItem) {
	serverInfo := strings.Split(item.Data(0, GetRole(RoleData)).ToString(), "-")
	rdmUi.WebEngineView.Page().RunJavaScript("window.showAddKeyModal(" + serverInfo[0] + ", " + serverInfo[1] + ")")
}

// ShowCliModal 展示添加key的模层
func (r *Js) ShowCliModal(item *widgets.QTreeWidgetItem) {
	serverInfo := strings.Split(item.Data(0, GetRole(RoleData)).ToString(), "-")
	rdmUi.WebEngineView.Page().RunJavaScript("window.showCliModal(" + serverInfo[0] + ", " + serverInfo[1] + ")")
}

// ShowSlowLog 展示慢日志
func (r *Js) ShowSlowLog(item *widgets.QTreeWidgetItem) {
	rdmUi.WebEngineView.Page().RunJavaScript("window.showSlowLog(" + item.Data(0, GetRole(RoleData)).ToString() + ")")
}

// FlushDB 清空
func (r *Js) FlushDB(item *widgets.QTreeWidgetItem) {
	serverInfo := strings.Split(item.Data(0, GetRole(RoleData)).ToString(), "-")
	rdmUi.WebEngineView.Page().RunJavaScript("window.flushDB(" + serverInfo[0] + ", " + serverInfo[1] + ")")
}
