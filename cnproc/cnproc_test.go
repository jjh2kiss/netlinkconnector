package cnproc

import (
	"reflect"
	"testing"
)

func TestDeserializeProcEvent(t *testing.T) {
	testcases := []struct {
		in       []byte
		expected *ProcEvent
	}{
		{
			//for little endian
			[]byte{
				1, 0, 0, 0, //Id.Idx
				2, 0, 0, 0, //Id.Val
				3, 0, 0, 0, 0, 0, 0, 0, //Ack
				4, 0, 0, 0, //data uint32(8)
			},
			&ProcEvent{
				What:      1,
				Cpu:       2,
				Timestamp: 3,
				Data:      []byte{4, 0, 0, 0},
			},
		},
	}

	for _, testcase := range testcases {
		actual, err := DeserializeProcEvent(testcase.in)
		if err != nil || reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}

func TestProcEventForkEvent(t *testing.T) {
	testcases := []struct {
		in       *ProcEvent
		expected *ForkEvent
	}{
		{
			&ProcEvent{
				What:      PROC_EVENT_FORK,
				Cpu:       2,
				Timestamp: 3,
				Data: []byte{
					1, 0, 0, 0,
					2, 0, 0, 0,
					3, 0, 0, 0,
					4, 0, 0, 0,
				},
			},

			&ForkEvent{
				ParentPid:  1,
				ParentTgid: 2,
				ChildPid:   3,
				ChildTgid:  4,
			},
		},
	}

	for _, testcase := range testcases {
		actual, err := testcase.in.ForkEvent()
		if err != nil || reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}

func TestProcEventForkEventInvalid(t *testing.T) {
	testcases := []struct {
		in *ProcEvent
	}{
		{
			&ProcEvent{
				What:      PROC_EVENT_NONE,
				Cpu:       2,
				Timestamp: 3,
				Data: []byte{
					1, 0, 0, 0,
					2, 0, 0, 0,
					3, 0, 0, 0,
					4, 0, 0, 0,
				},
			},
		},
	}

	for _, testcase := range testcases {
		_, err := testcase.in.ForkEvent()
		if err == nil {
			t.Errorf("expected nil but not %s", err.Error())
		}
	}

}

func TestProcEventExecEvent(t *testing.T) {
	testcases := []struct {
		in       *ProcEvent
		expected *ExecEvent
	}{
		{
			&ProcEvent{
				What:      PROC_EVENT_EXEC,
				Cpu:       2,
				Timestamp: 3,
				Data: []byte{
					1, 0, 0, 0,
					2, 0, 0, 0,
				},
			},

			&ExecEvent{
				ProcessPid:  1,
				ProcessTgid: 2,
			},
		},
	}

	for _, testcase := range testcases {
		actual, err := testcase.in.ExecEvent()
		if err != nil || reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}

func TestProcEventExecEventInvalid(t *testing.T) {
	testcases := []struct {
		in *ProcEvent
	}{
		{
			&ProcEvent{
				What:      PROC_EVENT_NONE,
				Cpu:       2,
				Timestamp: 3,
				Data: []byte{
					1, 0, 0, 0,
					2, 0, 0, 0,
					3, 0, 0, 0,
					4, 0, 0, 0,
				},
			},
		},
	}

	for _, testcase := range testcases {
		_, err := testcase.in.ExecEvent()
		if err == nil {
			t.Errorf("expected nil but not %s", err.Error())
		}
	}

}

func TestProcEventPtraceEvent(t *testing.T) {
	testcases := []struct {
		in       *ProcEvent
		expected *PtraceEvent
	}{
		{
			&ProcEvent{
				What:      PROC_EVENT_PTRACE,
				Cpu:       2,
				Timestamp: 3,
				Data: []byte{
					1, 0, 0, 0,
					2, 0, 0, 0,
					3, 0, 0, 0,
					4, 0, 0, 0,
				},
			},

			&PtraceEvent{
				ProcessPid:  1,
				ProcessTgid: 2,
				TracerPid:   3,
				TracerTgid:  4,
			},
		},
	}

	for _, testcase := range testcases {
		actual, err := testcase.in.PtraceEvent()
		if err != nil || reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}

func TestProcEventPtraceEventInvalid(t *testing.T) {
	testcases := []struct {
		in *ProcEvent
	}{
		{
			&ProcEvent{
				What:      PROC_EVENT_NONE,
				Cpu:       2,
				Timestamp: 3,
				Data: []byte{
					1, 0, 0, 0,
					2, 0, 0, 0,
					3, 0, 0, 0,
					4, 0, 0, 0,
				},
			},
		},
	}

	for _, testcase := range testcases {
		_, err := testcase.in.PtraceEvent()
		if err == nil {
			t.Errorf("expected nil but not %s", err.Error())
		}
	}

}

func TestProcEventCommEvent(t *testing.T) {
	testcases := []struct {
		in       *ProcEvent
		expected *CommEvent
	}{
		{
			&ProcEvent{
				What:      PROC_EVENT_COMM,
				Cpu:       2,
				Timestamp: 3,
				Data: []byte{
					1, 0, 0, 0,
					2, 0, 0, 0,
					97, 98, 99, 100,
					49, 50, 51, 52,
					97, 98, 99, 100,
					49, 50, 51, 52,
				},
			},

			&CommEvent{
				ProcessPid:  1,
				ProcessTgid: 2,
				Comm: [16]byte{
					97, 98, 99, 100,
					49, 50, 51, 52,
					97, 98, 99, 100,
					49, 50, 51, 52,
				},
			},
		},
	}

	for _, testcase := range testcases {
		actual, err := testcase.in.CommEvent()
		if err != nil || reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}

func TestProcEventCommEventInvalid(t *testing.T) {
	testcases := []struct {
		in *ProcEvent
	}{
		{
			&ProcEvent{
				What:      PROC_EVENT_NONE,
				Cpu:       2,
				Timestamp: 3,
				Data: []byte{
					1, 0, 0, 0,
					2, 0, 0, 0,
					3, 0, 0, 0,
					4, 0, 0, 0,
				},
			},
		},
	}

	for _, testcase := range testcases {
		_, err := testcase.in.CommEvent()
		if err == nil {
			t.Errorf("expected nil but not %s", err.Error())
		}
	}

}

func TestProcEventCoredumpEvent(t *testing.T) {
	testcases := []struct {
		in       *ProcEvent
		expected *CoredumpEvent
	}{
		{
			&ProcEvent{
				What:      PROC_EVENT_COREDUMP,
				Cpu:       2,
				Timestamp: 3,
				Data: []byte{
					1, 0, 0, 0,
					2, 0, 0, 0,
				},
			},

			&CoredumpEvent{
				ProcessPid:  1,
				ProcessTgid: 2,
			},
		},
	}

	for _, testcase := range testcases {
		actual, err := testcase.in.CoredumpEvent()
		if err != nil || reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}

func TestProcEventCoredumpEventInvalid(t *testing.T) {
	testcases := []struct {
		in *ProcEvent
	}{
		{
			&ProcEvent{
				What:      PROC_EVENT_NONE,
				Cpu:       2,
				Timestamp: 3,
				Data: []byte{
					1, 0, 0, 0,
					2, 0, 0, 0,
					3, 0, 0, 0,
					4, 0, 0, 0,
				},
			},
		},
	}

	for _, testcase := range testcases {
		_, err := testcase.in.CoredumpEvent()
		if err == nil {
			t.Errorf("expected nil but not %s", err.Error())
		}
	}

}

func TestProcEventExitEvent(t *testing.T) {
	testcases := []struct {
		in       *ProcEvent
		expected *ExitEvent
	}{
		{
			&ProcEvent{
				What:      PROC_EVENT_EXIT,
				Cpu:       2,
				Timestamp: 3,
				Data: []byte{
					1, 0, 0, 0,
					2, 0, 0, 0,
					3, 0, 0, 0,
					4, 0, 0, 0,
				},
			},

			&ExitEvent{
				ProcessPid:  1,
				ProcessTgid: 2,
				ExitCode:    3,
				ExitSignal:  4,
			},
		},
	}

	for i, testcase := range testcases {
		actual, err := testcase.in.ExitEvent()
		if err != nil || reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("%d. expected\n%v\nbut\n%v", i, testcase.expected, actual)
		}
	}

}

func TestProcEventExitEventInvalid(t *testing.T) {
	testcases := []struct {
		in *ProcEvent
	}{
		{
			&ProcEvent{
				What:      PROC_EVENT_NONE,
				Cpu:       2,
				Timestamp: 3,
				Data: []byte{
					1, 0, 0, 0,
					2, 0, 0, 0,
					3, 0, 0, 0,
					4, 0, 0, 0,
				},
			},
		},
	}

	for _, testcase := range testcases {
		_, err := testcase.in.ExitEvent()
		if err == nil {
			t.Errorf("expected nil but not %s", err.Error())
		}
	}

}
