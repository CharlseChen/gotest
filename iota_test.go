package carnival

import "testing"

const (
	A = iota
	B
	C = 10
	D = iota
	E
)

func TestIota(t *testing.T) {
	t.Log(A, B, C, D, E)
	//0 1 10 3 4
}
