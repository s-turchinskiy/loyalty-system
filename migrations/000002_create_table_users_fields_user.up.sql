CREATE TABLE IF NOT EXISTS loyaltysystem.users (
    id SERIAL PRIMARY KEY,
    login TEXT NOT NULL UNIQUE,
    hash TEXT NOT NULL,
    registration_at timestamptz DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO loyaltysystem.users (login, hash) VALUES ('test', '$2a$10$K37q5QUgbbJ9qSuhoj9FcecHIRHkh9gRykzNQUrIPXrI3EvyuzTOi');

ALTER TABLE loyaltysystem.ordersStatuses ADD COLUMN user_id INTEGER REFERENCES loyaltysystem.users (id);
ALTER TABLE loyaltysystem.balances ADD COLUMN user_id INTEGER REFERENCES loyaltysystem.users (id);


UPDATE loyaltysystem.ordersStatuses SET user_id = 1;
UPDATE loyaltysystem.balances SET user_id = 1;

ALTER TABLE loyaltysystem.ordersStatuses ALTER COLUMN user_id SET NOT NULL;
ALTER TABLE loyaltysystem.balances ALTER COLUMN user_id SET NOT NULL;