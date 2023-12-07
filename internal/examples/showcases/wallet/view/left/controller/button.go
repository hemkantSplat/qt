package controller

import (
	"github.com/hemkantSplat/qt/core"

	"github.com/hemkantSplat/qt/internal/examples/showcases/wallet/view/controller"
)

type buttonController struct {
	core.QObject

	_ func(code string) `signal:"clicked,auto"`
}

// lazy binding to the view/stack controller
func (c *buttonController) clicked(code string) { controller.StackController.Clicked(code) }
