package netlinkconnector

import "fmt"

const (
	SizeofCnMsg = 20
)

type CnMsg struct {
	Id     CbId
	Seq    uint32
	Ack    uint32
	Length uint16
	Flags  uint16
	Data   []byte
}

func (self *CnMsg) Len() int {
	return NlmAlignOf(SizeofCnMsg) + NlmAlignOf(int(self.Length))
}

func DeserializeCnMsg(b []byte) (*CnMsg, error) {
	if len(b) < SizeofCnMsg {
		return nil, fmt.Errorf("input bytes length should be greater equal than %d", SizeofCnMsg)
	}

	native := NativeEndian()
	msg := new(CnMsg)
	id, err := DeserializeCbId(b)
	if err != nil {
		return nil, err
	}
	msg.Id = *id
	b = b[NlmAlignOf(SizeofCbId):]
	msg.Seq = native.Uint32(b[0:4])
	msg.Ack = native.Uint32(b[4:8])
	msg.Length = native.Uint16(b[8:10])
	msg.Flags = native.Uint16(b[10:12])

	b = b[12:]
	if len(b) < int(msg.Length) {
		return nil, fmt.Errorf("length(%d) of Data should be greater eqaul than msg.Length(%d)", len(b), msg.Length)
	}

	msg.Data = b[:msg.Length]

	return msg, nil
}

func (self *CnMsg) Serialize() []byte {
	native := NativeEndian()
	buf := make([]byte, self.Len())

	id_buf := buf[:self.Id.Len()]
	msg_buf := buf[self.Id.Len():]

	copy(id_buf, self.Id.Serialize())

	native.PutUint32(msg_buf[0:4], self.Seq)
	native.PutUint32(msg_buf[4:8], self.Ack)
	native.PutUint16(msg_buf[8:10], self.Length)
	native.PutUint16(msg_buf[10:12], self.Flags)

	copy(buf[NlmAlignOf(SizeofCnMsg):], self.Data[:self.Length])
	return buf
}

func (self *CnMsg) AddData(data CnRequestData) {
	b := data.Serialize()
	self.Length = uint16(data.Len())
	self.Data = b[:self.Length]
}

type CnRequestData interface {
	Len() int
	Serialize() []byte
}
