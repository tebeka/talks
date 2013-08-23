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

        <form method="post">
            <label for="url">Url:</label><input name="url" /> 
            <input type="submit" value="Shorten">
        </form>

        <div class="footer">
            G'day {{ .User }}
        </div>
	</body>
</html>
`

// vim: ft=html
