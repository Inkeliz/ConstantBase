/**
 *	Copyright (c) 2017 Inkeliz
 *  Copyright (c) 2016 - 2017 Paragon Initiative Enterprises.
 *  Copyright (c) 2014 Steve "Sc00bz" Thomas (steve at tobtu dot com)
 *
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 */
package ConstantBase64

func Encode(b []byte) []byte {
	var destination []byte
	var i int

	bLen := len(b)

	for i = 0; i+3 <= bLen; i += 3 {
		val := int(b[i+0])<<16 | int(b[i+1])<<8 | int(b[i+2])

		destination = append(destination, encode8BitsTo6(val >> 18 & 63))
		destination = append(destination, encode8BitsTo6(val >> 12 & 63))
		destination = append(destination, encode8BitsTo6(val >> 6 & 63))
		destination = append(destination, encode8BitsTo6(val >> 0 & 63))
	}

	switch bLen - i {
	case 1:
		val := int(b[i+0]) << 16
		destination = append(destination, encode8BitsTo6(val >> 18 & 63))
		destination = append(destination, encode8BitsTo6(val >> 12 & 63))

	case 2:
		val := int(b[i+0])<<16 | int(b[i+1])<<8
		destination = append(destination, encode8BitsTo6(val >> 18 & 63))
		destination = append(destination, encode8BitsTo6(val >> 12 & 63))
		destination = append(destination, encode8BitsTo6(val >> 6 & 63))
	}

	return destination
}

func EncodeWithPad(b []byte) []byte {
	encoded := Encode(b)

	switch len(encoded) & 3 {
	case 2:
		encoded = append(encoded, 0x3D)
		fallthrough
	case 3:
		encoded = append(encoded, 0x3D)
	}

	return encoded
}

func EncodeToString(b []byte) string {
	return string(Encode(b))
}

func EncodeWithPadToString(b []byte) string {
	return string(EncodeWithPad(b))
}

func encode8BitsTo6(i int) byte {
	var diff = 65

	diff += ((25 - i) >> 8) & 6
	diff -= ((51 - i) >> 8) & 75
	diff -= ((61 - i) >> 8) & 15
	diff += ((62 - i) >> 8) & 3

	return byte(diff + i)
}
