{{define "layout"}}

<!doctype html>
<html lang="">
    <head>
        <title>Contact App</title>
        <link rel="stylesheet" href="/web/static/css/missing.min.css">
        <link rel="stylesheet" href="/web/static/css/site.css">
        <link rel="icon" href="data:,">
    </head>
    <body>
        <main>
            <header>
                <h1>
                    <span style="text-transform:uppercase;">graide.app</span>
                    <sub-title>A Grading Application</sub-title>
                </h1>
            </header>
            {{/* content block defined in each 'master' html file */}}
            {{ template "content" .}}
        </main>
    </body>
</html>

{{ end }}