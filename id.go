package netlinkconnector

import "unsafe"

const (
	SizeofCbId = 8
)

/*
 * idx and val are unique identifiers which
 * are used for message routing and
 * must be registered in connector.h for in-kernel usage.
 */
type CbId struct {
	Idx uint32
	Val uint32
}

func (self *CbId) Len() int {
	return NlmAlignOf(SizeofCbId)
}

func DeserializeCbId(b []byte) (*CbId, error) {
	return (*CbId)(unsafe.Pointer(&b[0:SizeofCbId][0])), nil
}

func (self *CbId) Serialize() []byte {
	native := NativeEndian()

	buf := make([]byte, self.Len())
	native.PutUint32(buf[0:4], self.Idx)
	native.PutUint32(buf[4:8], self.Val)
	return buf
}
