//author: https://github.com/longlongh4

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hemkantSplat/qt/core"
	"github.com/hemkantSplat/qt/gui"
	"github.com/hemkantSplat/qt/quick"
)

type QmlBridge struct {
	core.QObject

	_ func(data string)        `signal:"sendToQml"`
	_ func(data string) string `slot:"sendToGo"`
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var view = quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)

	var qmlBridge = NewQmlBridge(nil)
	qmlBridge.ConnectSendToGo(func(data string) string {
		fmt.Println("go:", data)
		return "hello from go"
	})

	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.SetSource(core.NewQUrl3("qrc:/qml/bridge2.qml", 0))

	go func() {
		for t := range time.NewTicker(time.Second * 1).C {
			qmlBridge.SendToQml(t.Format(time.ANSIC))
		}
	}()

	view.Show()

	gui.QGuiApplication_Exec()
}
