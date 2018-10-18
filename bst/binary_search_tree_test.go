package bst

import "testing"

func TestStringBinarySearchTree(t *testing.T) {
	type prepare func(*StringBinarySearchTree)
	type check func(*testing.T, *StringBinarySearchTree)

	tests := []struct {
		desc string
		prepare
		check
	}{
		{
			desc: "inseart one",
			prepare: func(bst *StringBinarySearchTree) {
				bst.Insert(100, "value")
			},
			check: func(t *testing.T, bst *StringBinarySearchTree) {
				got, found := bst.Search(100)
				if !found {
					t.Fatal("failed to search")
				}
				want := "value"
				if got != want {
					t.Errorf("got %v, want %v", got, want)
				}
			},
		},
		{
			desc: "inseart a lot",
			prepare: func(bst *StringBinarySearchTree) {
				bst.Insert(50, "v50")
				bst.Insert(10, "v10")
				bst.Insert(100, "v100")
				bst.Insert(20, "v20")
				bst.Insert(90, "v90")
				bst.Insert(30, "v30")
				bst.Insert(80, "v80")
				bst.Insert(60, "v60")
				bst.Insert(70, "v70")
				bst.Insert(40, "v40")
			},
			check: func(t *testing.T, bst *StringBinarySearchTree) {
				got, found := bst.Search(100)
				if !found {
					t.Fatal("failed to search")
				}
				want := "v100"
				if got != want {
					t.Errorf("got %v, want %v", got, want)
				}
				t.Logf("\n%s\n", bst.String())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			bst := NewStringBinarySearchTree()
			tt.prepare(bst)
			tt.check(t, bst)
		})
	}
}
