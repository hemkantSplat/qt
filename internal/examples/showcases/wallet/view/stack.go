package view

import (
	"github.com/hemkantSplat/qt/quick"

	_ "github.com/hemkantSplat/qt/internal/examples/showcases/wallet/view/controller"
)

func init() { stackTemplate_QmlRegisterType2("ViewTemplate", 1, 0, "StackTemplate") }

type stackTemplate struct {
	quick.QQuickItem

	_ func(code string) `signal:"Clicked,<-(controller.NewStackController(nil))"`
}
