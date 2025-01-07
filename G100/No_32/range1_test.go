package No_32

import (
	"fmt"
	"testing"
)

func TestRange1(t *testing.T) {
	b := &balance{
		id:   12,
		data: 1000,
	}
	s := &store{
		message: make(map[int]*balance),
	}
	s.put(12, b)
	fmt.Printf("store[message{id:%d,data:%d}]\n", 12, s.message[12].data)
	fmt.Printf("balance{id:%d,data:%d}\n", b.id, b.data)
	s.message[12].data = 2000
	fmt.Printf("balance{id:%d,data:%d}\n", b.id, b.data)
	b.data = 3000
	fmt.Printf("store[message{id:%d,data:%d}]\n", 12, s.message[12].data)
}

type balance struct {
	id   int
	data int
}

type store struct {
	message map[int]*balance
}

func (u store) put(id int, b *balance) {
	u.message[id] = b
}
