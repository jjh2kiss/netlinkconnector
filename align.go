package netlinkconnector

import (
	"syscall"
)

// Round the length of a netlink message up to align it properly.
// copy from src/syscall/netlink_linux.go
func NlmAlignOf(msglen int) int {
	return (msglen + syscall.NLMSG_ALIGNTO - 1) & ^(syscall.NLMSG_ALIGNTO - 1)
}
