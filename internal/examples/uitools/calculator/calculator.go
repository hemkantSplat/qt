package main

import (
	"os"

	"github.com/hemkantSplat/qt/widgets"

	"github.com/hemkantSplat/qt/internal/examples/uitools/calculator/ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	ui.NewCalculatorForm(nil).Show()

	widgets.QApplication_Exec()
}
