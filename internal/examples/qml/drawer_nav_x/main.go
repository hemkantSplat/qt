//source: https://github.com/ekke/drawer_nav_x

package main

import (
	"os"

	"github.com/hemkantSplat/qt/core"
	"github.com/hemkantSplat/qt/gui"
	"github.com/hemkantSplat/qt/qml"
	"github.com/hemkantSplat/qt/quickcontrols2"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	gui.NewQGuiApplication(len(os.Args), os.Args)

	if quickcontrols2.QQuickStyle_Name() == "" {
		quickcontrols2.QQuickStyle_SetStyle("Material")
	}

	var translator = core.NewQTranslator(nil)
	if translator.Load2(core.NewQLocale(), "drawer_nav_x", "_", ":/translations", ".qm") {
		core.QCoreApplication_InstallTranslator(translator)
	} else {
		println("cannot load translator", core.QLocale_System().Name(), "check content of translations.qrc")
	}

	var appui = initApplicationUI()
	var engine = qml.NewQQmlApplicationEngine(nil)

	// from QML we have access to ApplicationUI as myApp
	var context = engine.RootContext()
	context.SetContextProperty("myApp", appui)

	engine.Load(core.NewQUrl3("qrc:/main.qml", 0))

	gui.QGuiApplication_Exec()
}
