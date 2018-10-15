package compep

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLines(t *testing.T) {
	type checkFn func(*testing.T, Lines)

	check := func(t *testing.T, got, want interface{}) {
		t.Helper()
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("(-got +want)\n%s", diff)
		}
	}

	tests := []struct {
		desc string
		in   string
		checkFn
	}{
		{
			desc: "just one string line",
			in:   "golang or gohome\n",
			checkFn: func(t *testing.T, lines Lines) {
				got, want := lines.Next().String(), "golang or gohome"
				check(t, got, want)
			},
		},
		{
			desc: "just one int line",
			in:   "12345\n",
			checkFn: func(t *testing.T, lines Lines) {
				got, want := lines.Next().Int(), 12345
				check(t, got, want)
			},
		},
		{
			desc: "numbers",
			in:   "10 20 30 40\n",
			checkFn: func(t *testing.T, lines Lines) {
				got, want := lines.Next().Ints(), []int{10, 20, 30, 40}
				check(t, got, want)
			},
		},
		{
			desc: "strings",
			in:   "aaa bbb ccc ddd\n",
			checkFn: func(t *testing.T, lines Lines) {
				got, want := lines.Next().Strings(), []string{"aaa", "bbb", "ccc", "ddd"}
				check(t, got, want)
			},
		},
		{
			desc: "multi line numbers",
			in:   "10 20 30 40\n50 60 70 80\n90 100 110 120\n",
			checkFn: func(t *testing.T, lines Lines) {
				got, want := lines.Next().Ints(), []int{10, 20, 30, 40}
				check(t, got, want)
				got, want = lines.Next().Ints(), []int{50, 60, 70, 80}
				check(t, got, want)
				got, want = lines.Next().Ints(), []int{90, 100, 110, 120}
				check(t, got, want)
			},
		},
		{
			desc: "multi line strings",
			in:   "aaa AAA\nbbb BBB\nccc CCC\n",
			checkFn: func(t *testing.T, lines Lines) {
				got, want := lines.Next().Strings(), []string{"aaa", "AAA"}
				check(t, got, want)
				got, want = lines.Next().Strings(), []string{"bbb", "BBB"}
				check(t, got, want)
				got, want = lines.Next().Strings(), []string{"ccc", "CCC"}
				check(t, got, want)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			b := bytes.NewBufferString(tt.in)
			lines := NewLines(b, 1024)
			tt.checkFn(t, lines)
		})
	}
}
