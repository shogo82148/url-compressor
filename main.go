package main

import (
	"fmt"
	"io"
	"net/url"

	"github.com/shogo82148/base45"
)

func main() {
	fmt.Println(encode([]byte("shogo82148.github.io/blog/2023/10/01/2023-10-01-github-actions-notify-slack/")))
	v, err := decode("0-SLNDB9IQ9IIU%2FOR1NB0QGEF7$FNPDI%2F9-%2AHY:5MGPU%2AQYSQDFM4HVC27J4B6MO..2KQLE3I")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(v))
}

func encode(data []byte) string {
	buf := NewBuffer([]byte{})

	for _, c := range data {
		e := encodeTable[c]
		buf.WriteBitsLSB(e.code, e.length)
	}

	// Encode the terminator.
	for buf.Len()%8 != 0 {
		buf.WriteBit(1)
	}

	return url.PathEscape("0" + base45.EncodeToString(buf.Bytes()))
}

func decode(data string) ([]byte, error) {
	version := data[0]
	if version == '0' {
		return decode0(data[1:])
	}
	return nil, fmt.Errorf("unknown version: %c", version)
}

func decode0(data string) ([]byte, error) {
	tmp, err := base45.DecodeString(data)
	if err != nil {
		return nil, err
	}
	buf := NewBuffer(tmp)

	ret := []byte{}
	n := decodeTree
	for {
		bit, err := buf.ReadBit()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		n = n.children[bit]
		if n.children[0] == nil && n.children[1] == nil {
			ret = append(ret, byte(n.value))
			n = decodeTree
		}
	}
	return ret, nil
}
