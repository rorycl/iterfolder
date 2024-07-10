// dbCompoundExample2 is an example of an iterator function reading
// postgres row-wise using pgx.ForEachRow which is then fed to
// IterFolder.
//
// How to exit safely from the pgx.ForEachRow line is not clear.
package main

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rorycl/iterfolder"
)

type Person struct {
	ID          int `db:"personid"`
	Firstname   string
	Lastname    string
	CarCount    int `db:"person_car_count"`
	TicketCount int `db:"person_ticket_count"`
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s cars %d tickets %d", p.Firstname, p.Lastname, p.CarCount, p.TicketCount)
}

type Car struct {
	Registration string
	Manufacturer string
	Model        string
}

type Ticket struct {
	ID   string `db:"uuidpart"`
	Date time.Time
	Lat  float32
	Long float32
}

func (t Ticket) String() string {
	return fmt.Sprintf("%s %s lat/long %0.6f/%0.6f", t.ID, t.Date.Format("2006-01-02"), t.Lat, t.Long)
}

type PersonCarTicket struct {
	Person
	Car
	Ticket
}

var query = `
SELECT 
    personid
    ,firstname
    ,lastname
    ,ticket_count as person_ticket_count
    ,sum(pcc) over (partition by personid) as person_car_count
    ,registration
    ,manufacturer
    ,model
    ,uuidpart
    ,date
    ,lat
    ,long
FROM (
    SELECT 
        -- person
        p.id AS personid
        ,p.firstname
        ,p.lastname
        -- car
        ,c.registration
        ,c.manufacturer
        ,c.model
        -- ticket (showing just the left part of the uuid)
        ,left(t.uuid::text, 8) AS uuidpart
        ,date(t.datetime) as date
        ,lat
        ,long
        -- 
        ,count(t.uuid) OVER (PARTITION BY p.id) AS ticket_count
        ,CASE
         WHEN row_number() OVER (PARTITION BY p.id, c.registration) = 1 THEN 1
         ELSE 0
         END AS pcc 
    FROM
        people p
        JOIN cars c ON c.owner = p.id
        join tickets t ON t.car = c.registration
    ) x
ORDER BY
    ticket_count
    ,personid
    ,registration
;
`

func main() {
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, "postgres://iterf:iterf@localhost:5432/iterf")
	if err != nil {
		exit("failed creating pgx pool", err)
	}
	defer dbpool.Close()

	rows, err := dbpool.Query(ctx, query)
	if err != nil {
		exit("failed query", err)
	}

	// construct input iterator to send to iterFolder using pgx.ForEachRow
	type PCT = iterfolder.ABC[Person, Car, Ticket]
	ctsIter := func() iter.Seq[PCT] {
		return func(yield func(PCT) bool) {
			person := Person{}
			car := Car{}
			ticket := Ticket{}
			columnTypes := []any{
				&person.ID,
				&person.Firstname,
				&person.Lastname,
				&person.CarCount,
				&person.TicketCount,
				&car.Registration,
				&car.Manufacturer,
				&car.Model,
				&ticket.ID,
				&ticket.Date,
				&ticket.Lat,
				&ticket.Long,
			}
			_, err := pgx.ForEachRow(rows, columnTypes, func() error {
				if !yield(PCT{person, car, ticket}) {
					return errors.New("returned false from yield")
				}
				return nil
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	// run the folder
	for a := range iterfolder.IterFolder[Person, Car, Ticket](ctsIter) {
		fmt.Println(a.This)
		for b := range a.Iter() {
			fmt.Println(">", b.This)
			for c := range b.Iter() {
				fmt.Println("> >", c)
			}
		}
	}
}

func exit(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
	os.Exit(1)
}
