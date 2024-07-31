{{define "layout"}}

<!doctype html>
<html lang="">

<head>
    <title>Contact App</title>
    <link rel="stylesheet" href="/web/static/css/missing.min.css">
    <link rel="stylesheet" href="/web/static/css/site.css">
    <!-- calculate integrity at https://www.srihash.org/ -->
    <script src="https://unpkg.com/htmx.org@1.9.12/dist/htmx.min.js"
        integrity="sha384-ujb1lZYygJmzgSwoxRggbCHcjc0rB2XoQrxeTUQyRjrOnlCoYta87iKBWq3EsdM2"
        crossorigin="anonymous"></script>
    <link rel="icon" href="data:,">
    <!-- <div id="flash">
        <script>

            function flashRemove() {
                const flash = document.getElementById('flash');
                // const flash = document.querySelector('#flash');
                if (flash && flash != null) { flash.remove() }
                timeoutFunc()
            }
            function timeoutFunc() {
                setTimeout(flashRemove, 5000)

            }

            timeoutFunc()
            //    setTimeout(() => {
            //         //flash.textContent = '';
            //         flash.remove();
            //     }, 3000);
            // setTimeout(fade_out, 5000);
            // function fade_out() {
            //     $("#flash").fadeOut().empty();
            // }
        </script> -->
    <div class="dummy-wrap">
        <div class="success-wrap successfully-saved">
            <ul class="flash">
                <li>This is a longer message</li>
                <li>Message2</li>
                <li>Message3</li>
            </ul>
        </div>
    </div>
</head>

<body hx-boost="true" id="layout-body">
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