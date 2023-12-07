package main

import (
	"github.com/hemkantSplat/qt/core"
)

type Person struct {
	core.QObject

	_ string `property:"name"`
	_ int    `property:"shoeSize"`
}
