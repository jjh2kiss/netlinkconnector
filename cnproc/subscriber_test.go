package cnproc

import (
	"os/exec"
	"testing"

	"github.com/jjh2kiss/netlinkconnector"
)

func TestNewSubscriber(t *testing.T) {
	subscriber, err := NewSubscriber()
	defer subscriber.Close()

	if err != nil {
		t.Errorf("Fail to make new subscriber(%s)", err.Error())
	}
}

func TestEnableMcastListen(t *testing.T) {
	connector, err := netlinkconnector.NewNetlinkConnector(netlinkconnector.CN_IDX_PROC)
	if err != nil {
		t.Errorf("Fail to make new Connector(%s)", err.Error())
		return
	}
	defer connector.Close()

	err = enableMcastListen(connector)
	if err != nil {
		t.Errorf("Fail to Enable McastListen(%s)", err.Error())
		return
	}
}

func TestDisableMcastListen(t *testing.T) {
	connector, err := netlinkconnector.NewNetlinkConnector(netlinkconnector.CN_IDX_PROC)
	if err != nil {
		t.Errorf("Fail to make new Connector(%s)", err.Error())
		return
	}
	defer connector.Close()

	err = disableMcastListen(connector)
	if err != nil {
		t.Errorf("Fail to Disable McastListen(%s)", err.Error())
		return
	}
}

func TestSetProcCnMcastOptionListen(t *testing.T) {
	connector, err := netlinkconnector.NewNetlinkConnector(netlinkconnector.CN_IDX_PROC)
	if err != nil {
		t.Errorf("Fail to make new Connector(%s)", err.Error())
		return
	}
	defer connector.Close()

	err = setProcCnMcastOption(connector, PROC_CN_MCAST_LISTEN)
	if err != nil {
		t.Errorf("Fail to set PROC_CN_MCAST_LISTEN, (%s)", err.Error())
		return
	}

}

func TestSetProcCnMcastOptionIgnore(t *testing.T) {
	connector, err := netlinkconnector.NewNetlinkConnector(netlinkconnector.CN_IDX_PROC)
	if err != nil {
		t.Errorf("Fail to make new Connector(%s)", err.Error())
		return
	}
	defer connector.Close()

	err = setProcCnMcastOption(connector, PROC_CN_MCAST_IGNORE)
	if err != nil {
		t.Errorf("Fail to set PROC_CN_MCAST_IGNORE, (%s)", err.Error())
		return
	}

}

func TestSubscriberRead(t *testing.T) {
	subscriber, err := NewSubscriber()
	defer subscriber.Close()

	if err != nil {
		t.Errorf("Fail to make new subscriber(%s)", err.Error())
		return
	}

	go func() {
		for i := 0; i < 2; i++ {
			_, _ = exec.Command("sleep", "0.3").Output()
		}
	}()

	events, err := subscriber.Read()
	if err != nil {
		t.Errorf("Fail to read events from subscriber(%s)", err.Error())
		return
	}

	if len(events) == 0 {
		t.Errorf("Fail to read event, empty events")
		return
	}

}

func TestSubscribe(t *testing.T) {
	subscriber, err := NewSubscriber()
	defer subscriber.Close()

	if err != nil {
		t.Errorf("Fail to make new subscriber(%s)", err.Error())
		return
	}

	go func() {
		for i := 0; i < 2; i++ {
			_, _ = exec.Command("sleep", "0.3").Output()
		}
		subscriber.Close()
	}()

	ch, err := subscriber.Subscribe()
	if err != nil {
		t.Errorf("Fail to get subscriber channel")
		return
	}

	exec_count := 0
	for event := range ch {
		if event.What == PROC_EVENT_EXEC {
			exec_count++
		}
	}

	if exec_count == 0 {
		t.Errorf("Fail to read event from the channel")
	}
}
