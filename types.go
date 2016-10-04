package netlinkconnector

/*
 * idx and val are unique identifiers which
 * are used for message routing and
 * must be registered in connector.h for in-kernel usage.
 * get from /usr/include/connector.h
 */
const (
	CN_IDX_PROC            = 0x1
	CN_VAL_PROC            = 0x1
	CN_IDX_CIFS            = 0x2
	CN_VAL_CIFS            = 0x1
	CN_W1_IDX              = 0x3 /* w1 communication */
	CN_W1_VAL              = 0x1
	CN_IDX_V86D            = 0x4
	CN_VAL_V86D_UVESAF     = 0x1
	CN_IDX_BB              = 0x5 /* BlackBoard, from the TSP GPL sampling framework */
	CN_DST_IDX             = 0x6
	CN_DST_VAL             = 0x1
	CN_IDX_DM              = 0x7 /* Device Mapper */
	CN_VAL_DM_USERSPACE_LO = 0x1
	CN_IDX_DRBD            = 0x8
	CN_VAL_DRBD            = 0x1
	CN_KVP_IDX             = 0x9 /* HyperV KVP */
	CN_KVP_VAL             = 0x1 /* queries from the kernel */
	CN_VSS_IDX             = 0xA /* HyperV VSS */
	CN_VSS_VAL             = 0x1 /* queries from the kernel */

	CN_NETLINK_USERS = 11

	CONNECTOR_MAX_MSG_SIZE = 16384
)
