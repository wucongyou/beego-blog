<h1>Blog posts</h1>

<ul>
    {{range .blogs}}
    <li>
        <a href="/view/{{.Id}}">{{.Title}}</a>
        last update on {{.Mtime}}
        <a href="/edit/{{.Id}}">Edit</a>
        <a href="/delete/{{.Id}}">Delete</a>
    </li>
    {{end}}
</ul>
