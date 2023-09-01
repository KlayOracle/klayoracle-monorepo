-- noinspection SqlNoDataSourceInspectionForFile

DROP DATABASE IF EXISTS defaultdb;
DROP DATABASE IF EXISTS postgres;

CREATE DATABASE IF NOT EXISTS node1;
CREATE DATABASE IF NOT EXISTS node2;
CREATE DATABASE IF NOT EXISTS node3;

DROP TABLE IF EXISTS node1.node_responses, node1.node_jobs;
DROP TABLE IF EXISTS node2.node_responses, node2.node_jobs;
DROP TABLE IF EXISTS node3.node_responses, node3.node_jobs;

CREATE TABLE node1.node_responses (
                                      id UUID PRIMARY KEY,
                                      responses INT64[],
                                      round_answer INT64,
                                      adapter_id STRING(66),
                                      period TIMESTAMP
);

CREATE TABLE node1.node_jobs (
                                 id UUID PRIMARY KEY,
                                 adapter_id STRING(66),
                                 oracle_address STRING(42),
                                 request JSONB,
                                 period TIMESTAMP
);

CREATE TABLE node2.node_responses (
                                      id UUID PRIMARY KEY,
                                      responses INT64[],
                                      round_answer INT64,
                                      adapter_id STRING(66),
                                      period TIMESTAMP
);

CREATE TABLE node2.node_jobs (
                                 id UUID PRIMARY KEY,
                                 adapter_id STRING(66),
                                 oracle_address STRING(42),
                                 request JSONB,
                                 period TIMESTAMP
);

CREATE TABLE node3.node_responses (
                                      id UUID PRIMARY KEY,
                                      responses INT64[],
                                      round_answer INT64,
                                      adapter_id STRING(66),
                                      period TIMESTAMP
);

CREATE TABLE node3.node_jobs (
                                 id UUID PRIMARY KEY,
                                 adapter_id STRING(66),
                                 oracle_address STRING(42),
                                 request JSONB,
                                 period TIMESTAMP
);