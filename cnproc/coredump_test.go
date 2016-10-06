package cnproc

import (
	"reflect"
	"testing"
)

func TestDeserializeCoredumpEvent(t *testing.T) {
	testcases := []struct {
		in       []byte
		expected *CoredumpEvent
	}{
		{
			//for little endian
			[]byte{
				1, 0, 0, 0,
				2, 0, 0, 0,
			},
			&CoredumpEvent{
				ProcessPid:  1,
				ProcessTgid: 2,
			},
		},
	}

	for _, testcase := range testcases {
		actual, err := DeserializeCoredumpEvent(testcase.in)
		if err != nil {
			t.Error(err.Error())
		} else if reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}
