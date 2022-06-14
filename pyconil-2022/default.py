def default(obj):
    if isinstance(obj, datetime):
        return obj.isoformat()
    return obj
