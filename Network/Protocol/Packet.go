package Protocol

type Packet struct {
	length      int
	typeMessage int
	data        []byte
}
