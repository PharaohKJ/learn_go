// newmathパッケージはちょっとしたサンプルパッケージです。
package newmath

import (
	"testing"
)

// Sqrtはxの平方根の近似値を返します。
func TestSqrt(t *testing.T) {
	r := Sqrt(3)
	if r > 1.7 {
		t.Errorf("error! %v", r)
	}
}
