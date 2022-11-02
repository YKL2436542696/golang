package xinneng_test

import (
	"github.com/cncamp/golang/stu02/mod11-test/xinneng"
	"testing"
)

func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xinneng.MakeSliceWithoutAlloc()
	}
}

func BenchmarkMakeSliceWithPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xinneng.MakeSliceWithPreAlloc()
	}
}
