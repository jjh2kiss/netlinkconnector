package cnproc

import (
	"fmt"

	"github.com/jjh2kiss/netlinkconnector"
)

type ExitEvent struct {
	ProcessPid  int32
	ProcessTgid int32
	ExitCode    uint32
	ExitSignal  uint32
}

const SizeofExitEvent = 16

func DeserializeExitEvent(b []byte) (*ExitEvent, error) {
	if len(b) < SizeofExitEvent {
		return nil, fmt.Errorf("len(b) should be greater equal than %d", SizeofExitEvent)
	}

	native := netlinkconnector.NativeEndian()
	msg := new(ExitEvent)
	begin := 0
	msg.ProcessPid = int32(native.Uint32(b[begin:]))
	begin += 4
	msg.ProcessTgid = int32(native.Uint32(b[begin:]))
	begin += 4
	msg.ExitCode = native.Uint32(b[begin:])
	begin += 4
	msg.ExitSignal = native.Uint32(b[begin:])

	return msg, nil
}
