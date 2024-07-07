package iterfolder

import (
	"fmt"
	"iter"
)

func ExampleIterFolder() {

	// define input types and data
	type xyz = ABC[int, int, int]
	input := []xyz{
		xyz{1, 2, 3},
		xyz{1, 2, 4},
		xyz{2, 3, 5},
		xyz{3, 4, 6},
		xyz{3, 5, 7},
	}

	// construct input iterator to send to folder
	xyzIter := func() iter.Seq[xyz] {
		return func(yield func(xyz) bool) {
			for _, this := range input {
				yield(this)
			}
		}
	}()

	// run the folder
	fmt.Println("")
	for a := range IterFolder[int, int, int](xyzIter) {
		fmt.Println(a.This)
		for b := range a.Iter() {
			fmt.Println(">", b.This)
			for c := range b.Iter() {
				fmt.Println("> >", c)
			}
		}
	}

	// Output:
	// 1
	// > 2
	// > > 3
	// > > 4
	// 2
	// > 3
	// > > 5
	// 3
	// > 4
	// > > 6
	// > 5
	// > > 7
}
