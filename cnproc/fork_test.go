package cnproc

import (
	"reflect"
	"testing"
)

func TestDeserializeForkEvent(t *testing.T) {
	testcases := []struct {
		in       []byte
		expected *ForkEvent
	}{
		{
			//for little endian
			[]byte{
				1, 0, 0, 0,
				2, 0, 0, 0,
				3, 0, 0, 0,
				4, 0, 0, 0,
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
		actual, err := DeserializeForkEvent(testcase.in)
		if err != nil {
			t.Errorf(err.Error())
		} else if reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}
