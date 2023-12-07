package main

import (
	"os"

	"github.com/hemkantSplat/qt/core"
	"github.com/hemkantSplat/qt/gui"
	"github.com/hemkantSplat/qt/quick"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var translator = core.NewQTranslator(nil)
	translator.Load2(core.NewQLocale3(core.QLocale__French, core.QLocale__France), "translation", "_", ":/qml/i18n", ".qm")
	core.QCoreApplication_InstallTranslator(translator)

	var view = quick.NewQQuickView(nil)
	view.SetSource(core.NewQUrl3("qrc:/qml/translate.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	gui.QGuiApplication_Exec()
}
