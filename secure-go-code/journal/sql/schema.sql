CREATE TABLE IF NOT EXISTS journal (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    time TIMESTAMP,
    user TEXT,
    content TEXT
);
