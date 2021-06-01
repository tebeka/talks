SELECT time, user, content
FROM journal
WHERE time >= "%s" AND time <= "%s";
