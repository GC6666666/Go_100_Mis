package tests

import (
	"fmt"
	"testing"
)

type account struct {
	balance float32
}

func TestRange(t *testing.T) {
	accounts := []*account{
		{
			balance: 100,
		},
		{
			balance: 200,
		},
		{
			balance: 300,
		},
	}
	//for _, b := range accounts {
	//	b.balance += 1000
	//}
	for _, a := range accounts {
		a.balance += 1000
	}
	for _, a := range accounts {
		fmt.Printf("%.2f\n", a.balance)
	}
}
