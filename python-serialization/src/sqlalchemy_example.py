from sqlalchemy import create_engine, Column, Integer
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

engine = create_engine('sqlite:///:memory:')
Base = declarative_base(engine)

class Point(Base):
    __tablename__ = 'points'

    id = Column(Integer, primary_key=True)
    x = Column(Integer)
    y = Column(Integer)

    def __init__(self, x, y):
        self.x, self.y = x, y

Base.metadata.create_all()
