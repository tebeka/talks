from flask import Flask, render_template
from flask.ext.sqlalchemy import SQLAlchemy

from os.path import realpath, dirname, join

app = Flask(__name__)

here = dirname(realpath(__file__))
db_file = join(here, 'db.sqlite')

app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///{}'.format(db_file)
db = SQLAlchemy(app)


class Link(db.Model):
    short = db.Column(db.Text, primary_key=True)
    long = db.Column(db.Text)
    hits = db.Column(db.Integer)

    def __init__(self, short, long):
        self.short = short
        self.long = long
        self.hits = 0

    def __repr__(self):
        return '<Link> {short}->{long} [{hits}]'.format(**self.__dict__)


@app.route('/')
def index():
    link_count = db.session.query(Link).count()
    return render_template('index.html', link_count=link_count)


if __name__ == '__main__':
    db.create_all()
    app.run(debug=True)
