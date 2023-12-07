package swift

import "github.com/hemkantSplat/qt/interop"

func init() {
	interop.ReturnPointersAsStrings = true
	interop.SupportsSyncCallsIntoRemote = false
}
