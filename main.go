package main

import (
	"os"

	"goqt-redis/components"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
)

func main() {
	quickcontrols2.QQuickStyle_SetStyle("Fusion")
	core.QCoreApplication_SetAttribute(core.Qt__AA_UseSoftwareOpenGL, true)
	core.QCoreApplication_SetAttribute(core.Qt__AA_UseDesktopOpenGL, true)
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	core.QCoreApplication_SetAttribute(core.Qt__AA_ShareOpenGLContexts, true)
	core.QCoreApplication_SetAttribute(core.Qt__AA_UseOpenGLES, true)

	widgets.NewQApplication(len(os.Args), os.Args)

	fontDB := gui.NewQFontDatabase()
	fontPath := core.QCoreApplication_ApplicationDirPath() + "/msyh.ttc"
	fontID := fontDB.AddApplicationFont(fontPath)
	strs := gui.QFontDatabase_ApplicationFontFamilies(fontID)
	if len(strs) > 0 {
		font := gui.NewQFont()
		font.FromString(strs[0])
		font.SetWeight(14)
		widgets.QApplication_SetFont(font, "")
	}
	go components.InitRdm()
	components.InitRdmUI()
	widgets.QApplication_Exec()
}
