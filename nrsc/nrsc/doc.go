/* Serving resource from zip file appended to Go executable (this enables on file deploy).

Making it happen:
	1. Add code to serve resources (see example below)
	2. Compile your executable
	3. Run "nrsc /path/to/exe /path/to/resources"
	4. Deploy

Example code:

	package main

	import (
		"fmt"
		"net/http"
		"os"

		"bitbucket.org/tebeka/nrsc"
	)

	type params struct {
		Number  int
	}

	func indexHandler(w http.ResponseWriter, req *http.Request) {
		t, err := nrsc.LoadTemplates(nil, "t.html")
		if err != nil {
			http.NotFound(w, req)
		}
		if err = t.Execute(w, params{7}); err != nil {
			http.NotFound(w, req)
		}
	}

	func main() {
		nrsc.Handle("/static/")
		nrsc.Mask(".tmpl$")
		http.HandleFunc("/", indexHandler)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	}
*/
package nrsc
