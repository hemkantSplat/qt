package main

import (
	"os"

	"github.com/hemkantSplat/qt/core"
	"github.com/hemkantSplat/qt/gui"
	"github.com/hemkantSplat/qt/quick"
	"github.com/hemkantSplat/qt/sailfish"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	sailfish.SailfishApp_Application(len(os.Args), os.Args) //gui.NewQGuiApplication(len(os.Args), os.Args)

	var view = sailfish.SailfishApp_CreateView() //quick.NewQQuickView(nil)
	view.SetSource(core.NewQUrl3("qrc:/qml/sailfish.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	gui.QGuiApplication_Exec()
}
