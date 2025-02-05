package main

import (
	"os"

	"github.com/hemkantSplat/qt/gui"
	"github.com/hemkantSplat/qt/quick"
)

func main() {
	EnableHighDPI()

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var view = quick.NewQQuickView(nil)
	view.SetSource(NewQUrl("qrc:/qml/view.qml"))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	gui.QGuiApplication_Exec()
}
