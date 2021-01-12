package helper

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/xiusin/logger"
)

var workingDir string

func init() {
	workingDir, _ = os.Getwd()
}

// GetFileContent 获取文件内容
func GetFileContent(name string) ([]byte, error) {
	byts, err := ioutil.ReadFile(filepath.Join(workingDir, name))
	if err != nil {
		logger.Error("Error reading file: " + err.Error())
	}
	return byts, err
}

// ShowInfoMessage 弹出信息窗口
func ShowInfoMessage(title, content string) {
	widgets.NewQMessageBox2(widgets.QMessageBox__Information, title, content, widgets.QMessageBox__NoButton, nil, core.Qt__ToolTip).Show()
}

// ShowWarningMessage 弹出警告窗口
func ShowWarningMessage(title, content string) {
	widgets.NewQMessageBox2(widgets.QMessageBox__Warning, title, content, widgets.QMessageBox__NoButton, nil, core.Qt__ToolTip).Show()
}

// UserHomeDir 获取缓存目录
func UserHomeDir(file ...string) string {
	usr, _ := user.Current()
	homeDir := usr.HomeDir + "/.goqt_rdm"
	os.MkdirAll(homeDir, os.ModePerm)
	if len(file) != 0 {
		return homeDir + "/" + file[0]
	}
	return homeDir
}
