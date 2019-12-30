package snowflake

import (
	"testing"
)

// TestGenerate test will it conflict
func TestGenerate(t *testing.T) {
	node := NewNode(1, 1)
	m := map[int64]bool{}

	for i := 0; i < 10000000; i++ {
		id := node.NextId()
		if _, exist := m[id]; exist != false {
			t.Fail()
		}
		m[id] = true
	}
}

// BenchmarkGenerate bench generation
func BenchmarkGenerate(b *testing.B) {
	node := NewNode(1, 1)
	for i := 0; i < b.N; i++ {
		node.NextId()
	}
}
