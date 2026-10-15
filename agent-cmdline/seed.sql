INSTALL postgres;
LOAD postgres;

ATTACH 'host=localhost port=5432 user=postgres dbname=postgres' AS pg (TYPE postgres);

INSERT INTO pg.public.trips
SELECT
    VendorID,
    tpep_pickup_datetime,
    tpep_dropoff_datetime,
    passenger_count,
    trip_distance,
    RatecodeID,
    store_and_fwd_flag,
    PULocationID,
    DOLocationID,
    payment_type,
    fare_amount,
    extra,
    mta_tax,
    tip_amount,
    tolls_amount,
    improvement_surcharge,
    total_amount,
    congestion_surcharge,
    Airport_fee,
    cbd_congestion_fee
FROM 'yellow_tripdata_2026-05.parquet';
