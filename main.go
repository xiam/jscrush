/*
  Copyright (c) 2013 José Carlos Nieto, http://xiam.menteslibres.org/

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
)

const LowWord = 2
const HiWord = 50

type Script struct {
	Buf []byte
}

// Creates a dictionary of the given word length and counts ocurrences.
func dict(buf []byte, wlen int) map[string]int {

	dict := map[string]int{}

	l := len(buf)

	for i := 0; i+wlen < l; i++ {
		s := string(buf[i : i+wlen])
		if _, ok := dict[s]; ok == false {
			dict[s] = 0
		}
	}

	s := string(buf)
	for k := range dict {
		dict[k] = strings.Count(s, k)
	}

	return dict
}

// Returns the first unused byte, we need it to store a chunk.
func getUnusedByte(buf []byte) (byte, error) {
	var c byte
	for i := byte(126); i > byte(0); i-- {
		switch i {
		case
			// Non-printable bytes.
			byte(10),
			byte(13),
			byte(34),
			byte(39),
			byte(92):
		default:
			c = i
			if bytes.Contains(buf, []byte{c}) == false {
				return c, nil
			}
		}
	}
	return byte(0), errors.New("No more unused bytes\n")
}

// Takes a buffer and its keys and reconstructs the original string.
func Uncompress(buf []byte, keys []byte) ([]byte, error) {
	for i := 0; i < len(keys); i++ {

		slices := bytes.Split(buf, []byte{keys[i]})
		l := len(slices)

		buf = bytes.Join(slices[0:l-1], slices[l-1])
	}

	return buf, nil
}

// Compresses a buffer into a chunk and chunk keys.
func Compress(buf []byte) ([]byte, []byte, error) {
	keys := []byte{}

	mv := -1

	for true {

		t, err := getUnusedByte(buf)

		if err != nil {
			break
		}

		mk := ""

		for i := LowWord; i < HiWord; i++ {

			d := dict(buf, i)

			for k, v := range d {
				if v > 1 {
					// Calculating lenght.
					tlen := len(buf) - v*len(k) + v

					slen := tlen + len(k) + len(keys) + 2

					if slen < mv || mv < 0 {
						if tlen < len(buf) {
							mk = k
							mv = slen
						}
					}
				}
			}

		}

		if mk != "" {

			buf = bytes.Replace(buf, []byte(mk), []byte{t}, -1)
			buf = append(buf, t)
			buf = append(buf, []byte(mk)...)

			keys = append([]byte{t}, keys...)

		} else {
			break
		}
	}

	return buf, keys, nil
}

func Pack(buf []byte, keys []byte) string {
	return fmt.Sprintf("for(s='%s',i=0;j='%s'[i++];)with(s.split(j))s=join(pop());eval(s)", string(buf), string(keys))
}

func main() {
	buf := bytes.NewBuffer(nil)
	r := bufio.NewReader(os.Stdin)
	for true {
		b, err := r.ReadByte()
		if err != nil {
			break
		}
		buf.WriteByte(b)
	}
	t, k, err := Compress(buf.Bytes())
	if err == nil {
		fmt.Println(Pack(t, k))
	}
}
