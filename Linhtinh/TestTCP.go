package main

import (
	"fmt"
	"math"
	"net"
	"os"
)

var delimPacket byte = '|'

func main() {
	/*	addr, err := net.ResolveTCPAddr("tcp4", ":8080")
		checkError(err)
		listener, err := net.ListenTCP("tcp", addr)
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("cannot resolve client")
				continue
			}
			go handleConnection(conn)
		}*/
	var a uint32 = 8
	packet := Packet{}
	packet.WriteUInt32(a)

	packet.WriteString("hohoho")

	packet.offset = 0
	ret := packet.ReadUInt32()

	retStr := packet.ReadString()
	fmt.Println(ret)
	fmt.Println(retStr)

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	/*var buf [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}

		}

	}
	s := string(result.Bytes()[:])
	st := strings.Split(s, s[:1])
	for _, v := range st {
		fmt.Println(v)
	}
	*/
	/*reader := bufio.NewReader(conn)
	for {
		mess, err := reader.ReadBytes()
		if err != nil {
			break
		}
		if len(mess) > 0 {
			fmt.Println(mess)
		}

	}*/
}

type Packet struct {
	data   []byte
	offset int
}

func (packet *Packet) ReadByte() byte {
	packet.offset++
	return packet.data[packet.offset]
}

func (packet *Packet) WriteByte(b byte) {
	packet.data = append(packet.data, b)
	packet.offset++
}

func (packet *Packet) ReadBool() bool {
	packet.offset++
	if packet.data[packet.offset] == 0 {
		return false
	} else {
		return true
	}

}

func (packet *Packet) WriteBool(b bool) {
	if !b {
		packet.data = append(packet.data, 0)
	} else {
		packet.data = append(packet.data, 1)
	}
	packet.offset++
}

// short type in c#
func (packet *Packet) ReadUShort() uint8 {
	b := packet.data[packet.offset : packet.offset+2]
	packet.offset += 2
	return uint8(uint8(b[0]) + uint8(b[1])<<8)
}
func (packet *Packet) WriteUShort(v uint8) {
	b := make([]byte, 2)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff

	packet.data = append(packet.data, b...)
	packet.offset += 2
}

func (packet *Packet) ReadShort() int8 {
	b := packet.data[packet.offset : packet.offset+2]
	packet.offset += 2
	return int8(int8(b[0]) + int8(b[1])<<8)
}

func (packet *Packet) WriteShort(v int8) {
	b := make([]byte, 2)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff

	packet.data = append(packet.data, b...)
	packet.offset += 2
}
func (packet *Packet) ReadUInt32() uint32 {
	b := packet.data[packet.offset : packet.offset+4]
	packet.offset += 4
	return uint32(uint32(b[0]) + uint32(b[1])<<8 + uint32(b[2])<<16 + uint32(b[3])<<24)
}

func (packet *Packet) WriteInt32(v int) {
	b := make([]byte, 4)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff
	b[2] = byte(v>>8) & 0xff
	b[3] = byte(v>>24) & 0xff
	/*
		b[4] = byte(v>>32) & 0xff
		b[5] = byte(v>>40) & 0xff
		b[6] = byte(v>>48) & 0xff
		b[6] = byte(v>>56) & 0xff
	*/
	packet.data = append(packet.data, b...)
	packet.offset += 4
}

func (packet *Packet) ReadInt32() int {
	b := packet.data[packet.offset : packet.offset+4]
	packet.offset += 4
	return int(int(b[0]) + int(b[1])<<8 + int(b[2])<<16 + int(b[3])<<24)
	//	return int(int(b[0]) + int(b[1])<<8 + int(b[2])<<16 + int(b[3])<<24 + int(b[0])<<32 + int(b[1])<<40 + int(b[2])<<48 + int(b[3])<<56)
}

func (packet *Packet) WriteUInt32(v uint32) {
	b := make([]byte, 4)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff
	b[2] = byte(v>>8) & 0xff
	b[3] = byte(v>>24) & 0xff
	packet.data = append(packet.data, b...)
	packet.offset += 4
}

//long type in c#
func (packet *Packet) WriteLong(v int64) {
	b := make([]byte, 8)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff
	b[2] = byte(v>>8) & 0xff
	b[3] = byte(v>>24) & 0xff
	b[4] = byte(v>>32) & 0xff
	b[5] = byte(v>>40) & 0xff
	b[6] = byte(v>>48) & 0xff
	b[6] = byte(v>>56) & 0xff
	packet.data = append(packet.data, b...)
	packet.offset += 8
}

func (packet *Packet) ReadLong() int64 {
	b := packet.data[packet.offset : packet.offset+8]
	packet.offset += 8
	return int64(int64(b[0]) + int64(b[1])<<8 + int64(b[2])<<16 + int64(b[3])<<24 + int64(b[0])<<32 + int64(b[1])<<40 + int64(b[2])<<48 + int64(b[3])<<56)

}
func (packet *Packet) WriteULong(v uint64) {
	b := make([]byte, 8)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff
	b[2] = byte(v>>8) & 0xff
	b[3] = byte(v>>24) & 0xff
	b[4] = byte(v>>32) & 0xff
	b[5] = byte(v>>40) & 0xff
	b[6] = byte(v>>48) & 0xff
	b[6] = byte(v>>56) & 0xff
	packet.data = append(packet.data, b...)
	packet.offset += 8
}

func (packet *Packet) ReadULong() uint64 {
	b := packet.data[packet.offset : packet.offset+8]
	packet.offset += 8
	return uint64(uint64(b[0]) + uint64(b[1])<<8 + uint64(b[2])<<16 + uint64(b[3])<<24 + uint64(b[0])<<32 + uint64(b[1])<<40 + uint64(b[2])<<48 + uint64(b[3])<<56)
}

//float type in c#
func (packet *Packet) WriteFloat32(t float32) {
	v := math.Float32bits(t)
	b := make([]byte, 4)
	b[0] = byte(v) & 0xff
	b[1] = byte(v>>8) & 0xff
	b[2] = byte(v>>8) & 0xff
	b[3] = byte(v>>24) & 0xff
	packet.data = append(packet.data, b...)
	packet.offset += 4

}

func (packet *Packet) ReadFloat32() float32 {
	b := packet.data[packet.offset : packet.offset+8]
	packet.offset += 4
	v := uint32(uint32(b[0]) + uint32(b[1])<<8 + uint32(b[2])<<16 + uint32(b[3])<<24)
	return math.Float32frombits(v)
}

//Double type in c#
func (packet *Packet) WriteFloat64(t float64) {
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
	packet.data = append(packet.data, b...)
	packet.offset += 8

}

func (packet *Packet) ReadFloat64() float64 {
	b := packet.data[packet.offset : packet.offset+8]
	packet.offset += 8
	v := uint64(uint64(b[0]) + uint64(b[1])<<8 + uint64(b[2])<<16 + uint64(b[3])<<24 + uint64(b[0])<<32 + uint64(b[1])<<40 + uint64(b[2])<<48 + uint64(b[3])<<56)
	return math.Float64frombits(v)
}

func (packet *Packet) WriteString(s string) {
	b := []byte(s)

	packet.WriteInt32((len(b)))

	packet.data = append(packet.data, b...)
	packet.offset += len(b)
}

func (packet *Packet) ReadString() string {
	lenStr := packet.ReadInt32()

	ret := string(packet.data[packet.offset : packet.offset+lenStr])
	packet.offset = packet.offset + 4 + lenStr
	return ret
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: ", err.Error())
		os.Exit(2)
	}
}
