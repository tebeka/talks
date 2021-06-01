package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	entries = []Entry{
		{date(2021, 4, 1), "joe", "Who's on first?"},
		{date(2021, 4, 2), "joe", "What's on second?"},
		{date(2021, 4, 3), "joe", "I don't know on third!"},
	}
)

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func TestDB(t *testing.T) {
	require := require.New(t)

	dir := t.TempDir()
	dsn := fmt.Sprintf("%s/journal.db", dir)

	db, err := NewDB(dsn)
	require.NoError(err)
	defer db.Close()

	for _, e := range entries {
		err := db.Add(e)
		require.NoErrorf(err, "insert %#v", e)
	}

	start := entries[0].Time
	end := entries[1].Time

	entries, err := db.Query(start, end)
	require.NoError(err, "query")
	require.Equal(2, len(entries), "size")
}
