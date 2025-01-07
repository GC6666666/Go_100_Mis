package config

type IntConfig struct {
	Value int
}

func (c *IntConfig) Get() int {
	return c.Value
}

func (c *IntConfig) Set(val int) {
	c.Value = val
}

type IntConfigGetter interface {
	Get() int
}
