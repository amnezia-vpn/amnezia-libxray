package libXray

import (
	"syscall"

	xinternet "github.com/amnezia-vpn/amnezia-xray-core/transport/internet"
)

// Give a callback before connection beginning. Useful for Android development.
// It depends on xray:api:beta
func RegisterDialerController(controller DialerController) {
	xinternet.RegisterDialerController(func(network, address string, conn syscall.RawConn) error {
		return conn.Control(func(fd uintptr) {
			controller.ProtectFd(int(fd))
		})
	})
}

// Give a callback before listener beginning. Useful for Android development.
// It depends on xray:api:beta
func RegisterListenerController(controller DialerController) {
	xinternet.RegisterListenerController(func(network, address string, conn syscall.RawConn) error {
		return conn.Control(func(fd uintptr) {
			controller.ProtectFd(int(fd))
		})
	})
}
