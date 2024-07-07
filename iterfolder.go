// iterfolder is a package for go1.22 (using the "rangefunc" experiment)
// or higher that takes an iterator which is a composite of three types
// and "folds" left hand items of common value.
//
// This might be useful, for example, for putting the results of a sql
// "domain aggregate" query returning rows coerced to compound struct of
// structs (e.g. rows of `{perspon}{car}{ticket}`) into a tiered
// iterator which can be used as follows:
//
//	for _, person := range <IterFolder.results> {
//		// do something with person
//		for _, car := person.Iter() {
//			// do something with car
//			for _, ticket := car.Iter() {
//				// do something with ticket
//			}
//		}
//	}
//
// The input iterator should be sorted  and duplicate "rows" are not
// eliminated.
//
// Note that as of July 2024, Go templates do not yet support iterating
// over an iter.Seq. See https://go.dev/issue/66107.
//
// For more information about the rangefunc experiment, see
// https://go.dev/wiki/RangefuncExperiment.
package iterfolder

import (
	"fmt"
	"iter"
)

const debug = false

// Obj is a composable type supporting recursive iteration
type Obj[T comparable, U any] struct {
	This  T
	those []U
}

func (o *Obj[T, U]) add(u U) {
	o.those = append(o.those, u)
}

func (o *Obj[T, U]) replace(u U) {
	o.those[len(o.those)-1] = u
}

func (o *Obj[T, U]) eq(n T) bool {
	return o.This == n
}

// Iter provides an iterator over the subsiduary items contained in
// [Obj].
func (o *Obj[T, U]) Iter() iter.Seq[U] {
	if debug {
		fmt.Printf("calling Iter() on %T\n", o.those)
	}
	return func(yield func(U) bool) {
		for _, v := range o.those {
			if !yield(v) {
				return
			}
		}
	}
}

// ABC is a composite struct of the three types making up the recursive
// object/s. ABC is the required type of Iter.Seq to provide to
// [IterFolder].
type ABC[A, B, C comparable] struct {
	A A
	B B
	C C
}

// IterFolder takes an iter.Seq of [ABC] and returns a three-level
// iterable of iter.Seq[Obj[A, Obj[B, C]]] which "folds" the left hand
// items of common value. Conceptually IterFolder translates an iterable
// that would provide:
//
//	1 2 3
//	1 2 4
//	2 3 4
//
// into an iterable that provides:
//
//	1
//	  2
//	    3
//	    4
//	2
//	  3
//	    4
//
// Note that duplicate rows will provide duplicate right hand values in
// the output.
func IterFolder[A, B, C comparable](abc iter.Seq[ABC[A, B, C]]) iter.Seq[Obj[A, Obj[B, C]]] {

	type aT = Obj[A, Obj[B, C]]
	type bT = Obj[B, C]

	var aobj = aT{}
	var bobj = bT{}

	return func(yield func(aT) bool) {
		started := false
		for thisABC := range abc {
			ta, tb, tc := thisABC.A, thisABC.B, thisABC.C
			if !started {
				started = true
				bobj.This = tb
				bobj.add(tc)
				aobj.This = ta
				aobj.add(bobj)
				continue
			}
			switch {
			// a and b are unchanged
			case aobj.eq(ta) && bobj.eq(tb):
				bobj.add(tc)
				aobj.replace(bobj) // replace the last a.those
			// a is unchanged
			case aobj.eq(ta):
				bobj = bT{This: tb}
				bobj.add(tc)
				aobj.add(bobj)
			// all have changed
			default:
				if !yield(aobj) {
					return
				}
				bobj = bT{This: tb}
				bobj.add(tc)
				aobj = aT{This: ta} // reinit
				aobj.add(bobj)
			}
		}
		yield(aobj)
	}
}
