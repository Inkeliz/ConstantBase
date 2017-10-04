package ConstantBase64

func Encode(b []byte) string {
	return encode(b)
}

func EncodeWithPad(b []byte) string {
	encoded := encode(b)

	switch len(encoded) % 4 {
	case 2:
		encoded += "=="
	case 3:
		encoded += "="
	}

	return encoded
}

func encode(b []byte) string {
	var destination string
	var i int

	strLen := len(b)
	strByte := make([]byte, 0, (strLen/3)*3)
	strByte = append(strByte, b...)

	for i = 0; i < strLen; i += 3 {
		chuck := strByte[i:i+3]
		b0 := int(chuck[0])
		b1 := int(chuck[1])
		b2 := int(chuck[2])

		destination += encode8BitsTo6(b0 >> 2)
		destination += encode8BitsTo6(((b0 << 4) | (b1 >> 4)) & 63)
		destination += encode8BitsTo6(((b1 << 2) | (b2 >> 6)) & 63)
		destination += encode8BitsTo6(b2 & 63)
	}

	return destination[:len(destination)-(i-strLen)]
}

func encode8BitsTo6(i int) string {
	var diff = 65

	diff += ((25 - i) >> 8) & 6
	diff -= ((51 - i) >> 8) & 75
	diff -= ((61 - i) >> 8) & 15
	diff += ((62 - i) >> 8) & 3

	return string(diff + i)
}
