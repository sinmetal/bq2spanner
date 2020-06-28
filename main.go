package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/spanner"
)

type Config struct {
	FromBigQueryProjectID string
	ToSpannerProjectID    string
	ToSpannerInstance     string
	ToSpannerTableName    string
}

func main() {
	var q string
	q = ReadFile()
	fmt.Println(q)

	config := GetConfig()
	fmt.Printf("%+v\n", config)

	ctx := context.Background()

	bq, err := bigquery.NewClient(ctx, config.FromBigQueryProjectID)
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

	spa, err := spanner.NewClientWithConfig(ctx, fmt.Sprintf("projects/%s/instances/%s/databases/%s", config.ToSpannerProjectID, config.ToSpannerInstance, "sinmetal"),
		spanner.ClientConfig{
			SessionPoolConfig: spanner.SessionPoolConfig{
				MinOpened: 1,
				MaxOpened: 10,
			}})
	if err != nil {
		panic(err)
	}

	m, err := spanner.InsertStruct(config.ToSpannerTableName, &row)
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
	b, err := ioutil.ReadFile("query.sql")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func GetConfig() *Config {
	bqProjectID := os.Getenv("FROM_BIGQUERY_PROJECT_ID")
	spannerProjectID := os.Getenv("TO_SPANNER_PROJECT_ID")
	spannerInstance := os.Getenv("TO_SPANNER_INSTANCE")
	spannerTable := os.Getenv("TO_SPANNER_TABLE")

	return &Config{
		FromBigQueryProjectID: bqProjectID,
		ToSpannerProjectID:    spannerProjectID,
		ToSpannerInstance:     spannerInstance,
		ToSpannerTableName:    spannerTable,
	}
}
