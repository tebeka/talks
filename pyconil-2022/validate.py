
    def validate(self):
        if not self.id:
            raise ValueError('missing id')
        if not self.login:
            raise ValueError('missing login')
        now = datetime.now(tz=timezone.utc)
        if self.created > now + timedelta(hours=1):
            raise ValueError('created is in the future')
