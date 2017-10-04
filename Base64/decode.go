package ConstantBase64


func Decode(b []byte) string {
	var destination string
	var i int

	strLen := len(b)
	strByte := make([]byte, 0, strLen+(strLen%4))
	strByte = append(strByte, b...)

	if strByte[strLen-2] == 61 {
		strLen--
	}

	if strByte[strLen-1] == 61 {
		strLen--
	}

	for i = 0; i < strLen; i += 4 {
		chuck := strByte[i:i+4]
		c0 := decode6BitsTo8(int(chuck[0]))
		c1 := decode6BitsTo8(int(chuck[1]))
		c2 := decode6BitsTo8(int(chuck[2]))
		c3 := decode6BitsTo8(int(chuck[3]))

		destination += string(((c0 << 2) | (c1 >> 4)) & 255)
		destination += string(((c1 << 4) | (c2 >> 2)) & 255)
		destination += string(((c2 << 6) | c3) & 255)
	}

	return destination[:(strLen * 6) / 8]
}

func decode6BitsTo8(i int) int {
	var diff = -1

	diff += (((64 - i) & (i - 91)) >> 8) & (i - 64)
	diff += (((96 - i) & (i - 123)) >> 8) & (i - 70)
	diff += (((47 - i) & (i - 58)) >> 8) & (i + 5)
	diff += (((42 - i) & (i - 44)) >> 8) & 63
	diff += (((46 - i) & (i - 48)) >> 8) & 64

	return diff
}
