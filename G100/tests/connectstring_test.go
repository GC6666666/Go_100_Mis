package tests

import (
	"strings"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	values := []string{
		"my", "name", "is", "go",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concat(values)
	}
}
func concat(values []string) string {
	total := 0
	for i := range values {
		total += len(values[i])
	}
	s := strings.Builder{}
	s.Grow(total)
	for _, value := range values {
		_, _ = s.WriteString(value)
	}
	return s.String()
}
