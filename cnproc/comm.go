package cnproc

import (
	"fmt"

	"github.com/jjh2kiss/netlinkconnector"
)

type CommEvent struct {
	ProcessPid  int32
	ProcessTgid int32
	Comm        [16]byte
}

const SizeofCommEvent = 24

func DeserializeCommEvent(b []byte) (*CommEvent, error) {
	if len(b) < SizeofCommEvent {
		return nil, fmt.Errorf("len(b) should be greater equal than %d", SizeofCommEvent)
	}

	native := netlinkconnector.NativeEndian()
	msg := new(CommEvent)
	begin := 0
	msg.ProcessPid = int32(native.Uint32(b[begin:]))
	begin += 4
	msg.ProcessTgid = int32(native.Uint32(b[begin:]))
	begin += 4

	for i := 0; i < 16; i++ {
		msg.Comm[i] = b[begin+i]
	}

	return msg, nil
}
