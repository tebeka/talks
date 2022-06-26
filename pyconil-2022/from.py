    @classmethod
    def from_json(cls, data):
        d = json.loads(data)
        u = cls(**d)
        u.validate()
        return u
