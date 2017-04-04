package Protocol

import "math"

type PacketParser struct {
	packet *Packet
	offset int
}

func (parser *PacketParser) Decode(data []byte) {
	// read length
	bPacketType := data[:4]
	bData := data[0:4]
	parser.offset = 0
	parser.packet = &(Packet{PID: int(int(bPacketType[0]) + int(bPacketType[1])<<8 + int(bPacketType[2])<<16 + int(bPacketType[3])<<24), Data: bData})
}

func (parser *PacketParser) DecodePacket(packet *Packet) {
	// read length
	parser.offset = 0
	parser.packet = packet
}

func (parser *PacketParser) Encode() []byte {
	var msg []byte
	v := len(parser.packet.Data)
	v1 := parser.packet.PID
	b := make([]byte, 8)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff
	b[2] = byte(v>>8) & 0xff
	b[3] = byte(v>>24) & 0xff
	b[4] = byte(v1) & 0xff
	b[5] = byte(v1>>8) & 0xff
	b[6] = byte(v1>>8) & 0xff
	b[7] = byte(v1>>24) & 0xff
	msg = append(msg, b...)
	msg = append(msg, parser.packet.Data...)
	return msg
}

func (parser *PacketParser) Prepare(pid int) {
	parser.packet = &(Packet{PID: pid, Data: make([]byte, 0)})
	parser.offset = 0
}
func (parser *PacketParser) ReadByte() byte {
	ret := parser.packet.Data[parser.offset]
	parser.offset++
	return ret
}

func (parser *PacketParser) WriteByte(b byte) {
	parser.packet.Data = append(parser.packet.Data, b)
	parser.offset++
}

func (parser *PacketParser) ReadBool() bool {
	parser.offset++
	if parser.packet.Data[parser.offset-1] == 0 {
		return false
	} else {
		return true
	}

}

func (parser *PacketParser) WriteBool(b bool) {
	if !b {
		parser.packet.Data = append(parser.packet.Data, 0)
	} else {
		parser.packet.Data = append(parser.packet.Data, 1)
	}
	parser.offset++
}

// short type in c#
func (parser *PacketParser) ReadUShort() uint8 {
	b := parser.packet.Data[parser.offset : parser.offset+2]
	parser.offset += 2
	return uint8(uint8(b[0]) + uint8(b[1])<<8)
}
func (parser *PacketParser) WriteUShort(v uint8) {
	b := make([]byte, 2)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff

	parser.packet.Data = append(parser.packet.Data, b...)
	parser.offset += 2
}

func (parser *PacketParser) ReadShort() int8 {
	b := parser.packet.Data[parser.offset : parser.offset+2]
	parser.offset += 2
	return int8(int8(b[0]) + int8(b[1])<<8)
}

func (parser *PacketParser) WriteShort(v int8) {
	b := make([]byte, 2)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff

	parser.packet.Data = append(parser.packet.Data, b...)
	parser.offset += 2
}
func (parser *PacketParser) ReadUInt32() uint32 {
	b := parser.packet.Data[parser.offset : parser.offset+4]
	parser.offset += 4
	return uint32(uint32(b[0]) + uint32(b[1])<<8 + uint32(b[2])<<16 + uint32(b[3])<<24)
}

func (parser *PacketParser) WriteInt32(v int) {
	b := make([]byte, 4)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff
	b[2] = byte(v>>8) & 0xff
	b[3] = byte(v>>24) & 0xff
	parser.packet.Data = append(parser.packet.Data, b...)
	parser.offset += 4
}

func (parser *PacketParser) ReadInt32() int {
	b := parser.packet.Data[parser.offset : parser.offset+4]
	parser.offset += 4
	return int(int(b[0]) + int(b[1])<<8 + int(b[2])<<16 + int(b[3])<<24)
}

func (parser *PacketParser) WriteUInt32(v uint32) {
	b := make([]byte, 4)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff
	b[2] = byte(v>>8) & 0xff
	b[3] = byte(v>>24) & 0xff
	parser.packet.Data = append(parser.packet.Data, b...)
	parser.offset += 4
}

