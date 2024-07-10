README

This example runs a "compound" query using nested structs on a
Postgresql database. The use of a struct with embedded structs for each
row utilizes the [pgx.RowToStructByName](https://pkg.go.dev/github.com/jackc/pgx/v5#RowToStructByName)
function.

To run the example setup up a postgresql version 16 database and then load the
`run.sql` file and then run the programme.

## steps

```
docker pull postgres:16

docker run -itd -e POSTGRES_USER=iterf \
           -e POSTGRES_PASSWORD=iterf \
           -p 5432:5432 --name iterf postgres

PGPASSWORD=iterf psql -h localhost -p 5432 -U iterf -f run.sql

GOEXPERIMENT=rangefunc go run dbCompoundExample.go
```

You should get output along the following lines (although `run.sql` will
create different data on each invocation):

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
Karl May cars 3 tickets 4
> {ABB-00145-A Lexus GX}
> > 4287a3b4 2024-01-17 lat/long 51.877621/-2.892015
> {DEB-00549-D Chevrolet Silverado 2500 HD Regular Cab}
> > a2cf986f 2024-01-21 lat/long 51.901222/-3.069015
> {FBB-00347-F Honda Insight}
> > e5667c1c 2024-04-07 lat/long 51.642822/-1.131015
> > a253703e 2023-08-30 lat/long 51.645222/-1.149015
Edgar Rice Burroughs cars 2 tickets 4
> {BBE-00093-B Honda CR-V}
> > 6a85f407 2023-08-13 lat/long 51.595623/-0.777015
> > f82cc194 2023-10-05 lat/long 51.832821/-2.556015
> {CDE-00598-C Honda Passport}
> > b9391c0c 2023-10-17 lat/long 51.809223/-2.379015
...
```

