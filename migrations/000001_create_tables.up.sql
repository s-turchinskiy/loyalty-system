CREATE SCHEMA IF NOT EXISTS loyaltysystem

CREATE TABLE IF NOT EXISTS loyaltysystem.statuses (
    id SERIAL PRIMARY KEY,
    status TEXT NOT NULL UNIQUE
);

INSERT INTO loyaltysystem.statuses (status) VALUES ('NEW');
INSERT INTO loyaltysystem.statuses (status) VALUES ('PROCESSING');
INSERT INTO loyaltysystem.statuses (status) VALUES ('INVALID');
INSERT INTO loyaltysystem.statuses (status) VALUES ('PROCESSED');

CREATE TABLE IF NOT EXISTS loyaltysystem.ordersStatuses (
    id SERIAL PRIMARY KEY,
    order_id TEXT NOT NULL UNIQUE,
    status_id INTEGER NOT NULL REFERENCES loyaltysystem.statuses (id),
    uploaded_at timestamptz DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS loyaltysystem.operations (
    id SERIAL PRIMARY KEY,
    operation TEXT NOT NULL UNIQUE
);

INSERT INTO loyaltysystem.operations (operation) VALUES ('REFILL');
INSERT INTO loyaltysystem.operations (operation) VALUES ('WITHDRAWAL');


CREATE TABLE IF NOT EXISTS loyaltysystem.balances (
    id SERIAL PRIMARY KEY,
    operation_id INTEGER NOT NULL REFERENCES loyaltysystem.operations (id),
    order_id TEXT NOT NULL UNIQUE,
    uploaded_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    sum numeric(15,2) NOT NULL
    );