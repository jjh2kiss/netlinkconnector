package cnproc

import (
	"reflect"
	"testing"
)

func TestDeserializeExitEvent(t *testing.T) {
	testcases := []struct {
		in       []byte
		expected *ExitEvent
	}{
		{
			//for little endian
			[]byte{
				1, 0, 0, 0,
				2, 0, 0, 0,
				3, 0, 0, 0,
				4, 0, 0, 0,
			},
			&ExitEvent{
				ProcessPid:  1,
				ProcessTgid: 2,
				ExitCode:    3,
				ExitSignal:  4,
			},
		},
	}

	for _, testcase := range testcases {
		actual, err := DeserializeExitEvent(testcase.in)
		if err != nil {
			t.Errorf(err.Error())
		} else if reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}
