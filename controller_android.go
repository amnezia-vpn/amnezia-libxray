package libXray

import (
	"syscall"

	"github.com/amnezia-vpn/amnezia-libxray/nodep"
	xinternet "github.com/xtls/xray-core/transport/internet"
)

// Give a callback before connection beginning. Useful for Android development.
// It depends on xray:api:beta
func RegisterDialerController(controller DialerController) string {
	err := xinternet.RegisterDialerController(func(network, address string, conn syscall.RawConn) error {
		return conn.Control(func(fd uintptr) {
			controller.ProtectFd(int(fd))
		})
	})
	return nodep.WrapError(err)
}

// Give a callback before listener beginning. Useful for Android development.
// It depends on xray:api:beta
func RegisterListenerController(controller DialerController) string {
	err := xinternet.RegisterListenerController(func(network, address string, conn syscall.RawConn) error {
		return conn.Control(func(fd uintptr) {
			controller.ProtectFd(int(fd))
		})
	})
	return nodep.WrapError(err)
}
