package cnproc

import (
	"bytes"
	"testing"
)

func TestMultiCastOpSerialize(t *testing.T) {
	testcases := []struct {
		in       *MultiCastOp
		expected []byte
	}{
		{
			&MultiCastOp{1},
			//for little endian
			[]byte{1, 0, 0, 0},
		},
		{
			&MultiCastOp{2},
			//for little endian
			[]byte{2, 0, 0, 0},
		},
	}

	for _, testcase := range testcases {
		actual := testcase.in.Serialize()
		if bytes.Equal(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}
