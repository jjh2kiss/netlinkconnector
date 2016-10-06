package cnproc

import (
	"reflect"
	"testing"
)

func TestDeserializeCommEvent(t *testing.T) {
	testcases := []struct {
		in       []byte
		expected *CommEvent
	}{
		{
			//for little endian
			[]byte{
				1, 0, 0, 0,
				2, 0, 0, 0,
				97, 98, 99, 100,
				49, 50, 51, 52,
				97, 98, 99, 100,
				49, 50, 51, 52,
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
		actual, err := DeserializeCommEvent(testcase.in)
		if err != nil || reflect.DeepEqual(actual, testcase.expected) == false {
			if err != nil {
				t.Errorf(err.Error())
			} else {
				t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
			}
		}
	}

}
