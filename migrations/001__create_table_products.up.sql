create table products
(
    id          BIGSERIAL NOT NULL,
    name        VARCHAR(255),
    description VARCHAR(255),
    price       DECIMAL,
    quantity    INTEGER
);
