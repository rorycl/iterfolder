package iterfolder

import (
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
	"testing"
	"text/template"

	"github.com/google/go-cmp/cmp"
)

func TestAlternative(t *testing.T) {

}

func TestSimpleIter(t *testing.T) {
	type x int
	type y int
	type z int
	// type alias
	type xyz = ABC[x, y, z]

	xyzIter := func() iter.Seq[xyz] {
		input := []xyz{
			xyz{1, 2, 3},
			xyz{1, 2, 4},
			xyz{2, 3, 5},
			xyz{3, 4, 6},
			xyz{3, 5, 7},
		}
		return func(yield func(xyz) bool) {
			for _, in := range input {
				yield(in)
			}
		}
	}()

	expected := `1
> 2
> > 3
> > 4
2
> 3
> > 5
3
> 4
> > 6
> 5
> > 7
`

	// run the folder
	output := ""
	for aa := range IterFolder[x, y, z](xyzIter) {
		output += fmt.Sprintln(aa.This)
		for bb := range aa.Iter() {
			output += fmt.Sprintln(">", bb.This)
			for cc := range bb.Iter() {
				output += fmt.Sprintln("> >", cc)
			}
		}
	}
	if diff := cmp.Diff(expected, output); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
	t.Log(output)
}

func TestGenIter(t *testing.T) {

	// define types
	type a struct{ id, f1, f2 string }
	type b struct{ id, f2, f3 float64 }
	type c struct{ id, f3, f4 string }

	// instantiate a/b/c structs
	rowsFromStrings := func(ss []string) iter.Seq[ABC[a, b, c]] {
		return func(yield func(ABC[a, b, c]) bool) {
			for _, line := range ss {
				s := strings.Fields(line)
				_a := a{s[0], s[1], s[2]}
				_b := func() b {
					f := func(si string) float64 {
						f32, _ := strconv.ParseFloat(si, 10)
						return f32
					}
					return b{f(s[3]), f(s[4]), f(s[5])}
				}()
				_c := c{s[6], s[7], s[8]}
				yield(ABC[a, b, c]{_a, _b, _c})
			}
		}
	}
	var input = `a1 a1a a1b 1.1 1.2 1.3 n1 n1a n1b
			 a1 a1a a1b 1.1 1.2 1.3 n2 n2a n2b
			 a1 a1a a1b 1.1 1.2 1.3 n3 n3a n3b
			 a1 a1a a1b 2.1 2.2 2.3 n4 n4a n4b
			 a2 a2a a2b 3.1 3.2 3.3 n5 n5a n6b
			 a2 a2a a2b 3.1 3.2 3.3 n6 n5a n6b`

	var expected = strings.ReplaceAll(`{a1 a1a a1b}
			 > {1.1 1.2 1.3}
			 > > {n1 n1a n1b}
			 > > {n2 n2a n2b}
			 > > {n3 n3a n3b}
			 > {2.1 2.2 2.3}
			 > > {n4 n4a n4b}
			 {a2 a2a a2b}
			 > {3.1 3.2 3.3}
			 > > {n5 n5a n6b}
			 > > {n6 n5a n6b}
`, "			 ", "")

	// construct input
	allLines := strings.Split(input, "\n")
	lines := make([]string, len(allLines))
	for i, a := range allLines {
		lines[i] = strings.TrimSpace(a)
	}
	iterABC := rowsFromStrings(lines)

	output := ""
	// run the folder
	for aa := range IterFolder[a, b, c](iterABC) {
		output += fmt.Sprintln(aa.This)
		for bb := range aa.Iter() {
			output += fmt.Sprintln(">", bb.This)
			for cc := range bb.Iter() {
				output += fmt.Sprintln("> >", cc)
			}
		}
	}

	if diff := cmp.Diff(expected, output); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
	t.Log(output)
}

// as at 07 July 2024, text/template does not support iterating over an
// iter.Seq. See https://go.dev/issue/66107
func TestTemplate(t *testing.T) {

	//
	t.SkipNow()

	type n string
	type o string
	type p string
	// type alias
	type nop = ABC[n, o, p]

	nopIter := func() iter.Seq[nop] {
		input := []nop{
			nop{"one", "two", "three"},
			nop{"one", "two", "four"},
			nop{"two", "three", "five"},
			nop{"three", "four", "six"},
			nop{"three", "five", "seven"},
		}
		return func(yield func(nop) bool) {
			for _, in := range input {
				yield(in)
			}
		}
	}()

	resulter := IterFolder[n, o, p](nopIter)

	// this resulter does not work either
	// resulter := func(yield func(int) bool) {
	// 	for _, v := range []int{1, 2, 3} {
	// 		if !yield(v) {
	// 			return
	// 		}
	// 	}
	// }
	//
	// this iterable works
	// resulter := []string{"a", "b"}

	tpl := `
{{ range $aa := . }}
{{ $aa.This }}
	{{ range $bb := $aa.Iter }}
	{{ $bb.This }}
		{{ range $cc := $bb.Iter }}
		{{ $cc }}
{{ end }}
{{ end }}
{{ end }}
`
	tplParsed := template.Must(template.New("test").Parse(tpl))
	err := tplParsed.Execute(os.Stdout, resulter)
	if err != nil {
		t.Fatal(err)
	}
}
