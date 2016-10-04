package netlinkconnector

import (
	"fmt"
	"os"
	"syscall"

	"github.com/vishvananda/netlink/nl"
)

type NetlinkConnector struct {
	fd  int //File descriptor for Netlink Socket
	seq uint32
	lsa *syscall.SockaddrNetlink
}

// Get a new Handler for Netlink connector
func NewNetlinkConnector(idx uint32) (*NetlinkConnector, error) {
	sa := &syscall.SockaddrNetlink{
		Family: syscall.AF_NETLINK,
		Pid:    uint32(os.Getpid()),
		Groups: idx,
	}

	fd, err := syscall.Socket(
		syscall.AF_NETLINK,
		syscall.SOCK_RAW,
		syscall.NETLINK_CONNECTOR,
	)

	if err != nil {
		return nil, err
	}

	socket := new(NetlinkConnector)
	socket.fd = fd
	socket.lsa = sa

	if err := syscall.Bind(fd, sa); err != nil {
		syscall.Close(fd)
		return nil, err
	}
	return socket, nil
}

//Close Netlink Connector
func (self *NetlinkConnector) Close() {
	if self.fd >= 0 {
		syscall.Close(self.fd)
		self.fd = -1
	}
}

//Return a file descriptor
func (self *NetlinkConnector) GetFd() int {
	return self.fd
}

//Write Netlink Connector Message
func (self *NetlinkConnector) Write(cnmsg *CnMsg) error {
	if self.fd < 0 {
		return fmt.Errorf("write called on a closed socket")
	}

	request := &nl.NetlinkRequest{}
	request.Pid = uint32(os.Getpid())
	request.Type = syscall.NLMSG_DONE
	request.Seq = self.seq
	request.AddData(cnmsg)

	self.seq++

	_, err := syscall.Write(self.fd, request.Serialize())
	if err != nil {
		return err
	}
	return nil
}

//Returns the slice containing the CnMsg(Netlink Connector Message) structures
func (self *NetlinkConnector) Read() ([]*CnMsg, error) {
	if self.fd < 0 {
		return nil, fmt.Errorf("Read called on a closed socket")
	}

	p := make([]byte, syscall.Getpagesize())
	n, _, err := syscall.Recvfrom(self.fd, p, 0)
	if err != nil {
		return nil, err
	}

	if n < syscall.NLMSG_HDRLEN {
		return nil, fmt.Errorf("Got short netlink message from socket")
	}

	cnmsgs, err := ParseNetlinkConnectorMessage(p)
	if err != nil {
		return nil, err
	}

	return cnmsgs, nil
}

//ParseNetlinkConnectorMessage parses b as an array of netlink connector messages and
// returns the slice containing the CnMsg structures.
func ParseNetlinkConnectorMessage(b []byte) ([]*CnMsg, error) {
	msgs, err := syscall.ParseNetlinkMessage(b)
	if err != nil {
		return nil, err
	}

	cnmsgs := make([]*CnMsg, 0, len(msgs))
	for _, msg := range msgs {
		if msg.Header.Type == syscall.NLMSG_ERROR || msg.Header.Type == syscall.NLMSG_NOOP {
			continue
		}

		cnmsg, err := DeserializeCnMsg(msg.Data)
		if err != nil {
			continue
		}
		cnmsgs = append(cnmsgs, cnmsg)
	}
	return cnmsgs, nil

}
