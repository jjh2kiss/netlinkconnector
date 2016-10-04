package netlinkconnector

import (
	"encoding/binary"

	"github.com/vishvananda/netlink/nl"
)

// Get native endianness for the system
func NativeEndian() binary.ByteOrder {
	return nl.NativeEndian()
}
