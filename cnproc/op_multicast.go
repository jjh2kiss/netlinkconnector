package cnproc

import (
	"github.com/jjh2kiss/netlinkconnector"
)

const SizeofMulticastOp = 4

/*
#linux/cn_proc.h
* Userspace sends this enum to register with the kernel that it is listening
* for events on the connector.
*/

const (
	PROC_CN_MCAST_LISTEN = 1
	PROC_CN_MCAST_IGNORE = 2
)

type MultiCastOp struct {
	Val uint32
}

func (self *MultiCastOp) Len() int {
	return netlinkconnector.NlmAlignOf(SizeofMulticastOp)
}

func (self *MultiCastOp) Serialize() []byte {
	native := netlinkconnector.NativeEndian()
	buf := make([]byte, self.Len())
	native.PutUint32(buf[0:4], self.Val)
	return buf
}
