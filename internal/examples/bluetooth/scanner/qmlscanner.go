//source: https://doc.qt.io/qt-5/qtbluetooth-scanner-example.html

package main

import (
	"os"

	"github.com/hemkantSplat/qt/core"
	"github.com/hemkantSplat/qt/gui"
	"github.com/hemkantSplat/qt/quick"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	//core.QLoggingCategory_SetFilterRules("qt.bluetooth* = true")
	app := gui.NewQGuiApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetSource(core.NewQUrl3("qrc:/scanner.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Engine().ConnectQuit(app.Quit)
	view.SetGeometry2(core.NewQRect4(100, 100, 360, 640))
	view.Show()

	gui.QGuiApplication_Exec()
}
