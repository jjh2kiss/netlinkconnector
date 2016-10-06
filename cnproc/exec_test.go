package cnproc

import (
	"reflect"
	"testing"
)

func TestDeserializeExecEvent(t *testing.T) {
	testcases := []struct {
		in       []byte
		expected *ExecEvent
	}{
		{
			//for little endian
			[]byte{
				1, 0, 0, 0,
				2, 0, 0, 0,
			},
			&ExecEvent{
				ProcessPid:  1,
				ProcessTgid: 2,
			},
		},
	}

	for _, testcase := range testcases {
		actual, err := DeserializeExecEvent(testcase.in)
		if err != nil {
			t.Error(err.Error())
		} else if reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}
