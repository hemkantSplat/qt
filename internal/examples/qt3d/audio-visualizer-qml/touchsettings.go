package main

import (
	"os"
	"runtime"

	"github.com/hemkantSplat/qt/core"
	"github.com/hemkantSplat/qt/gui"
)

type TouchSettings struct {
	core.QObject

	_ func() bool `slot:"isHoverEnabled"`
}

func (s *TouchSettings) isHoverEnabled() bool {
	if runtime.GOOS == "darwin" && runtime.GOARCH == "arm64" || runtime.GOOS == "android" {
		return false
	}

	var isTouch bool
	for _, dev := range gui.QTouchDevice_Devices() {
		if dev.Type() == gui.QTouchDevice__TouchScreen {
			isTouch = true
		}
		break
	}

	var isMobile bool
	if os.Getenv("QT_QUICK_CONTROLS_MOBILE") != "" {
		isMobile = true
	}

	return !isTouch && !isMobile
}
