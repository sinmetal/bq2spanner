# bq2spanner
BigQueryに投げたQuery結果をSpannerに入るか試す

## Run

`row.go` を適当に編集する

```
export FROM_BIGQUERY_PROJECT_ID=hoge
export TO_SPANNER_PROJECT_ID=fuga
export TO_SPANNER_INSTANCE=moge
export TO_SPANNER_TABLE=momo
```