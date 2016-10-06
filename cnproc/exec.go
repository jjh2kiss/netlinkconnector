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
type ExecEvent struct {
	ProcessPid  int32
	ProcessTgid int32
}

const SizeofExecEvent = 8

func DeserializeExecEvent(b []byte) (*ExecEvent, error) {
	if len(b) < SizeofExecEvent {
		return nil, fmt.Errorf("len(b) should be greater equal than %d", SizeofExecEvent)
	}

	native := netlinkconnector.NativeEndian()
	msg := new(ExecEvent)
	begin := 0
	msg.ProcessPid = int32(native.Uint32(b[begin:]))
	begin += 4
	msg.ProcessTgid = int32(native.Uint32(b[begin:]))
	return msg, nil
}
