package netlinkconnector

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDeserializeCbId(t *testing.T) {
	testcases := []struct {
		in       []byte
		expected *CbId
	}{
		{
			//for little endian
			[]byte{1, 0, 0, 0, 2, 0, 0, 0},
			&CbId{1, 2},
		},
		{
			//for little endian
			[]byte{255, 0, 0, 0, 254, 0, 0, 0},
			&CbId{255, 254},
		},
	}

	for _, testcase := range testcases {
		actual, err := DeserializeCbId(testcase.in)
		if err != nil {
			t.Errorf("fail to descrizeCbId")
		}
		if reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}
}

func TestCbIdSerialize(t *testing.T) {
	testcases := []struct {
		in       *CbId
		expected []byte
	}{
		{
			&CbId{1, 2},
			//for little endian
			[]byte{1, 0, 0, 0, 2, 0, 0, 0},
		},
		{
			&CbId{255, 254},
			//for little endian
			[]byte{255, 0, 0, 0, 254, 0, 0, 0},
		},
	}

	for _, testcase := range testcases {
		actual := testcase.in.Serialize()
		if bytes.Equal(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}
