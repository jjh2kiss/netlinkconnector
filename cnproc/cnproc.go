package cnproc

import (
	"fmt"

	"github.com/jjh2kiss/netlinkconnector"
)

const (
	PROC_EVENT_NONE   = 0x00000000
	PROC_EVENT_FORK   = 0x00000001
	PROC_EVENT_EXEC   = 0x00000002
	PROC_EVENT_UID    = 0x00000004
	PROC_EVENT_GID    = 0x00000040
	PROC_EVENT_SID    = 0x00000080
	PROC_EVENT_PTRACE = 0x00000100
	PROC_EVENT_COMM   = 0x00000200
	/* "next" should be 0x00000400 */
	/* "last" is the last process event: exit,
	 * while "next to last" is coredumping event */
	PROC_EVENT_COREDUMP = 0x40000000
	PROC_EVENT_EXIT     = 0x80000000
)

func ProcEventString(v uint32) string {
	switch v {
	case PROC_EVENT_NONE:
		return "none"
	case PROC_EVENT_FORK:
		return "fork"
	case PROC_EVENT_EXEC:
		return "exec"
	case PROC_EVENT_UID:
		return "uid"
	case PROC_EVENT_GID:
		return "gid"
	case PROC_EVENT_SID:
		return "sid"
	case PROC_EVENT_PTRACE:
		return "ptrace"
	case PROC_EVENT_COMM:
		return "comm"
	case PROC_EVENT_COREDUMP:
		return "coredump"
	case PROC_EVENT_EXIT:
		return "exit"
	}
	return "unknown"
}

const SizeofProcEvent = 16

type ProcEvent struct {
	What      uint32
	Cpu       uint32
	Timestamp uint64
	Data      []byte
}

func (self *ProcEvent) ForkEvent() (*ForkEvent, error) {
	if self.What&PROC_EVENT_FORK == PROC_EVENT_FORK {
		return DeserializeForkEvent(self.Data)
	}
	return nil, fmt.Errorf("event type is not PROC_EVENT_FORK")
}

func (self *ProcEvent) ExecEvent() (*ExecEvent, error) {
	if self.What&PROC_EVENT_EXEC == PROC_EVENT_EXEC {
		return DeserializeExecEvent(self.Data)
	}
	return nil, fmt.Errorf("event type is not PROC_EVENT_EXEC")
}

func (self *ProcEvent) PtraceEvent() (*PtraceEvent, error) {
	if self.What&PROC_EVENT_PTRACE == PROC_EVENT_PTRACE {
		return DeserializePtraceEvent(self.Data)
	}
	return nil, fmt.Errorf("event type is not PROC_EVENT_PTRACE")
}

func (self *ProcEvent) CommEvent() (*CommEvent, error) {
	if self.What&PROC_EVENT_COMM == PROC_EVENT_COMM {
		return DeserializeCommEvent(self.Data)
	}
	return nil, fmt.Errorf("event type is not PROC_EVENT_COMM")
}

func (self *ProcEvent) CoredumpEvent() (*CoredumpEvent, error) {
	if self.What&PROC_EVENT_COREDUMP == PROC_EVENT_COREDUMP {
		return DeserializeCoredumpEvent(self.Data)
	}
	return nil, fmt.Errorf("event type is not PROC_EVENT_COMREDUMP")
}

func (self *ProcEvent) ExitEvent() (*ExitEvent, error) {
	if self.What&PROC_EVENT_EXIT == PROC_EVENT_EXIT {
		return DeserializeExitEvent(self.Data)
	}
	return nil, fmt.Errorf("event type is not PROC_EVENT_EXIT")
}

func DeserializeProcEvent(b []byte) (*ProcEvent, error) {
	if len(b) < SizeofProcEvent {
		return nil, fmt.Errorf("len(b) should be greater equal than %d", SizeofProcEvent)
	}

	native := netlinkconnector.NativeEndian()
	msg := new(ProcEvent)

	begin := 0
	msg.What = native.Uint32(b[begin:])
	begin += 4
	msg.Cpu = native.Uint32(b[begin:])
	begin += 4
	msg.Timestamp = native.Uint64(b[begin:])
	begin += 8
	msg.Data = b[begin:]
	return msg, nil
}
