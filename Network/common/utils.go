package common

func ByteToInt(b []byte) int {
	return int(int(b[0]) + int(b[1])<<8 + int(b[2])<<16 + int(b[3])<<24)
}
