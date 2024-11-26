import tracemalloc


class Car:
    def __init__(self, lat, lng):
        self.lat = lat
        self.lng = lng
        

tracemalloc.start()
snap1 = tracemalloc.take_snapshot()
cars = [Car(2.0686173, 34.7912133) for _ in range(1_000_000)]
snap2 = tracemalloc.take_snapshot()
stats = snap2.compare_to(snap1, 'lineno')
for stat in stats[:10]:
    print(stat)
