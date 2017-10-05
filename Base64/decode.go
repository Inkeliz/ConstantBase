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

import (
	"errors"
)

func Decode(b []byte) ([]byte, error) {
	var destination []byte
	var i, err int
	var bLen = len(b)

	bLen -= (((60 - int(b[bLen-1])) & (int(b[bLen-1]) - 62)) >> 8) & 1
	bLen -= (((60 - int(b[bLen-1])) & (int(b[bLen-1]) - 62)) >> 8) & 1

	for i = 0; i+4 <= bLen; i += 4 {
		c0, c1, c2, c3 := decode6BitsTo8(int(b[i+0])), decode6BitsTo8(int(b[i+1])), decode6BitsTo8(int(b[i+2])),	decode6BitsTo8(int(b[i+3]))
		val := c0<<18 | c1<<12 | c2 <<6 | c3
		err |= (c0 | c1 | c2 | c3) >> 8

		destination = append(destination, byte(val >> 16 & 255))
		destination = append(destination, byte(val >> 8 & 255))
		destination = append(destination, byte(val >> 0 & 255))
	}

	switch bLen - i {
	case 2:
		c0, c1 := decode6BitsTo8(int(b[i+0])), decode6BitsTo8(int(b[i+1]))
		val := c0<<18 | c1<<12
		err |= (c0 | c1) >> 8

		destination= append(destination, byte(val >> 16 & 255))
	case 3:
		c0, c1, c2 := decode6BitsTo8(int(b[i+0])), decode6BitsTo8(int(b[i+1])), decode6BitsTo8(int(b[i+2]))
		val := c0<<18 | c1<<12 | c2 << 6
		err |= (c0 | c1 | c2) >> 8

		destination = append(destination, byte(val >> 16 & 255))
		destination = append(destination, byte(val >> 8 & 255))
	}

	if err != 0 {
		return nil, errors.New("invalid characters")
	}

	return destination, nil
}

func DecodeToString(b []byte) (string, error) {
	destination, err := Decode(b)
	return string(destination), err
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
