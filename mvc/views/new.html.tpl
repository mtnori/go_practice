{{ define "new" }}
  {{ template "header" }}
    {{ template "menu" }}
   <h2>登録</h2>
    <form method="POST" action="insert">
      <label>名前:</label><input type="text" name="name" /><br />
      <label>出身地:</label><input type="text" name="city" /><br />
      <input type="submit" value="登録" />
    </form>
  {{ template "footer" }}
{{ end }}
