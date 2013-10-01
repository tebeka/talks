Session = sessionmaker(bind=engine)

session = Session()
session.add_all([Point(1, 1), Point(2, 2)])
session.commit()

p1 = session.query(Point).filter_by(x=1).first()
assert p1.y == 1, 'corrupted data'
