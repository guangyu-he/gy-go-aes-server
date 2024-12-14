package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

const PGURL string = "postgres://test:test@localhost:5432/test"

type TestGo struct {
	ID     int
	Name   string
	Weight int
}

func (tg *TestGo) Insert() error {
	conn, err := pgx.Connect(context.Background(), PGURL)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	var id int
	_, err = conn.Exec(context.Background(), "INSERT INTO testgo (name, weight) VALUES ($1, $2)", tg.Name, tg.Weight)
	if err != nil {
		return err
	}

	err = conn.QueryRow(context.Background(), "SELECT id FROM testgo WHERE name=$1 AND weight=$2", tg.Name, tg.Weight).Scan(&id)
	if err != nil {
		return err
	}
	tg.ID = id

	return err
}

func (tg *TestGo) Update() error {
	conn, err := pgx.Connect(context.Background(), PGURL)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "UPDATE testgo SET name=$1, weight=$2 WHERE id=$3", tg.Name, tg.Weight, tg.ID)
	if err != nil {
		return err
	}
	return nil
}

func (tg *TestGo) Delete() error {
	conn, err := pgx.Connect(context.Background(), PGURL)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "DELETE FROM testgo WHERE id=$1", tg.ID)
	if err != nil {
		return err
	}
	return nil
}

func (tg *TestGo) Select() error {
	conn, err := pgx.Connect(context.Background(), PGURL)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(), "SELECT id, name, weight FROM testgo WHERE id=$1", tg.ID).Scan(&tg.ID, &tg.Name, &tg.Weight)
	if err != nil {
		return err
	}

	tg.ID = tg.ID
	tg.Name = tg.Name
	tg.Weight = tg.Weight

	return nil
}

func PGCreateTable() {

	conn, err := pgx.Connect(context.Background(), PGURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// create a table called testgo
	_, err = conn.Exec(context.Background(), "CREATE TABLE testgo (id serial PRIMARY KEY, name VARCHAR(50), weight INT)")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Table created")
	}
}

func PG() {

	//var testgo TestGo = TestGo{
	//	Name:   "test",
	//	Weight: 100,
	//}
	//
	//err := testgo.Insert()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(testgo.ID, testgo.Name, testgo.Weight)

	var testgo TestGo = TestGo{
		ID: 4,
	}

	err := testgo.Select()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(testgo.ID, testgo.Name, testgo.Weight)

	testgo.Name = "test2"

	err = testgo.Update()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(testgo.ID, testgo.Name, testgo.Weight)

}
