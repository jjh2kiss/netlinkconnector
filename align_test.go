package netlinkconnector

import (
	"syscall"
	"testing"
)

func TestNlmAlignOf(t *testing.T) {
	testcases := []struct {
		in       int
		expected int
	}{
		{syscall.NLMSG_ALIGNTO, syscall.NLMSG_ALIGNTO},
		{syscall.NLMSG_ALIGNTO - 1, syscall.NLMSG_ALIGNTO},
		{syscall.NLMSG_ALIGNTO * 2, syscall.NLMSG_ALIGNTO * 2},
		{syscall.NLMSG_ALIGNTO*2 - 1, syscall.NLMSG_ALIGNTO * 2},
		{syscall.NLMSG_ALIGNTO*2 + 1, syscall.NLMSG_ALIGNTO * 3},
	}

	for _, testcase := range testcases {
		actual := NlmAlignOf(testcase.in)
		if actual != testcase.expected {
			t.Errorf("expected %d but %d\n", testcase.expected, actual)
		}
	}
}
