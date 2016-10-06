package cnproc

import (
	"fmt"
	"os"

	"github.com/jjh2kiss/netlinkconnector"
)

type Subscriber struct {
	Connector *netlinkconnector.NetlinkConnector
}

// NewSubscriber make a new subscriber for process event
func NewSubscriber() (*Subscriber, error) {
	if os.Geteuid() != 0 {
		return nil, fmt.Errorf("Need to run with root privilege")
	}

	connector, err := netlinkconnector.NewNetlinkConnector(netlinkconnector.CN_IDX_PROC)
	if err != nil {
		return nil, err
	}

	if err := enableMcastListen(connector); err != nil {
		connector.Close()
		return nil, err
	}

	subscriber := new(Subscriber)
	subscriber.Connector = connector

	return subscriber, nil
}

func enableMcastListen(connector *netlinkconnector.NetlinkConnector) error {
	return setProcCnMcastOption(connector, PROC_CN_MCAST_LISTEN)
}

func disableMcastListen(connector *netlinkconnector.NetlinkConnector) error {
	return setProcCnMcastOption(connector, PROC_CN_MCAST_IGNORE)
}

func setProcCnMcastOption(connector *netlinkconnector.NetlinkConnector, option uint32) error {
	op := &MultiCastOp{option}

	cn_msg := new(netlinkconnector.CnMsg)
	cn_msg.Id.Val = netlinkconnector.CN_VAL_PROC
	cn_msg.Id.Idx = netlinkconnector.CN_IDX_PROC
	cn_msg.AddData(op)

	err := connector.Write(cn_msg)
	if err != nil {
		return err
	}
	return nil
}

// Close close subscriber
func (self *Subscriber) Close() {
	if self != nil {
		self.Connector.Close()
	}
}

// Read returns Process event information, such as EXEC, EXIT, FORK, PTRACE, ...
func (self *Subscriber) Read() ([]*ProcEvent, error) {
	cnmsgs, err := self.Connector.Read()
	if err != nil {
		return nil, err
	}

	events := make([]*ProcEvent, 0, len(cnmsgs))
	for _, cnmsg := range cnmsgs {
		event, err := DeserializeProcEvent(cnmsg.Data)
		if err != nil {
			continue
		}

		if cnmsg.Id.Idx != netlinkconnector.CN_IDX_PROC {
			continue
		}

		if cnmsg.Id.Val != netlinkconnector.CN_VAL_PROC {
			continue
		}

		events = append(events, event)
	}

	return events, nil
}

// Subscribe returns a channel to receive process event
func (self *Subscriber) Subscribe() (<-chan *ProcEvent, error) {
	ch := make(chan *ProcEvent, 100)

	go func() {
		defer close(ch)

		for {
			events, err := self.Read()
			if err != nil {
				return
			}

			for _, event := range events {
				ch <- event
			}
		}
	}()

	return ch, nil
}
