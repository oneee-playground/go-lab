package bytes_test

import (
	"testing"
	"unsafe"
)

func doWork(s string) []byte {
	// 컴파일러가 최적화시키지 않도록 변환된 string을 사용
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func doTypeConversion(b []byte) []byte {
	s := string(b)
	return doWork(s)
}

func doUnsafeConversion(b []byte) []byte {
	s := unsafe.String(unsafe.SliceData(b), len(b))
	return doWork(s)
}

func BenchmarkTypeConversion(b *testing.B) {
	bytes := []byte{'h', 'e', 'l', 'l', 'o'}
	for i := 0; i < b.N; i++ {
		bytes = doTypeConversion(bytes)
	}
}

func BenchmarkUnsafeConversion(b *testing.B) {
	bytes := []byte{'h', 'e', 'l', 'l', 'o'}
	for i := 0; i < b.N; i++ {
		bytes = doUnsafeConversion(bytes)
	}
}
