deg_km = 111


def distance(lat1, lng1, lat2, lng2):
    """Return approximate distance between two coordinates in Km"""
    # Manhattan distance
    dlat = abs(lat1 - lat2)
    dlng = abs(lng1 - lng2)
    return (dlat + dlng) * deg_km


dist = distance(32.5253837, 34.9408283, 32.517574, 34.9455938)
print(dist)
