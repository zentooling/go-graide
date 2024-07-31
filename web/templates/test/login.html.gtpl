
{{ template "layout" . }}

{{ define "content" }}

<form action="/login" method="post">
    <fieldset>
        <legend>Enter Username and Password to continue</legend>
        <p>
            <label for="user_name">User</label>
            <input name="user_name" id="user_name" type="text" value="">
        </p>
        <p>
            <label for="password">Password</label>
            <input name="password" id="password" type="password" value="">
        </p>
        <button>Login</button>
    </fieldset>
</form>

{{ end }}