package netlinkconnector

import (
	"os"
	"reflect"
	"testing"

	"github.com/vishvananda/netlink/nl"
)

//it need root privilege
func TestNewNetlinkConnector(t *testing.T) {
	testcases := []struct {
		idx uint32
	}{
		{CN_IDX_PROC},
		{CN_IDX_CIFS},
		{CN_IDX_DM},
		{CN_KVP_IDX},
		{CN_VSS_IDX},
	}

	for _, testcase := range testcases {
		nlcn, err := NewNetlinkConnector(testcase.idx)

		if err != nil {
			t.Errorf("expected err is nil, but not nil(%s)", err.Error())
		} else {
			nlcn.Close()
		}
	}
}

//it need root privilege
func TestNetlinkConnectorGetFd(t *testing.T) {
	testcases := []struct {
		nlcn     NetlinkConnector
		expected int
	}{
		{NetlinkConnector{fd: 0}, 0},
		{NetlinkConnector{fd: 1}, 1},
		{NetlinkConnector{fd: 2}, 2},
	}

	for _, testcase := range testcases {
		actual := testcase.nlcn.GetFd()
		if actual != testcase.expected {
			t.Errorf("expected %d, but %d", testcase.expected, actual)
		}
	}
}

//it need root privilege
func TestNetlinkConnectorClose(t *testing.T) {
	nlcn, err := NewNetlinkConnector(CN_IDX_PROC)
	if err != nil {
		t.Errorf("Fail to test function Close(%s)", err.Error())
		return
	}

	before := nlcn.fd
	nlcn.Close()
	after := nlcn.fd

	if before == after {
		t.Errorf(
			"expected not equal between before(%d) and after(%d), but equal",
			before,
			after,
		)
	}
}

func TestParseNetlinkConnectorMessage(t *testing.T) {
	cnmsg := CnMsg{
		Id:     CbId{1, 2},
		Length: 10,
		Seq:    1,
		Ack:    2,
		Flags:  1,
		Data:   []byte("abcde12345"),
	}

	req := nl.NetlinkRequest{}
	req.Type = 1
	req.Flags = 2
	req.Seq = 3
	req.Pid = uint32(os.Getpid())
	req.AddData(&cnmsg)

	breq := req.Serialize()
	actual_msgs, err := ParseNetlinkConnectorMessage(breq)
	if err != nil {
		t.Errorf("fail to parse netlink connector message, %s", err.Error())
		return
	}

	for _, actual_msg := range actual_msgs {
		if reflect.DeepEqual(*actual_msg, cnmsg) {
			t.Errorf("expected equal, but not equal")
		}
	}

}

func TestParseNetlinkConnectorMessages(t *testing.T) {
	cnmsg := CnMsg{
		Id:     CbId{1, 2},
		Length: 10,
		Seq:    1,
		Ack:    2,
		Flags:  1,
		Data:   []byte("abcde12345"),
	}

	req := nl.NetlinkRequest{}
	req.Type = 1
	req.Flags = 2
	req.Seq = 3
	req.Pid = uint32(os.Getpid())
	req.AddData(&cnmsg)

	breq := req.Serialize()
	breqs := make([]byte, len(breq)*2)
	copy(breqs[:len(breq)], breq)
	copy(breqs[len(breq):], breq)
	actual_msgs, err := ParseNetlinkConnectorMessage(breqs)
	if err != nil {
		t.Errorf("fail to parse netlink connector messages, %s", err.Error())
		return
	}

	for _, actual_msg := range actual_msgs {
		if reflect.DeepEqual(*actual_msg, cnmsg) {
			t.Errorf("expected equal, but not equal")
		}
	}

}
