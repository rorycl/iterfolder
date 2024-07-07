# iterfolder

iterfolder is a package for go1.22 (using the "rangefunc" experiment)
or higher that takes an iterator which is a composite of three types
and "folds" left hand items of common value.

This might be useful, for example, for putting the results of a sql
"domain aggregate" query returning rows coerced to compound struct of
structs (e.g. rows of `{person}{car}{ticket}`) into a tiered iterator
which can be used as follows:

```go
for _, person := range <IterFolder.results> {
	// do something with person
	for _, car := person.Iter() {
		// do something with car
		for _, ticket := car.Iter() {
			// do something with ticket
		}
	}
}
```
Note that the input iterator should be pre-sorted and that duplicate
"rows" are not squashed.

Note that as of July 2024, Go templates do not yet support iterating
over an iter.Seq. See https://go.dev/issue/66107.

For more information about the rangefunc experiment, see
https://go.dev/wiki/RangefuncExperiment.

## License

MIT
