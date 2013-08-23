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
        </div>

        <form method="post">
            <label for="url">Url:</label><input name="url" /> 
            <input type="submit" value="Shorten">
        </form>

        <div class="footer">
            G'day
        </div>
	</body>
</html>
`

// vim: ft=html
