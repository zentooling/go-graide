{{define "layout"}}

<!doctype html>
<html lang="">

<head>
    <title>Contact App</title>
    <link rel="stylesheet" href="/web/static/css/missing.min.css">
    <link rel="stylesheet" href="/web/static/css/site.css">
    <!-- calculate integrity at https://www.srihash.org/ -->
    <script src="https://unpkg.com/htmx.org@1.9.12/dist/htmx.min.js" 
    integrity="sha384-ujb1lZYygJmzgSwoxRggbCHcjc0rB2XoQrxeTUQyRjrOnlCoYta87iKBWq3EsdM2" crossorigin="anonymous"></script>
    <link rel="icon" href="data:,">
</head>

<body hx-boost="true">
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