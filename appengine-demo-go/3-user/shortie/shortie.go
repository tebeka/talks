/* AppEngine demo - a URL shortener */
package shortie

import (
	"html/template"
	"net/http"

	"appengine"
	"appengine/user"
)

var homeTemplate *template.Template

func init() {
	homeTemplate = template.Must(template.New("home").Parse(homeHTML))
	http.HandleFunc("/", rootHandler)
}

// Parameters for homeTemplate.
type homeParams struct {
	User       string
	LoginTitle string
	LoginURL   string
	Error      string
}

// fillUser fills user details in template parameters
func fillUser(r *http.Request, ctx appengine.Context, params *homeParams) error {
	var err error
	u := user.Current(ctx)

	if u != nil {
		params.User = u.String()
		params.LoginTitle = "Logout"
		params.LoginURL, err = user.LogoutURL(ctx, r.URL.String())
	} else {
		params.User = "Stranger"
		params.LoginTitle = "Login"
		params.LoginURL, err = user.LoginURL(ctx, r.URL.String())
	}

	return err
}

// rootHandler handles the main page.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	var err error
	params := new(homeParams)

	// Run at end. We check "err" and update params if needed. Then serve homeTemplate.
	defer func() {
		if err != nil {
			params.Error = err.Error()
			ctx.Errorf("%v", err) // Log error
		}
		homeTemplate.Execute(w, params)
	}()

	err = fillUser(r, ctx, params)
}
