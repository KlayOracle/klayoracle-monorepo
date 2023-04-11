-- noinspection SqlNoDataSourceInspectionForFile

DROP TABLE IF EXISTS node_feeds, node_jobs;

CREATE TABLE node_feeds (
    id UUID PRIMARY KEY,
    responses INT64[],
    round_answer INT64,
    adapter_id STRING(64),
    period TIMESTAMP
);

CREATE TABLE node_jobs (
    id UUID PRIMARY KEY,
    host_ip STRING(50),
    request JSONB,
    period TIMESTAMP
);