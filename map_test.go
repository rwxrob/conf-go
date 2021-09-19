package conf

import "testing"

func TestNewMap(t *testing.T) {
	f := func(m Map) { return }
	m := NewMap()
	f(m)
}
