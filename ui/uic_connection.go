package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __connserverwin struct{}

func (*__connserverwin) init() {}

type ConnServerWin struct {
	*__connserverwin
	*widgets.QWidget
	FormLayoutWidget       *widgets.QWidget
	FormLayout             *widgets.QFormLayout
	VerticalSpacer_5       *widgets.QSpacerItem
	Label_1                *widgets.QLabel
	ServerName             *widgets.QLineEdit
	VerticalSpacer_3       *widgets.QSpacerItem
	Label_2                *widgets.QLabel
	ServerIp               *widgets.QLineEdit
	VerticalSpacer_2       *widgets.QSpacerItem
	Label_3                *widgets.QLabel
	ServerPort             *widgets.QLineEdit
	VerticalSpacer         *widgets.QSpacerItem
	Label_4                *widgets.QLabel
	ServerPwd              *widgets.QLineEdit
	VerticalSpacer_4       *widgets.QSpacerItem
	HorizontalLayoutWidget *widgets.QWidget
	HorizontalLayout       *widgets.QHBoxLayout
	ConnAddBtn             *widgets.QPushButton
	ConnTestBtn            *widgets.QPushButton
}

func NewConnServerWin(p widgets.QWidget_ITF) *ConnServerWin {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &ConnServerWin{QWidget: widgets.NewQWidget(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *ConnServerWin) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("ConnServerWin")
	}
	w.Resize2(420, 268)
	w.FormLayoutWidget = widgets.NewQWidget(w, 0)
	w.FormLayoutWidget.SetObjectName("formLayoutWidget")
	w.FormLayoutWidget.SetGeometry(core.NewQRect4(20, 0, 381, 231))
	w.FormLayout = widgets.NewQFormLayout(w.FormLayoutWidget)
	w.FormLayout.SetObjectName("formLayout")
	w.FormLayout.SetContentsMargins(0, 0, 0, 0)
	w.VerticalSpacer_5 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.FormLayout.SetItem(1, widgets.QFormLayout__SpanningRole, w.VerticalSpacer_5)
	w.Label_1 = widgets.NewQLabel(w.FormLayoutWidget, 0)
	w.Label_1.SetObjectName("label_1")
	w.Label_1.SetMinimumSize(core.NewQSize2(80, 0))
	w.Label_1.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.FormLayout.SetWidget(2, widgets.QFormLayout__LabelRole, w.Label_1)
	w.ServerName = widgets.NewQLineEdit(w.FormLayoutWidget)
	w.ServerName.SetObjectName("serverName")
	w.ServerName.SetMinimumSize(core.NewQSize2(0, 0))
	w.ServerName.SetClearButtonEnabled(true)
	w.FormLayout.SetWidget(2, widgets.QFormLayout__FieldRole, w.ServerName)
	w.VerticalSpacer_3 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.FormLayout.SetItem(3, widgets.QFormLayout__SpanningRole, w.VerticalSpacer_3)
	w.Label_2 = widgets.NewQLabel(w.FormLayoutWidget, 0)
	w.Label_2.SetObjectName("label_2")
	w.Label_2.SetMinimumSize(core.NewQSize2(80, 0))
	w.Label_2.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.FormLayout.SetWidget(4, widgets.QFormLayout__LabelRole, w.Label_2)
	w.ServerIp = widgets.NewQLineEdit(w.FormLayoutWidget)
	w.ServerIp.SetObjectName("serverIp")
	w.ServerIp.SetClearButtonEnabled(true)
	w.FormLayout.SetWidget(4, widgets.QFormLayout__FieldRole, w.ServerIp)
	w.VerticalSpacer_2 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.FormLayout.SetItem(5, widgets.QFormLayout__SpanningRole, w.VerticalSpacer_2)
	w.Label_3 = widgets.NewQLabel(w.FormLayoutWidget, 0)
	w.Label_3.SetObjectName("label_3")
	sizePolicy := widgets.NewQSizePolicy2(widgets.QSizePolicy__Preferred, widgets.QSizePolicy__Preferred, 0)
	sizePolicy.SetHorizontalStretch(0)
	sizePolicy.SetVerticalStretch(0)
	sizePolicy.SetHeightForWidth(w.Label_3.SizePolicy().HasHeightForWidth())
	w.Label_3.SetSizePolicy(sizePolicy)
	w.Label_3.SetMinimumSize(core.NewQSize2(80, 0))
	w.Label_3.SetBaseSize(core.NewQSize2(0, 0))
	w.Label_3.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.FormLayout.SetWidget(6, widgets.QFormLayout__LabelRole, w.Label_3)
	w.ServerPort = widgets.NewQLineEdit(w.FormLayoutWidget)
	w.ServerPort.SetObjectName("serverPort")
	w.ServerPort.SetClearButtonEnabled(true)
	w.FormLayout.SetWidget(6, widgets.QFormLayout__FieldRole, w.ServerPort)
	w.VerticalSpacer = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.FormLayout.SetItem(7, widgets.QFormLayout__SpanningRole, w.VerticalSpacer)
	w.Label_4 = widgets.NewQLabel(w.FormLayoutWidget, 0)
	w.Label_4.SetObjectName("label_4")
	w.Label_4.SetMinimumSize(core.NewQSize2(80, 0))
	w.Label_4.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.FormLayout.SetWidget(8, widgets.QFormLayout__LabelRole, w.Label_4)
	w.ServerPwd = widgets.NewQLineEdit(w.FormLayoutWidget)
	w.ServerPwd.SetObjectName("serverPwd")
	w.ServerPwd.SetEchoMode(widgets.QLineEdit__Password)
	w.ServerPwd.SetClearButtonEnabled(true)
	w.FormLayout.SetWidget(8, widgets.QFormLayout__FieldRole, w.ServerPwd)
	w.VerticalSpacer_4 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.FormLayout.SetItem(9, widgets.QFormLayout__SpanningRole, w.VerticalSpacer_4)
	w.HorizontalLayoutWidget = widgets.NewQWidget(w, 0)
	w.HorizontalLayoutWidget.SetObjectName("horizontalLayoutWidget")
	w.HorizontalLayoutWidget.SetGeometry(core.NewQRect4(20, 230, 381, 34))
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w.HorizontalLayoutWidget)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.HorizontalLayout.SetContentsMargins(0, 0, 0, 0)
	w.ConnAddBtn = widgets.NewQPushButton(w.HorizontalLayoutWidget)
	w.ConnAddBtn.SetObjectName("connAddBtn")
	w.HorizontalLayout.QLayout.AddWidget(w.ConnAddBtn)
	w.ConnTestBtn = widgets.NewQPushButton(w.HorizontalLayoutWidget)
	w.ConnTestBtn.SetObjectName("connTestBtn")
	w.HorizontalLayout.QLayout.AddWidget(w.ConnTestBtn)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *ConnServerWin) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("ConnServerWin", "\346\267\273\345\212\240\346\226\260\346\234\215\345\212\241\345\231\250", "", 0))
	w.Label_1.SetText(core.QCoreApplication_Translate("ConnServerWin", "\345\220\215\347\247\260\357\274\232", "", 0))
	w.Label_2.SetText(core.QCoreApplication_Translate("ConnServerWin", "\345\234\260\345\235\200\357\274\232", "", 0))
	w.Label_3.SetText(core.QCoreApplication_Translate("ConnServerWin", "\347\253\257\345\217\243\357\274\232", "", 0))
	w.Label_4.SetText(core.QCoreApplication_Translate("ConnServerWin", "\345\257\206\347\240\201\357\274\232", "", 0))
	w.ConnAddBtn.SetText(core.QCoreApplication_Translate("ConnServerWin", "\347\241\256\345\256\232", "", 0))
	w.ConnTestBtn.SetText(core.QCoreApplication_Translate("ConnServerWin", "\346\265\213\350\257\225", "", 0))

}
