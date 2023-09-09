package goroutine_test

import (
	"testing"

	"github.com/onee-only/go-lab/test/benchmark/concurrency/goroutine"
)

func BenchmarkWithoutGoroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goroutine.ReallyLongTask()
	}
}

func BenchmarkWithoGoroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go goroutine.ReallyLongTask()
	}
}
