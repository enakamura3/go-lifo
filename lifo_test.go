package lifo

import (
	"testing"
)

func TestAdd(t *testing.T) {
	l := New()
	Add(&l, "hello")
	val := len(l.v)
	want := 1
	if val != want {
		t.Fatalf("Want %v, got %v\n", want, val)
	}
}

func TestRemove(t *testing.T) {
	l := New()
	Add(&l, "hello")
	Add(&l, "world")
	val := Remove(&l)
	want := "world"
	if val != want {
		t.Fatalf("Want %v, got %v\n", want, val)
	}
}

func TestIsEmpty(t *testing.T) {
	l := New()
	val := IsEmpty(&l)
	want := true
	if val != want {
		t.Fatalf("Want %v, got %v\n", want, val)
	}
}
