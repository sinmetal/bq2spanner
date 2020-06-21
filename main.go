package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/spanner"
)

type Row struct {
}

func main() {
	var q string
	q = ReadFile()
	fmt.Println(q)

	bqProjectID, spannerProjectID := GetProjectID()
	ctx := context.Background()

	bq, err := bigquery.NewClient(ctx, bqProjectID)
	if err != nil {
		panic(err)
	}
	iter, err := bq.Query(q).Read(ctx)
	if err != nil {
		panic(err)
	}
	var row Row
	if err := iter.Next(&row); err != nil {
		panic(err)
	}

	spa, err := spanner.NewClient(ctx, fmt.Sprintf("projects/%s/instances/%s/databases/%s", spannerProjectID, "test20200621", "sinmetal"))
	if err != nil {
		panic(err)
	}

	m, err := spanner.InsertStruct("TABLE", &row)
	if err != nil {
		panic(err)
	}
	ms := []*spanner.Mutation{
		m,
	}

	_, err = spa.Apply(ctx, ms)
	if err != nil {
		panic(err)
	}
	fmt.Println("DONE")
}

func ReadFile() string {
	b, err := ioutil.ReadFile("q.sql")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func GetProjectID() (string, string) {
	bqProjectID := os.Getenv("FROM_BIGQUERY_PROJECT_ID")
	spannerProjectID := os.Getenv("TO_SPANNER_PROJECT_ID")

	return bqProjectID, spannerProjectID
}
