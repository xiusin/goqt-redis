package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __settingformui struct{}

func (*__settingformui) init() {}

type SettingFormUI struct {
	*__settingformui
	*widgets.QWidget
	FormLayoutWidget       *widgets.QWidget
	FormLayout             *widgets.QFormLayout
	VerticalSpacer_5       *widgets.QSpacerItem
	Label_1                *widgets.QLabel
	CheckBox               *widgets.QCheckBox
	VerticalSpacer_3       *widgets.QSpacerItem
	Label_2                *widgets.QLabel
	ServerIp               *widgets.QLineEdit
	VerticalSpacer_2       *widgets.QSpacerItem
	Label_3                *widgets.QLabel
	HorizontalLayout_2     *widgets.QHBoxLayout
	CheckBox_4             *widgets.QCheckBox
	CheckBox_3             *widgets.QCheckBox
	CheckBox_2             *widgets.QCheckBox
	CheckBox_5             *widgets.QCheckBox
	HorizontalSpacer       *widgets.QSpacerItem
	VerticalSpacer         *widgets.QSpacerItem
	Label_4                *widgets.QLabel
	CheckBox_6             *widgets.QCheckBox
	VerticalSpacer_4       *widgets.QSpacerItem
	Label_5                *widgets.QLabel
	ListView               *widgets.QListView
	HorizontalLayoutWidget *widgets.QWidget
	HorizontalLayout       *widgets.QHBoxLayout
	ConnAddBtn             *widgets.QPushButton
}

