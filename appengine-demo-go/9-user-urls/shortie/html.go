package shortie

const homeHTML = `
<html>
    <head>
        <title>Shortie - The URL Shortener</title>
        <link rel="stylesheet" href="/static/style.css" />
        <link rel="shortcut icon" href="/static/logo.png" />
    </head>
    <body>
        <div class="header">
            The URL Shortener
            <span class="right">
				<a href="{{ .LoginURL }}">{{ .LoginTitle }}</a>
            </span>
        </div>

        {{ if .Error }}
        <div class="error">
            ERROR: {{ .Error | html}}
        </div>
        {{ end }}

        {{ if .ShortURL }}
        <div class="short">
            Url shortened to <a href="{{ .ShortURL }}">{{ .ShortURL }}</a>
        </div>
        {{ end }}

        <form method="post">
            <label for="url">Url:</label><input name="url" /> 
            <input type="submit" value="Shorten">
        </form>

		{{ with .URLs }}
        <b>Your Urls</b>
        <table class="urls">
            <tr><th>URL</th><th>Created</th><th>Hits</th></tr>
			{{ range $idx, $url := . }}
				<tr class="{{ $idx | oddeven }}">
					<td title="{{ .Long | html }}">
						<a href="{{ .Long }}">{{ .Short }}</a>
					</td>
					<td>{{ .Created | strftime }}</td>
					<td>{{ .Hits }}</td>
            </tr>
            {{ end }}
        </table>
        {{ end }}

        <div class="footer">
            G'day {{ .User }}, we've shortened {{ .Count }} urls so far.
        </div>
	</body>
</html>
`

// vim: ft=html
