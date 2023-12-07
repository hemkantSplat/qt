package main

import (
	"os"

	"github.com/hemkantSplat/qt/core"
	"github.com/hemkantSplat/qt/gui"
	"github.com/hemkantSplat/qt/quick"

	_ "github.com/hemkantSplat/qt/internal/examples/qml/extending/components/test_qml/component"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/app.qml", 0))
	view.Show()

	gui.QGuiApplication_Exec()
}
