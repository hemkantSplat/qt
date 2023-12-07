//go:build !qml
// +build !qml

package main

import (
	"os"

	"github.com/hemkantSplat/qt/core"
	"github.com/hemkantSplat/qt/widgets"

	"github.com/hemkantSplat/qt/internal/examples/sql/masterdetail_qml/controller"

	"github.com/hemkantSplat/qt/internal/examples/sql/masterdetail_qml/view"
)

func main() {
	qApp := widgets.NewQApplication(len(os.Args), os.Args)

	controller.NewController(nil).InitWith(core.NewQFile2(":/albumdetails.xml"), qApp)

	view.NewViewController(nil, 0).Show()

	qApp.Exec()
}
