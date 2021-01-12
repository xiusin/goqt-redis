package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __stringkey struct{}

func (*__stringkey) init() {}

type StringKey struct {
	*__stringkey
	*widgets.QWidget
	VerticalLayoutWidget *widgets.QWidget
	VerticalLayout       *widgets.QVBoxLayout
	HorizontalLayout     *widgets.QHBoxLayout
	Label                *widgets.QLabel
	LineEdit             *widgets.QLineEdit
	HorizontalSpacer     *widgets.QSpacerItem
	PushButton_4         *widgets.QPushButton
	PushButton_5         *widgets.QPushButton
	PushButton_3         *widgets.QPushButton
	PlainTextEdit        *widgets.QPlainTextEdit
}

func NewStringKey(p widgets.QWidget_ITF) *StringKey {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &StringKey{QWidget: widgets.NewQWidget(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *StringKey) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("StringKey")
	}
	w.Resize2(652, 513)
	w.VerticalLayoutWidget = widgets.NewQWidget(w, 0)
	w.VerticalLayoutWidget.SetObjectName("verticalLayoutWidget")
	w.VerticalLayoutWidget.SetGeometry(core.NewQRect4(0, 0, 651, 511))
	w.VerticalLayout = widgets.NewQVBoxLayout2(w.VerticalLayoutWidget)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.VerticalLayout.SetContentsMargins(0, 0, 0, 0)
	w.HorizontalLayout = widgets.NewQHBoxLayout()
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.Label = widgets.NewQLabel(w.VerticalLayoutWidget, 0)
	w.Label.SetObjectName("label")
	w.HorizontalLayout.QLayout.AddWidget(w.Label)
	w.LineEdit = widgets.NewQLineEdit(w.VerticalLayoutWidget)
	w.LineEdit.SetObjectName("lineEdit")
	w.HorizontalLayout.QLayout.AddWidget(w.LineEdit)
	w.HorizontalSpacer = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout.AddItem(w.HorizontalSpacer)
	w.PushButton_4 = widgets.NewQPushButton(w.VerticalLayoutWidget)
	w.PushButton_4.SetObjectName("pushButton_4")
	w.HorizontalLayout.QLayout.AddWidget(w.PushButton_4)
	w.PushButton_5 = widgets.NewQPushButton(w.VerticalLayoutWidget)
	w.PushButton_5.SetObjectName("pushButton_5")
	w.HorizontalLayout.QLayout.AddWidget(w.PushButton_5)
	w.PushButton_3 = widgets.NewQPushButton(w.VerticalLayoutWidget)
	w.PushButton_3.SetObjectName("pushButton_3")
	w.HorizontalLayout.QLayout.AddWidget(w.PushButton_3)
	w.VerticalLayout.AddLayout(w.HorizontalLayout, 0)
	w.PlainTextEdit = widgets.NewQPlainTextEdit(w.VerticalLayoutWidget)
	w.PlainTextEdit.SetObjectName("plainTextEdit")
	w.VerticalLayout.QLayout.AddWidget(w.PlainTextEdit)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *StringKey) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("StringKey", "Form", "", 0))
	w.Label.SetText(core.QCoreApplication_Translate("StringKey", "STRING:", "", 0))
	w.PushButton_4.SetText(core.QCoreApplication_Translate("StringKey", "\345\210\240\351\231\244", "", 0))
	w.PushButton_5.SetText(core.QCoreApplication_Translate("StringKey", "\345\210\267\346\226\260\345\200\274", "", 0))
	w.PushButton_3.SetText(core.QCoreApplication_Translate("StringKey", "TTL:-1", "", 0))

}
