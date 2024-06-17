{{ template "layout" . }}

{{ define "content" }}

    <form action="/contact" method="get" class="tool-bar">
        <label for="search">Search Term</label>
        <input id="search" type="search" name="q" value="{{ .Search }}" />
        <input type="submit" value="Search"/>
    </form>
    <table>
        <thead>
            <tr>
                <th>from Contact.html</th>
            </tr>
            </thead>
        <tbody>
            {{range .Contacts}}
                <tr>
                    <td>{{ .First }}</td>
                    <td>{{ .Last }}</td>
                    <td>{{ .Phone }}</td>
                    <td>{{ .Email }}</td>
                    <td>
                        <a href="/contact/{{ .ID }}/edit">Edit</a>
                        <a href="/contact/{{ .ID }}">View</a>
                    </td>
                </tr>
            {{ end }}
        </tbody>
    </table>
    <p>
        <a href="/contact/new">Add Contact</a>
    </p>
    
{{ end }}

