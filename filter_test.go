package filter

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := uint(1000)
	b := New(s)
	l := uint(len(b.bits))
	if l != s {
		t.Errorf("Number of bits was incorrect. Got: %d, want: %d", l, s)
	}
	if b.n != 0 {
		t.Errorf("Number of items should be 0, got: %d", b.n)
	}
	if len(b.hashFuncs) < 1 {
		t.Errorf("Invalid number of hash functions. Should contain at least one, got: %d", len(b.hashFuncs))
	}
}

func TestAdd(t *testing.T) {
	b := New(1000)
	b.Add([]byte("potato"))
	if b.n != 1 {
		t.Errorf("Number of items was incorrect. Got: %d, want %d", b.n, 1)
	}

	b.Add([]byte("another potato"))
	if b.n != 2 {
		t.Errorf("Number of items was incorrect. Got: %d, want %d", b.n, 2)
	}

	for i := 0; i < 200; i++ {
		b.Add([]byte(string(i)))
	}
	if b.n != 202 {
		t.Errorf("Number of items was incorrect. Got: %d, want %d", b.n, 202)
	}
}

func TestLookup(t *testing.T) {
	n := uint(10000000)
	c := n / 256
	b := New(n)
	b.Add([]byte("potato"))
	if !b.Lookup([]byte("potato")) {
		t.Errorf("False negative when searching for %s.", "potato")
	}
	if b.Lookup([]byte("llama")) {
		t.Errorf("False positive when searching for %s.", "llama")
	}

	for i := uint(0); i < c; i++ {
		b.Add([]byte(string(i)))
	}

	fp := uint(0)
	for i := uint(0); i < c; i++ {
		if !b.Lookup([]byte(string(i))) {
			t.Errorf("False negative when searching for %d.", i)
		}
		if b.Lookup([]byte(string(n - i))) {
			fp++
			fmt.Println("False Positive:", n-i)
		}
	}
	fmt.Printf("Size of Filter: %d\n"+
		"Number of items: %d\n"+
		"Number of False Positives: %d\n"+
		"Percent False Positives: %f%%\n", n, b.n, fp, float64(fp)*100/float64(n/2))
}
