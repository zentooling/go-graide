
{{ template "layout" . }}

{{ define "content" }}
<form action="/contact/new" method="post">
    <fieldset>
        <legend>Contact Values</legend>
        <div class="table rows">
            <p>
                <label for="email">Email</label>
                <input name="email" id="email" type="email" placeholder="Email" value="{{ .Email }}">
                <span class="error">{{ index .Errors "email"  }}</span>
            </p>
            <p>
                <label for="first_name">First Name</label>
                <input name="first_name" id="first_name" type="text" placeholder="First Name" value="{{ .First }}">
                <span class="error">{{ index .Errors "first"  }}</span>
            </p>
            <p>
                <label for="last_name">Last Name</label>
                <input name="last_name" id="last_name" type="text" placeholder="Last Name" value="{{ .Last }}">
                <span class="error">{{ index .Errors "last"  }}</span>
            </p>
            <p>
                <label for="phone">Phone</label>
                <input name="phone" id="phone" type="text" placeholder="Phone" value="{{ .Phone  }}">
                <span class="error">{{ index .Errors "phone"  }}</span>
            </p>
        </div>
        <button>Save</button>
    </fieldset>
</form>

<p>
    <a href="/contact">Back</a>
</p>

{{ end }}

