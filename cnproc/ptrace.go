package cnproc

import (
	"fmt"

	"github.com/jjh2kiss/netlinkconnector"
)

type PtraceEvent struct {
	ProcessPid  int32
	ProcessTgid int32
	TracerPid   int32
	TracerTgid  int32
}

const SizeofPtraceEvent = 16

func DeserializePtraceEvent(b []byte) (*PtraceEvent, error) {
	if len(b) < SizeofPtraceEvent {
		return nil, fmt.Errorf("len(b) should be greater equal than %d", SizeofPtraceEvent)
	}

	native := netlinkconnector.NativeEndian()
	msg := new(PtraceEvent)
	begin := 0
	msg.ProcessPid = int32(native.Uint32(b[begin:]))
	begin += 4
	msg.ProcessTgid = int32(native.Uint32(b[begin:]))
	begin += 4
	msg.TracerPid = int32(native.Uint32(b[begin:]))
	begin += 4
	msg.TracerTgid = int32(native.Uint32(b[begin:]))

	return msg, nil
}
