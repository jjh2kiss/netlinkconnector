package cnproc

import (
	"reflect"
	"testing"
)

func TestDeserializePtraceEvent(t *testing.T) {
	testcases := []struct {
		in       []byte
		expected *PtraceEvent
	}{
		{
			//for little endian
			[]byte{
				1, 0, 0, 0,
				2, 0, 0, 0,
				3, 0, 0, 0,
				4, 0, 0, 0,
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
		actual, err := DeserializePtraceEvent(testcase.in)
		if err != nil {
			t.Error(err.Error())
		} else if reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}
