/* AppEngine demo - a URL shortener */
package shortie

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/delay"
	"appengine/memcache"
	"appengine/user"
)

const (
	counterKeyName = "counter-key-name"
	counterKind    = "Counter"
	urlKind        = "Url"
)

// Global counter of urls
type Counter struct {
	Count int64
}

// URL stored in database
type URL struct {
	Short   string
	Long    string
	User    string
	Created time.Time
	Hits    int64
}

// Parameters for homeTemplate.
type homeParams struct {
	User       string
	LoginTitle string
	LoginURL   string
	Error      string
	Count      int64
	ShortURL   string
}

var homeTemplate *template.Template

func init() {
	homeTemplate = template.Must(template.New("home").Parse(homeHTML))
	http.HandleFunc("/", rootHandler)
}

// rootHandler handles the main page.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// If it has something, we assume it's a short url
	if r.URL.Path != "/" {
		redirectHandler(w, r)
		return
	}

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
	if err != nil {
		return
	}
	params.Count, err = urlCount(ctx)

	if r.Method == "POST" {
		longURL := strings.TrimSpace(r.FormValue("url"))

		if !isValidURL(longURL) {
			err = fmt.Errorf("Bad URL - %s", longURL)
			return
		}

		if !hasSchema(longURL) {
			longURL = fmt.Sprintf("http://%s", longURL)
		}

		var id string
		id, err = newShortURL(ctx, longURL, params.User)
		if err != nil {
			return
		}

		params.ShortURL = fullURL(r, id)
	}
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

// urlCount return the current count of urls.
func urlCount(ctx appengine.Context) (int64, error) {
	key := datastore.NewKey(ctx, counterKind, counterKeyName, 0, nil)
	counter := new(Counter)
	if err := datastore.Get(ctx, key, counter); err != nil && err != datastore.ErrNoSuchEntity {
		return 0, err
	}

	return counter.Count, nil
}

/* nextId returns the next short url. 
We use the global counter and then encode the last count in base62.
*/
func nextId(ctx appengine.Context) (string, error) {
	var count int64

	err := datastore.RunInTransaction(ctx, func(ctx appengine.Context) error {
		key := datastore.NewKey(ctx, counterKind, counterKeyName, 0, nil)
		counter := new(Counter)
		if err := datastore.Get(ctx, key, counter); err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}

		counter.Count++
		if _, err := datastore.Put(ctx, key, counter); err != nil {
			return err
		}
		count = counter.Count
		return nil
	}, nil)

	return base62Encode(uint64(count)), err
}

// isValidURL check that URL is valid
func isValidURL(url string) bool {
	return (len(url) > 0) && strings.Contains(url, ".")
}

// hasSchema check if url has schema prefix.
func hasSchema(url string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z]+://", url)
	return match
}

/* fullURL adds http://<host> suffix to short url. 
This works both locally and on AppEngine.
*/
func fullURL(r *http.Request, id string) string {
	return fmt.Sprintf("http://%s/%s", r.Host, id)
}

func newShortURL(ctx appengine.Context, longURL, user string) (string, error) {
	var id string
	id, err := nextId(ctx)

	if err != nil {
		return "", err
	}

	url := &URL{
		Short:   id,
		Long:    longURL,
		User:    user,
		Created: time.Now(),
		Hits:    0,
	}
	key := datastore.NewKey(ctx, urlKind, id, 0, nil)
	_, err = datastore.Put(ctx, key, url)
	if err != nil {
		return "", err
	}

	return id, nil
}

// delayedInc is a "delayed" call to incHits.
var delayedInc = delay.Func("hits", incHits)

/* redirectHandler handles redirects.
All urls that are not / and worker are assumed to be redirects (short).
*/
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	path := r.URL.Path[1:]

	var longURL string
	// Try memcache first and if not get URL from datastore and updatge memcache
	if item, err := memcache.Get(ctx, path); err == memcache.ErrCacheMiss {
		url, err1 := getURL(ctx, path)
		if err1 != nil {
			ctx.Errorf("Short URL not found - %s", path)
			http.NotFound(w, r)
			return
		}
		longURL = url.Long

		item1 := &memcache.Item{
			Key:   path,
			Value: []byte(longURL),
		}

		if err1 := memcache.Set(ctx, item1); err1 != nil {
			ctx.Errorf("memcache setting error: %v", err1)
		}
	} else if err != nil {
		ctx.Errorf("memcache error - %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else { // Found in memcache
		longURL = string(item.Value)
	}

	// Offload hit counter update to a task
	delayedInc.Call(ctx, path)

	http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
}

// getURL fetches a URL from the datastore by short url.
func getURL(ctx appengine.Context, short string) (*URL, error) {
	key := urlKey(ctx, short)
	url := new(URL)
	err := datastore.Get(ctx, key, url)

	return url, err
}

// urlKey returns a datastore key for short url.
func urlKey(ctx appengine.Context, short string) *datastore.Key {
	return datastore.NewKey(ctx, urlKind, short, 0, nil)
}

// incHits increments hit count on url (this is done when short url is resolved).
func incHits(ctx appengine.Context, short string) error {
	return datastore.RunInTransaction(ctx, func(ctx appengine.Context) error {
		url, err := getURL(ctx, short)
		if err != nil {
			return err
		}

		url.Hits++
		key := urlKey(ctx, short)
		if _, err := datastore.Put(ctx, key, url); err != nil {
			return err
		}
		return nil
	}, nil)
}
