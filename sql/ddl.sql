CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    customer VARCHAR(255),
    quantity INT,
    price NUMERIC(10, 2),
    timestamp TIMESTAMP
);
