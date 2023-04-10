-- noinspection SqlNoDataSourceInspectionForFile

DROP TABLE IF EXISTS node_feeds;

CREATE TABLE node_feeds (
    id UUID PRIMARY KEY,
    responses INT64[],
    round_answer INT64,
    adapter_id STRING(64),
    period TIMESTAMP
);