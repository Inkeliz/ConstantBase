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

func Encode(b []byte) string {
	var destination string
	var i int

	bLen := len(b)

	for i = 0; i+3 <= bLen; i += 3 {
		val := int(b[i+0])<<16 | int(b[i+1])<<8 | int(b[i+2])

		destination += encode8BitsTo6(val >> 18 & 63)
		destination += encode8BitsTo6(val >> 12 & 63)
		destination += encode8BitsTo6(val >> 6 & 63)
		destination += encode8BitsTo6(val >> 0 & 63)
	}

	switch bLen - i {
	case 1:
		val := int(b[i+0]) << 16
		destination += encode8BitsTo6(val >> 18 & 63)
		destination += encode8BitsTo6(val >> 12 & 63)

	case 2:
		val := int(b[i+0])<<16 | int(b[i+1])<<8
		destination += encode8BitsTo6(val >> 18 & 63)
		destination += encode8BitsTo6(val >> 12 & 63)
		destination += encode8BitsTo6(val >> 6 & 63)
	}

	return destination
}

func EncodeWithPad(b []byte) string {
	encoded := Encode(b)

	switch len(encoded) % 4 {
	case 2:
		encoded += "=="
	case 3:
		encoded += "="
	}

	return encoded
}

func encode8BitsTo6(i int) string {
	var diff = 65

	diff += ((25 - i) >> 8) & 6
	diff -= ((51 - i) >> 8) & 75
	diff -= ((61 - i) >> 8) & 15
	diff += ((62 - i) >> 8) & 3

	return string(diff + i)
}
