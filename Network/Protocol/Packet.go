package Protocol

type Packet struct {
	PID  int
	Data []byte
}

var (
	LENGTH_HEADER      int = 4
	LENGTH_TYPE_PACKET int = 4
)
