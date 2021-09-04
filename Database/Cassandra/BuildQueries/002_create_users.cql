CREATE TABLE IF NOT EXISTS users
(
    id             UUID,
    phone          VARCHAR,
    balance        INT,
    driver_details driver,
    created_at     TIMESTAMP,
    PRIMARY KEY ( id )
);