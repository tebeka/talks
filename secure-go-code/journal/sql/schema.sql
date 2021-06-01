CREATE TABLE IF NOT EXISTS journal (
    time TIMESTAMP,
    user TEXT,
    content TEXT
);
CREATE INDEX IF NOT EXISTS entries_time ON journal(time);
