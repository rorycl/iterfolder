// iterfolder is a package that takes an iterator which is a
// composite of three types and "folds" the right hand items into
// common values in the left hand items.
//
// This might be useful, for example, for putting the results of a
// compound sql "domain aggregate" query returning rows coerced to
// {people}{cars}{tickets} into an iterator which can then be printed,
// or otherwise used, to allow nested iteration such as:
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
// object/s. ABC is the required type of Iter.Seq to provide to IterFolder
type ABC[A, B, C comparable] struct {
	A A
	B B
	C C
}

// IterFolder takes an iter.Seq of ABC and returns a three-level
// iterable of iter.Seq[Obj[A, Obj[B, C]]] which "folds" the right hand
// types into the left hand types. So,
//
//	1 2 3
//	1 2 4
//	2 3 4
//
// will (if put into an interator of the ABC type), be folded to
//
//	1
//	  2
//	    3
//	    4
//	2
//	  3
//	    4
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
			case aobj.eq(ta) && bobj.eq(tb):
				bobj.add(tc)
				aobj.replace(bobj)
			case aobj.eq(ta):
				bobj = bT{This: tb}
				bobj.add(tc)
				aobj.add(bobj)
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
