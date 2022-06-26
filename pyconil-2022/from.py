    @classmethod
    def from_json(cls, data):
        d = json.loads(data)
        u = cls(**d)
        return u
