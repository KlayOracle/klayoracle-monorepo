-- noinspection SqlNoDataSourceInspectionForFile

DROP TABLE IF EXISTS node_responses, node_jobs;

CREATE TABLE node_responses (
    id UUID PRIMARY KEY,
    responses INT64[],
    round_answer INT64,
    adapter_id STRING(66),
    period TIMESTAMP
);

CREATE TABLE node_jobs (
    id UUID PRIMARY KEY,
    adapter_id STRING(66),
    request JSONB,
    period TIMESTAMP
);