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
    orderId TEXT NOT NULL UNIQUE,
    statusId INTEGER NOT NULL REFERENCES loyaltysystem.statuses (Id),
    uploadedAt timestamptz DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS loyaltysystem.balances (
    id SERIAL PRIMARY KEY,
    orderId TEXT NOT NULL UNIQUE,
    uploadedAt timestamptz DEFAULT CURRENT_TIMESTAMP,
    sum numeric(15,2) NOT NULL
    );