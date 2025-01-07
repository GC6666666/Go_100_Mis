package tests

import (
	"fmt"
	"testing"
)

func TestForRangePointer(t *testing.T) {
	customers := []*Customer{
		{
			ID:      "1",
			Balance: 10,
		},
		{
			ID:      "2",
			Balance: 20,
		},
		{
			ID:      "3",
			Balance: 30,
		},
	}

	s := Store{
		m: make(map[string]*Customer),
	}
	s.storeCustomers(customers)
	fmt.Println(s)
}

type Customer struct {
	ID      string
	Balance float64
}

func (c Customer) String() string {
	return fmt.Sprintf("{ID is %s, Balance is %.2f}\n", c.ID, c.Balance)
}

type Store struct {
	m map[string]*Customer
}

func (s *Store) storeCustomers(customers []*Customer) {
	for _, customer := range customers {
		s.m[customer.ID] = customer
	}
}

func (s Store) String() string {
	str := fmt.Sprintf("Store is {\n")
	for key, v := range s.m {
		str += fmt.Sprintf("key is %s, value is %s", key, v)
	}
	str += "}"
	return str
}
