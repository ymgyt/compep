package compep

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsPrime(t *testing.T) {
	primes := []int{
		2, 3, 5, 7, 11, 13,
		3557, 3559, 3571,
	}

	composites := []int{
		100, 1000, 10002,
	}

	if IsPrime(1) {
		t.Errorf("1 is not prime number")
	}

	for _, p := range primes {
		if !IsPrime(p) {
			t.Errorf("%d IsPrime => false, want true", p)
		}
	}

	for _, c := range composites {
		if IsPrime(c) {
			t.Errorf("%d IsPrime => true, want false", c)
		}
	}
}

func TestGenPrimes(t *testing.T) {
	want := []int{2, 3, 5, 7, 11, 13, 17}
	got := GenPrimes(len(want))

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("(-got +want)\n%s", diff)
	}

	a := GenPrimes(100)
	fmt.Println(a)
}
