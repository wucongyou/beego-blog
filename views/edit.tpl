<h1>Edit</h1>
<form action="" method="post">
    Title: <input type="text" name="title" value="{{.Post.Title}}"><br>
    Content: <textarea name="content" colspan="3" rowspan="10">{{.Post.Content}}</textarea>
    <input type="hidden" name="id" value="{{.Post.Id}}">
    <br/>
    <input type="submit">
</form>
