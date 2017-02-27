<h1>Blog posts</h1>

<ul>
    {{range .blogs}}
    <li>
        <a href="/view/{{.Id}}">{{.Title}}</a>
        last update on {{.Mtime.Format "2006-01-02 15:04:05"}}
        <a href="/edit/{{.Id}}">Edit</a>
        <a href="/delete/{{.Id}}">Delete</a>
    </li>
    {{end}}
</ul>
