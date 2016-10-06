package cnproc

import (
	"fmt"

	"github.com/jjh2kiss/netlinkconnector"
)

/*
 * From the user's point of view, the process
 * ID is the thread group ID and thread ID is the internal
 * kernel "pid". So, fields are assigned as follow:
 *
 *  In user space     -  In  kernel space
 *
 * parent process ID  =  parent->tgid
 * parent thread  ID  =  parent->pid
 * child  process ID  =  child->tgid
 * child  thread  ID  =  child->pid
 */
type ForkEvent struct {
	ParentPid  int32
	ParentTgid int32
	ChildPid   int32
	ChildTgid  int32
}

const SizeofForkEvent = 16

func DeserializeForkEvent(b []byte) (*ForkEvent, error) {
	if len(b) < SizeofForkEvent {
		return nil, fmt.Errorf("len(b) should be greater equal than %d", SizeofForkEvent)
	}

	native := netlinkconnector.NativeEndian()
	msg := new(ForkEvent)
	begin := 0
	msg.ParentPid = int32(native.Uint32(b[begin:]))
	begin += 4
	msg.ParentTgid = int32(native.Uint32(b[begin:]))
	begin += 4
	msg.ChildPid = int32(native.Uint32(b[begin:]))
	begin += 4
	msg.ChildTgid = int32(native.Uint32(b[begin:]))

	return msg, nil
}
