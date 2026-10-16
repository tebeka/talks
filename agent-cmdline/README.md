# Command Line Tools for Humans and AI

This is code from the talk.
I made the code as simple as possible, some (many?) things are omitted for brevity.

In order to run it you'll first need to run the database.
For that you'll need docker, duckdb and to download [yellow_tripdata_2026-05.parquet][pr]
from the [NYC Trip Record Data][nyc].
The run `make db-start`.

[nyc]: https://www.nyc.gov/site/tlc/about/tlc-trip-record-data.page
[pr]: https://d37ci6vzurychx.cloudfront.net/trip-data/yellow_tripdata_2026-05.parquet

## Code

Each step is in it's own directory. See [playbook.md][pl] for the order and how to run commands.
You can also diff each step, for example    

```
$ diff tty initial 
diff --color tty/main.go initial/main.go
13d12
< 	"golang.org/x/term"
98,99d96
< var isTTY = term.IsTerminal(int(os.Stdout.Fd()))
< 
101,104d97
< 	if !isTTY {
< 		return s
< 	}
< 
```

I tried to make the diffs between steps as small as possible.

[pl]: ./playbook.md

## License

[CC BY 4.0](LICENSE) © Miki Tebeka

---
Keep hacking, stay curious,
Miki
