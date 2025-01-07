package tests

import (
	"fmt"
	"io"
	"testing"
)

func TestReceiver(t *testing.T) {
	c := ccustomer{&data{balance: 2.00}}
	c.add(3.00)
	fmt.Println(c.data.balance)
}

type ccustomer struct {
	data *data
}

type data struct {
	balance float64
}

func (c ccustomer) add(v float64) {
	c.data.balance += v
}

func readFull(r io.Reader, buf []byte) (n int, err error) {
	if len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}
