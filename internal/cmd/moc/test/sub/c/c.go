package c

import (
	"github.com/hemkantSplat/qt/core"

	_ "github.com/hemkantSplat/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoC struct{}
type StructSubMocC struct{ core.QObject }
