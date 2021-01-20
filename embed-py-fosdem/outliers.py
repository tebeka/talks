"""Detect outliers"""
import numpy as np


def detect(data):
    """Return indices where values more than 2 standard deviations from mean"""
    z_score = (data - data.mean()) / data.std()
    out = np.where(np.abs(z_score) > 3)  # Use the 68–95–99.7 rule
    # np.where returns a tuple for each dimension, we want the 1st element
    return out[0]
