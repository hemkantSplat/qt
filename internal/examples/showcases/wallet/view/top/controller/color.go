package controller

import (
	"github.com/hemkantSplat/qt/core"

	"github.com/hemkantSplat/qt/internal/examples/showcases/wallet/theme/controller"
)

type colorController struct {
	core.QObject

	_ func() `signal:"change,auto"`
}

// lazy binding to the (qml singleton) theme controller
func (c *colorController) change() { controller.Controller.Change() }
