<h1>Blog posts</h1>

<ul>
    {{range .blogs}}
    <li>
        <a href="/view/{{.Id}}"><b>{{.Title}}</b></a>
        <i>last update on {{.Mtime.Format "2006-01-02 15:04:05"}}</i>
        <a href="/edit/{{.Id}}">edit</a>
        <a href="/delete/{{.Id}}">delete</a>
    </li>
    {{end}}
</ul>
