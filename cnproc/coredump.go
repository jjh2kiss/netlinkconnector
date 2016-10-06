package cnproc

import (
	"fmt"

	"github.com/jjh2kiss/netlinkconnector"
)

type CoredumpEvent struct {
	ProcessPid  int32
	ProcessTgid int32
}

const SizeofCoredumpEvent = 8

func DeserializeCoredumpEvent(b []byte) (*CoredumpEvent, error) {
	if len(b) < SizeofCoredumpEvent {
		return nil, fmt.Errorf("len(b) should be greater equal than %d", SizeofCoredumpEvent)
	}

	native := netlinkconnector.NativeEndian()
	msg := new(CoredumpEvent)
	begin := 0
	msg.ProcessPid = int32(native.Uint32(b[begin:]))
	begin += 4
	msg.ProcessTgid = int32(native.Uint32(b[begin:]))

	return msg, nil
}
