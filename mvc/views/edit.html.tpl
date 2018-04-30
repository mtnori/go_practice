{{ define "edit" }}
  {{ template "header" }}
    {{ template "menu" }}
   <h2>編集</h2>
    <form method="POST" action="update">
      <input type="hidden" name="uid" value="{{ .Id }}" />
      <label>名前:</label><input type="text" name="name" value="{{ .Name }}" /><br />
      <label>出身地:</label><input type="text" name="city" value="{{ .City }}" /><br />
      <input type="submit" value="更新" />
    </form><br />
  {{ template "footer" }}
{{ end }}
