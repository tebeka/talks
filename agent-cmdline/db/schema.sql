-- Codes from https://www.nyc.gov/assets/tlc/downloads/pdf/data_dictionary_trip_records_yellow.pdf

CREATE TABLE vendor (
    id INT PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO vendor VALUES
    (1, 'Creative Mobile Technologies, LLC'),
    (2, 'Curb Mobility, LLC'),
    (6, 'Myle Technologies Inc'),
    (7, 'Helix');

CREATE TABLE rate_code (
    id INT PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO rate_code VALUES
    (1, 'Standard rate'),
    (2, 'JFK'),
    (3, 'Newark'),
    (4, 'Nassau or Westchester'),
    (5, 'Negotiated fare'),
    (6, 'Group ride'),
    (99, 'Null/unknown');

CREATE TABLE payment_type (
    id INT PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO payment_type VALUES
    (0, 'Flex Fare trip'),
    (1, 'Credit card'),
    (2, 'Cash'),
    (3, 'No charge'),
    (4, 'Dispute'),
    (5, 'Unknown'),
    (6, 'Voided trip');

CREATE TABLE store_and_fwd (
    flag CHAR(1) PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO store_and_fwd VALUES
    ('Y', 'store and forward trip'),
    ('N', 'not a store and forward trip');

CREATE TABLE trips (
    vendor_id INT REFERENCES vendor(id),
    pickup_time TIMESTAMP,
    dropoff_time TIMESTAMP,
    passenger_count INT,
    trip_distance DOUBLE PRECISION,
    rate_code_id INT REFERENCES rate_code(id),
    store_and_fwd_flag CHAR(1) REFERENCES store_and_fwd(flag),
    pu_location_id INT,
    do_location_id INT,
    payment_type_id INT REFERENCES payment_type(id),
    fare_amount NUMERIC(10, 2),
    extra NUMERIC(10, 2),
    mta_tax NUMERIC(10, 2),
    tip_amount NUMERIC(10, 2),
    tolls_amount NUMERIC(10, 2),
    improvement_surcharge NUMERIC(10, 2),
    total_amount NUMERIC(10, 2),
    congestion_surcharge NUMERIC(10, 2),
    airport_fee NUMERIC(10, 2),
    cbd_congestion_fee NUMERIC(10, 2)
);
