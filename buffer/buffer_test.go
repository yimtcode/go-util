package buffer

import (
	"bytes"
	"io"
	"testing"
)

func TestNewBuffer(t *testing.T) {
	data := []byte{1, 2, 3}
	buf := NewBuffer(data)

	if !bytes.Equal(buf.Bytes(), data) {
		t.Fatal("Buffer NewBuffer exception")
	}
}

func TestBuffer_TryRead(t *testing.T) {
	// simulation tcp read package
	// package format:
	// 		package length | package data
	//  	1 byte		   | any byte
	data := []byte{
		1, 1,
		2, 2, 2,
		3, 3, 3, 3,
	}
	buffer := NewBuffer(data)


	readCount := func(data []byte)  int {
		length := len(data)
		if length == 0 {
			return 0
		}

		packageLength := int(data[0]) +1
		if packageLength > length {
			return 0
		}

		return packageLength
	}

	buf := make([]byte, 1024)
	for true {
		n, err := buffer.TryRead(buf, readCount)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				t.Fatal(err.Error())
			}
		}

		if n == 0 {
			break
		}

		t.Log(buf[:n])
	}
}