package netlinkconnector

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDeserializeCnMsg(t *testing.T) {
	testcases := []struct {
		in       []byte
		expected *CnMsg
	}{
		{
			//for little endian
			[]byte{
				1, 0, 0, 0, //Id.Idx
				2, 0, 0, 0, //Id.Val
				4, 0, 0, 0, //Seq
				5, 0, 0, 0, //Ack
				4, 0, //Len
				7, 0, //Flags
				8, 0, 0, 0, //data uint32(8)
			},
			&CnMsg{
				Id:     CbId{1, 2},
				Seq:    4,
				Ack:    5,
				Length: 4,
				Flags:  7,
				Data:   []byte{8, 0, 0, 0},
			},
		},
	}

	for _, testcase := range testcases {
		actual, err := DeserializeCnMsg(testcase.in)
		if err != nil {
			t.Errorf("fail to Deserialize CnMsg")
		}

		if reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}
}

func TestCnMsgSerialize(t *testing.T) {
	testcases := []struct {
		in       *CnMsg
		expected []byte
	}{
		{
			&CnMsg{
				Id:     CbId{1, 2},
				Seq:    4,
				Ack:    5,
				Length: 4,
				Flags:  7,
				Data:   []byte{8, 0, 0, 0},
			},

			//for little endian
			[]byte{
				1, 0, 0, 0, //Id.Idx
				2, 0, 0, 0, //Id.Val
				4, 0, 0, 0, //Seq
				5, 0, 0, 0, //Ack
				4, 0, //Len
				7, 0, //Flags
				8, 0, 0, 0, //data uint32(8)
			},
		},
	}

	for _, testcase := range testcases {
		actual := testcase.in.Serialize()
		if bytes.Equal(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}

type AddDataTestType struct {
	Val uint32
}

func (self *AddDataTestType) Len() int {
	return 4
}

func (self *AddDataTestType) Serialize() []byte {
	b := make([]byte, self.Len())
	native := NativeEndian()
	native.PutUint32(b[0:4], self.Val)
	return b
}

func TestCnMsgAddData(t *testing.T) {
	testcases := []struct {
		in       *AddDataTestType
		expected *CnMsg
	}{
		{
			&AddDataTestType{
				Val: 1,
			},
			&CnMsg{
				Length: 4,
				Data:   []byte{1, 0, 0, 0},
			},
		},
	}

	for _, testcase := range testcases {
		actual := new(CnMsg)
		actual.AddData(testcase.in)
		if reflect.DeepEqual(actual, testcase.expected) == false {
			t.Errorf("expected\n%v\nbut\n%v", testcase.expected, actual)
		}
	}

}