func NewSettingFormUI(p widgets.QWidget_ITF) *SettingFormUI {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &SettingFormUI{QWidget: widgets.NewQWidget(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *SettingFormUI) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("SettingFormUI")
	}
	w.Resize2(383, 510)
	w.FormLayoutWidget = widgets.NewQWidget(w, 0)
	w.FormLayoutWidget.SetObjectName("formLayoutWidget")
	w.FormLayoutWidget.SetGeometry(core.NewQRect4(0, 0, 381, 461))
	w.FormLayout = widgets.NewQFormLayout(w.FormLayoutWidget)
	w.FormLayout.SetObjectName("formLayout")
	w.FormLayout.SetContentsMargins(0, 0, 0, 0)
	w.VerticalSpacer_5 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.FormLayout.SetItem(1, widgets.QFormLayout__SpanningRole, w.VerticalSpacer_5)
	w.Label_1 = widgets.NewQLabel(w.FormLayoutWidget, 0)
	w.Label_1.SetObjectName("label_1")
	w.Label_1.SetMinimumSize(core.NewQSize2(100, 0))
	w.Label_1.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.FormLayout.SetWidget(2, widgets.QFormLayout__LabelRole, w.Label_1)
	w.CheckBox = widgets.NewQCheckBox(w.FormLayoutWidget)
	w.CheckBox.SetObjectName("checkBox")
	w.FormLayout.SetWidget(2, widgets.QFormLayout__FieldRole, w.CheckBox)
	w.VerticalSpacer_3 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.FormLayout.SetItem(3, widgets.QFormLayout__SpanningRole, w.VerticalSpacer_3)
	w.Label_2 = widgets.NewQLabel(w.FormLayoutWidget, 0)
	w.Label_2.SetObjectName("label_2")
	w.Label_2.SetMinimumSize(core.NewQSize2(100, 0))
	w.Label_2.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.FormLayout.SetWidget(4, widgets.QFormLayout__LabelRole, w.Label_2)
	w.ServerIp = widgets.NewQLineEdit(w.FormLayoutWidget)
	w.ServerIp.SetObjectName("serverIp")
	w.ServerIp.SetMinimumSize(core.NewQSize2(120, 0))
	w.ServerIp.SetMaximumSize(core.NewQSize2(120, 16777215))
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
	w.Label_3.SetMinimumSize(core.NewQSize2(100, 0))
	w.Label_3.SetBaseSize(core.NewQSize2(0, 0))
	w.Label_3.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.FormLayout.SetWidget(6, widgets.QFormLayout__LabelRole, w.Label_3)
	w.HorizontalLayout_2 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_2.SetObjectName("horizontalLayout_2")
	w.CheckBox_4 = widgets.NewQCheckBox(w.FormLayoutWidget)
	w.CheckBox_4.SetObjectName("checkBox_4")
	w.HorizontalLayout_2.QLayout.AddWidget(w.CheckBox_4)
	w.CheckBox_3 = widgets.NewQCheckBox(w.FormLayoutWidget)
	w.CheckBox_3.SetObjectName("checkBox_3")
	w.HorizontalLayout_2.QLayout.AddWidget(w.CheckBox_3)
	w.CheckBox_2 = widgets.NewQCheckBox(w.FormLayoutWidget)
	w.CheckBox_2.SetObjectName("checkBox_2")
	w.HorizontalLayout_2.QLayout.AddWidget(w.CheckBox_2)
	w.CheckBox_5 = widgets.NewQCheckBox(w.FormLayoutWidget)
	w.CheckBox_5.SetObjectName("checkBox_5")
	w.HorizontalLayout_2.QLayout.AddWidget(w.CheckBox_5)
	w.HorizontalSpacer = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout_2.AddItem(w.HorizontalSpacer)
	w.FormLayout.SetLayout(6, widgets.QFormLayout__FieldRole, w.HorizontalLayout_2)
	w.VerticalSpacer = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.FormLayout.SetItem(7, widgets.QFormLayout__SpanningRole, w.VerticalSpacer)
	w.Label_4 = widgets.NewQLabel(w.FormLayoutWidget, 0)
	w.Label_4.SetObjectName("label_4")
	w.Label_4.SetMinimumSize(core.NewQSize2(100, 0))
	w.Label_4.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.FormLayout.SetWidget(8, widgets.QFormLayout__LabelRole, w.Label_4)
	w.CheckBox_6 = widgets.NewQCheckBox(w.FormLayoutWidget)
	w.CheckBox_6.SetObjectName("checkBox_6")
	w.FormLayout.SetWidget(8, widgets.QFormLayout__FieldRole, w.CheckBox_6)
	w.VerticalSpacer_4 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.FormLayout.SetItem(9, widgets.QFormLayout__SpanningRole, w.VerticalSpacer_4)
	w.Label_5 = widgets.NewQLabel(w.FormLayoutWidget, 0)
	w.Label_5.SetObjectName("label_5")
	w.Label_5.SetMinimumSize(core.NewQSize2(100, 0))
	w.Label_5.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.FormLayout.SetWidget(10, widgets.QFormLayout__LabelRole, w.Label_5)
	w.ListView = widgets.NewQListView(w.FormLayoutWidget)
	w.ListView.SetObjectName("listView")
	w.FormLayout.SetWidget(10, widgets.QFormLayout__FieldRole, w.ListView)
	w.HorizontalLayoutWidget = widgets.NewQWidget(w, 0)
	w.HorizontalLayoutWidget.SetObjectName("horizontalLayoutWidget")
	w.HorizontalLayoutWidget.SetGeometry(core.NewQRect4(100, 460, 171, 51))
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w.HorizontalLayoutWidget)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.HorizontalLayout.SetContentsMargins(0, 0, 0, 0)
	w.ConnAddBtn = widgets.NewQPushButton(w.HorizontalLayoutWidget)
	w.ConnAddBtn.SetObjectName("connAddBtn")
	w.HorizontalLayout.QLayout.AddWidget(w.ConnAddBtn)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *SettingFormUI) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("SettingFormUI", "Form", "", 0))
	w.Label_1.SetText(core.QCoreApplication_Translate("SettingFormUI", "\347\277\273\350\257\221\345\277\253\346\215\267\351\224\256\357\274\232", "", 0))
	w.CheckBox.SetText(core.QCoreApplication_Translate("SettingFormUI", "\345\217\214\345\207\273alt\345\220\257\345\212\250\347\277\273\350\257\221\347\252\227\344\275\223", "", 0))
	w.Label_2.SetText(core.QCoreApplication_Translate("SettingFormUI", "\345\217\214\345\207\273\351\227\264\351\232\224(ms)\357\274\232", "", 0))
	w.Label_3.SetText(core.QCoreApplication_Translate("SettingFormUI", "\345\220\257\345\212\250\347\277\273\350\257\221\345\274\225\346\223\216\357\274\232", "", 0))
	w.CheckBox_4.SetText(core.QCoreApplication_Translate("SettingFormUI", "\350\205\276\350\256\257", "", 0))
	w.CheckBox_3.SetText(core.QCoreApplication_Translate("SettingFormUI", "\347\231\276\345\272\246", "", 0))
	w.CheckBox_2.SetText(core.QCoreApplication_Translate("SettingFormUI", "\346\234\211\351\201\223", "", 0))
	w.CheckBox_5.SetText(core.QCoreApplication_Translate("SettingFormUI", "\350\256\257\351\243\236", "", 0))
	w.Label_4.SetText(core.QCoreApplication_Translate("SettingFormUI", "\346\273\221\350\257\221\357\274\232", "", 0))
	w.CheckBox_6.SetText("")
	w.Label_5.SetText(core.QCoreApplication_Translate("SettingFormUI", "\345\220\257\347\224\250\350\217\234\345\215\225\357\274\232", "", 0))
	w.ConnAddBtn.SetText(core.QCoreApplication_Translate("SettingFormUI", "\347\241\256\345\256\232", "", 0))

}
