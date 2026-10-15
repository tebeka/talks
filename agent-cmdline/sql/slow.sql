-- Median fare, tip, distance and duration per pickup zone, slow because of the ordered-set aggregates
SELECT
    pu_location_id,
    COUNT(*) AS trips,
    percentile_cont(0.5) WITHIN GROUP (ORDER BY fare_amount) AS median_fare,
    percentile_cont(0.5) WITHIN GROUP (ORDER BY tip_amount) AS median_tip,
    percentile_cont(0.5) WITHIN GROUP (ORDER BY trip_distance) AS median_distance,
    percentile_cont(0.5) WITHIN GROUP (ORDER BY dropoff_time - pickup_time) AS median_duration,
    COUNT(DISTINCT do_location_id) AS destinations
FROM trips
GROUP BY pu_location_id
ORDER BY trips DESC
LIMIT 10
;
