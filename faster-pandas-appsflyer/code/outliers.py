def find_outliers(data):
    """Find outliers in data, return indices of outliers"""
    out = data[(data - data.mean()).abs() > 2 * data.std()]
    return out.index
