package main

import (
	"os"

	"github.com/hemkantSplat/qt/core"
	"github.com/hemkantSplat/qt/gui"
	"github.com/hemkantSplat/qt/qml"
	"github.com/hemkantSplat/qt/quickcontrols2"
)

func main() {

	core.QCoreApplication_SetApplicationName("Gallery")
	core.QCoreApplication_SetOrganizationName("QtProject")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var (
		settings = core.NewQSettings5(nil)
		style    = quickcontrols2.QQuickStyle_Name()
	)
	if style != "" {
		settings.SetValue("style", core.NewQVariant1(style))
	} else {
		quickcontrols2.QQuickStyle_SetStyle(settings.Value("style", core.NewQVariant1("")).ToString())
	}

	var engine = qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/gallery.qml", 0))

	gui.QGuiApplication_Exec()
}
