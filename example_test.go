package iterfolder

import (
	"fmt"
	"iter"
)

func Example_IterFolder() {
	// define input types and data
	type x int
	type y int
	type z int
	input := [][]int{[]int{1, 2, 3}, []int{1, 2, 4}, []int{2, 3, 5}, []int{3, 4, 6}, []int{3, 5, 7}}

	// construct input iterator to send to folder
	type xyz = ABC[x, y, z]
	xyzIter := func() iter.Seq[xyz] {
		return func(yield func(xyz) bool) {
			for _, in := range input {
				yield(xyz{x(in[0]), y(in[1]), z(in[2])})
			}
		}
	}()

	// run the folder
	fmt.Println("")
	for a := range IterFolder[x, y, z](xyzIter) {
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
