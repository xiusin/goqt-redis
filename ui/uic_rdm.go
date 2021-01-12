package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/webengine"
	"github.com/therecipe/qt/widgets"
)

type __rdmui struct{}

func (*__rdmui) init() {}

type RdmUi struct {
	*__rdmui
	*widgets.QMainWindow
	Centralwidget            *widgets.QWidget
	VerticalLayoutWidget     *widgets.QWidget
	VerticalLayout           *widgets.QVBoxLayout
	HorizontalLayout         *widgets.QHBoxLayout
	TreeWidget               *widgets.QTreeWidget
	VerticalLayout_2         *widgets.QVBoxLayout
	WebViewBorderLine        *widgets.QFrame
	WebEngineView            *webengine.QWebEngineView
	HorizontalLayoutWidget_2 *widgets.QWidget
	HorizontalLayout_2       *widgets.QHBoxLayout
	ConnectionServer         *widgets.QPushButton
	HorizontalSpacer         *widgets.QSpacerItem
	CloseBtn                 *widgets.QPushButton
}

func NewRdmUi(p widgets.QWidget_ITF) *RdmUi {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &RdmUi{QMainWindow: widgets.NewQMainWindow(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *RdmUi) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("RdmUi")
	}
	w.Resize2(1253, 793)
	w.Centralwidget = widgets.NewQWidget(w, 0)
	w.Centralwidget.SetObjectName("centralwidget")
	w.Centralwidget.SetStyleSheet("")
	w.VerticalLayoutWidget = widgets.NewQWidget(w.Centralwidget, 0)
	w.VerticalLayoutWidget.SetObjectName("verticalLayoutWidget")
	w.VerticalLayoutWidget.SetGeometry(core.NewQRect4(0, 40, 1251, 751))
	w.VerticalLayout = widgets.NewQVBoxLayout2(w.VerticalLayoutWidget)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.VerticalLayout.SetContentsMargins(0, 0, 0, 0)
	w.HorizontalLayout = widgets.NewQHBoxLayout()
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.TreeWidget = widgets.NewQTreeWidget(w.VerticalLayoutWidget)
	__qtreewidgetitem := widgets.NewQTreeWidgetItem(0)
	__qtreewidgetitem.SetText(0, "1")
	w.TreeWidget.SetHeaderItem(__qtreewidgetitem)
	w.TreeWidget.SetObjectName("treeWidget")
	w.TreeWidget.SetMinimumSize(core.NewQSize2(240, 0))
	w.TreeWidget.SetMaximumSize(core.NewQSize2(310, 16777215))
	w.TreeWidget.SetContextMenuPolicy(core.Qt__CustomContextMenu)
	w.TreeWidget.SetEditTriggers(widgets.QAbstractItemView__CurrentChanged | widgets.QAbstractItemView__DoubleClicked | widgets.QAbstractItemView__EditKeyPressed)
	w.TreeWidget.Header().SetVisible(false)
	w.HorizontalLayout.QLayout.AddWidget(w.TreeWidget)
	w.VerticalLayout_2 = widgets.NewQVBoxLayout()
	w.VerticalLayout_2.SetObjectName("verticalLayout_2")
	w.WebViewBorderLine = widgets.NewQFrame(w.VerticalLayoutWidget, 0)
	w.WebViewBorderLine.SetObjectName("webViewBorderLine")
	w.WebViewBorderLine.SetStyleSheet("border-top: 1px solid #ccc")
	w.WebViewBorderLine.SetFrameShape(widgets.QFrame__HLine)
	w.WebViewBorderLine.SetFrameShadow(widgets.QFrame__Sunken)
	w.VerticalLayout_2.QLayout.AddWidget(w.WebViewBorderLine)
	w.WebEngineView = webengine.NewQWebEngineView(w.VerticalLayoutWidget)
	w.WebEngineView.SetObjectName("webEngineView")
	font := gui.NewQFont()
	font.SetStyleStrategy(gui.QFont__PreferAntialias)
	w.WebEngineView.SetFont(font)
	w.WebEngineView.SetContextMenuPolicy(core.Qt__NoContextMenu)
	w.WebEngineView.SetUrl(core.NewQUrl3("http://localhost:6784/#/", 0))
	w.VerticalLayout_2.QLayout.AddWidget(w.WebEngineView)
	w.HorizontalLayout.AddLayout(w.VerticalLayout_2, 0)
	w.VerticalLayout.AddLayout(w.HorizontalLayout, 0)
	w.HorizontalLayoutWidget_2 = widgets.NewQWidget(w.Centralwidget, 0)
	w.HorizontalLayoutWidget_2.SetObjectName("horizontalLayoutWidget_2")
	w.HorizontalLayoutWidget_2.SetGeometry(core.NewQRect4(0, 0, 1251, 41))
	w.HorizontalLayout_2 = widgets.NewQHBoxLayout2(w.HorizontalLayoutWidget_2)
	w.HorizontalLayout_2.SetSpacing(20)
	w.HorizontalLayout_2.SetObjectName("horizontalLayout_2")
	w.HorizontalLayout_2.SetContentsMargins(5, 0, 0, 0)
	w.ConnectionServer = widgets.NewQPushButton(w.HorizontalLayoutWidget_2)
	w.ConnectionServer.SetObjectName("connectionServer")
	w.HorizontalLayout_2.QLayout.AddWidget(w.ConnectionServer)
	w.HorizontalSpacer = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout_2.AddItem(w.HorizontalSpacer)
	w.CloseBtn = widgets.NewQPushButton(w.HorizontalLayoutWidget_2)
	w.CloseBtn.SetObjectName("CloseBtn")
	w.CloseBtn.SetMinimumSize(core.NewQSize2(32, 32))
	w.CloseBtn.SetMaximumSize(core.NewQSize2(32, 32))
	font1 := gui.NewQFont()
	font1.SetKerning(true)
	w.CloseBtn.SetFont(font1)
	w.HorizontalLayout_2.QLayout.AddWidget(w.CloseBtn)
	w.SetCentralWidget(w.Centralwidget)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *RdmUi) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("RdmUi", "RedisManager Desktop", "", 0))
	if true {
		w.WebEngineView.SetStatusTip("")
	}
	w.ConnectionServer.SetText(core.QCoreApplication_Translate("RdmUi", "\350\277\236\346\216\245\346\234\215\345\212\241\345\231\250", "", 0))
	w.CloseBtn.SetText("")

}
