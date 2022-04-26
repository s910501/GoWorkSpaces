package split_string

import (
	"reflect"
	"testing"
)

// 测试组
func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := []testCase{
		testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		testCase{"abcef", "bc", []string{"a", "ef"}},
		testCase{"品亦萨拉丁有", "萨", []string{"品亦", "拉丁有"}},
	}
	for _, tc := range testGroup {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("want:%v got:%v\n", tc.want, got)
		}
	}
}

// 子测试
func Test10Split(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{
		"case_1": testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case_2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case_3": testCase{"abcef", "bc", []string{"a", "ef"}},
		"case_4": testCase{"品亦萨拉丁有", "萨", []string{"品亦", "拉丁有"}},
	}
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}

//基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}

// 性能比较测试
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}

func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}

func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 10)
}

// go test
// go test -run=Test10Split/case_1
// go test -cover
// go test -cover -coverprofile=c
// go tool cover -html=c

// go test -bench=Split
// go test -bench=Split -benchmem
// go test -bench=Fib2
// go test -bench=Split -cpu=1

// go example ===> doc
