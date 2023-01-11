package runtimeproxy

import (
	"CRTProxy/pkg/runtimeproxy/server/cri"
)

func main() {
	criServer := cri.RuntimeCriServer{}
	criServer.Run()
}