//long type in c#
func (parser *PacketParser) WriteLong(v int64) {
	b := make([]byte, 8)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff
	b[2] = byte(v>>8) & 0xff
	b[3] = byte(v>>24) & 0xff
	b[4] = byte(v>>32) & 0xff
	b[5] = byte(v>>40) & 0xff
	b[6] = byte(v>>48) & 0xff
	b[6] = byte(v>>56) & 0xff
	parser.packet.Data = append(parser.packet.Data, b...)
	parser.offset += 8
}

func (parser *PacketParser) ReadLong() int64 {
	b := parser.packet.Data[parser.offset : parser.offset+8]
	parser.offset += 8
	return int64(int64(b[0]) + int64(b[1])<<8 + int64(b[2])<<16 + int64(b[3])<<24 + int64(b[0])<<32 + int64(b[1])<<40 + int64(b[2])<<48 + int64(b[3])<<56)

}
func (parser *PacketParser) WriteULong(v uint64) {
	b := make([]byte, 8)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff
	b[2] = byte(v>>8) & 0xff
	b[3] = byte(v>>24) & 0xff
	b[4] = byte(v>>32) & 0xff
	b[5] = byte(v>>40) & 0xff
	b[6] = byte(v>>48) & 0xff
	b[6] = byte(v>>56) & 0xff
	parser.packet.Data = append(parser.packet.Data, b...)
	parser.offset += 8
}

func (parser *PacketParser) ReadULong() uint64 {
	b := parser.packet.Data[parser.offset : parser.offset+8]
	parser.offset += 8
	return uint64(uint64(b[0]) + uint64(b[1])<<8 + uint64(b[2])<<16 + uint64(b[3])<<24 + uint64(b[0])<<32 + uint64(b[1])<<40 + uint64(b[2])<<48 + uint64(b[3])<<56)
}

//float type in c#
func (parser *PacketParser) WriteFloat32(t float32) {
	v := math.Float32bits(t)
	b := make([]byte, 4)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff
	b[2] = byte(v>>8) & 0xff
	b[3] = byte(v>>24) & 0xff
	parser.packet.Data = append(parser.packet.Data, b...)
	parser.offset += 4

}

func (parser *PacketParser) ReadFloat32() float32 {
	b := parser.packet.Data[parser.offset : parser.offset+8]
	parser.offset += 4
	v := uint32(uint32(b[0]) + uint32(b[1])<<8 + uint32(b[2])<<16 + uint32(b[3])<<24)
	return math.Float32frombits(v)
}

//Double type in c#
func (parser *PacketParser) WriteFloat64(t float64) {
	v := math.Float64bits(t)
	b := make([]byte, 8)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff
	b[2] = byte(v>>8) & 0xff
	b[3] = byte(v>>24) & 0xff
	b[4] = byte(v>>32) & 0xff
	b[5] = byte(v>>40) & 0xff
	b[6] = byte(v>>48) & 0xff
	b[6] = byte(v>>56) & 0xff
	parser.packet.Data = append(parser.packet.Data, b...)
	parser.offset += 8

}

func (parser *PacketParser) ReadFloat64() float64 {
	b := parser.packet.Data[parser.offset : parser.offset+8]
	parser.offset += 8
	v := uint64(uint64(b[0]) + uint64(b[1])<<8 + uint64(b[2])<<16 + uint64(b[3])<<24 + uint64(b[0])<<32 + uint64(b[1])<<40 + uint64(b[2])<<48 + uint64(b[3])<<56)
	return math.Float64frombits(v)
}

func (parser *PacketParser) WriteString(s string) {
	b := []byte(s)
	parser.WriteInt32((len(b)))
	parser.packet.Data = append(parser.packet.Data, b...)
	parser.offset += len(b)
}

func (parser *PacketParser) ReadString() string {
	lenStr := parser.ReadInt32()
	ret := string(parser.packet.Data[parser.offset : parser.offset+lenStr])
	parser.offset = parser.offset + 4 + lenStr
	return ret
}
