package tests

import (
	"errors"
	"log"
	"strings"
	"testing"
)

func TestNilReceiver(t *testing.T) {
	customer := Customers{
		age:  30,
		name: "John",
	}
	if err := customer.Validate(); err != nil {
		log.Fatalf("customer is invalid: %v", err)
	}
}

type MultiError struct {
	errs []string
}

func (m *MultiError) add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

type Customers struct {
	age  int
	name string
}

func (c Customers) Validate() error {
	var m *MultiError
	if c.age < 0 {
		m = &MultiError{}
		m.add(errors.New("age is negative"))
	}
	if c.name == "" {
		if m != nil {
			m = &MultiError{}
		}
		m.add(errors.New("name is nil"))
	}
	if m != nil {
		return m
	}
	return nil
}
