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

## database usage example

See [examples](./examples) for a database example using pgx, embedded
structs and `pgx.RowToStructByName` via a docker example.

Example output:

```
Roald Dahl cars 3 tickets 4
> {ACA-00138-A Kia Stinger}
> > c2890f7b 2023-11-03 lat/long 51.829620/-2.532015
> {BDD-00845-B Honda Insight}
> > 6b30c01e 2024-02-21 lat/long 51.726021/-1.755015
> > fafb4c88 2023-06-06 lat/long 51.824020/-2.490015
> {EEE-00744-E Volvo V90}
> > af36e4bd 2024-02-17 lat/long 51.873222/-2.859015
Gosho Aoyama cars 3 tickets 4
> {ADB-00245-A Volvo V60}
> > 61509c9b 2023-05-08 lat/long 51.847622/-2.667015
> {BBA-00043-B Chevrolet Colorado Extended Cab}
> > dc978c72 2023-10-24 lat/long 51.806824/-2.361015
> > 37c4e259 2023-05-28 lat/long 51.796421/-2.283015
> {FBA-00548-F Chevrolet Silverado 2500 HD Double Cab}
> > 1087a451 2024-06-28 lat/long 51.634422/-1.068015
```

## License

MIT
