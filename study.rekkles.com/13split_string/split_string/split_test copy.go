package split_string

import (
	"reflect"
	"testing"
)

func Test1Split(t *testing.T) {
	ret := Split("babcbef", "b")
	want := []string{"", "a", "c", "ef"}
	if !reflect.DeepEqual(ret, want) {
		t.Errorf("want:%v but got %v", want, ret)
	}
}

func Test2Split(t *testing.T) {
	ret := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(ret, want) {
		t.Errorf("want:%v but got %v", want, ret)
	}
}

func Test3Split(t *testing.T) {
	ret := Split("abcef", "bc")
	want := []string{"a", "ef"}
	if !reflect.DeepEqual(ret, want) {
		t.Fatalf("want:%v but got %v", want, ret)
	}
}
