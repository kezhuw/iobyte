// Package iobyte provides reader that returns unlimited bytes to its caller.
// If you want limited reader, wrap it with io.LimitedReader.
package iobyte

import "io"

type Reader interface {
	io.Reader
	io.ByteReader
}

type reader byte

func (c reader) ReadByte() (byte, error) {
	return byte(c), nil
}

func (c reader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = byte(c)
	}
	return len(p), nil
}

// ByteReader creates an reader that returns unlimited bytes with value c.
func ByteReader(c byte) Reader {
	return reader(c)
}

// ZeroReader is an reader that returns unlimited zero bytes.
//
// os.Open("/dev/zero") returns a similar reader.
var ZeroReader = ByteReader(0)
