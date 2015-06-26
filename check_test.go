package inn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheck(t *testing.T) {

	if assert.True(t, Check("")) {

		assert.True(t, Check("500100732259"))
		assert.True(t, Check("7830002293"))
		assert.True(t, Check("7736207543"))
		assert.False(t, Check("500100732258"))
		assert.False(t, Check("7830002213"))
		assert.False(t, Check("3245234"))
	}
}

func BenchmarkCheck(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {

		Check("500100732259")
	}
}
