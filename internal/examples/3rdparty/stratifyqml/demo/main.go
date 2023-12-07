//source: https://github.com/StratifyLabs/StratifyQML

package main

import (
	"os"

	"github.com/hemkantSplat/qt/core"
	"github.com/hemkantSplat/qt/qml"
	"github.com/hemkantSplat/qt/widgets"
)

func main() {

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	widgets.NewQApplication(len(os.Args), os.Args)

	engine := qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	widgets.QApplication_Exec()
}
