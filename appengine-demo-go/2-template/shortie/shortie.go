/* AppEngine demo - a URL shortener */
package shortie

import (
	"html/template"
	"net/http"
)

var homeTemplate *template.Template

func init() {
	homeTemplate = template.Must(template.New("home").Parse(homeHTML))
	http.HandleFunc("/", rootHandler)
}

// rootHandler handles the main page.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, nil)
}
