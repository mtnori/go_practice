{{ define "index" }}
  {{ template "header" }}
    {{ template "menu"  }}
    <h2>ユーザー一覧</h2>
    <table border="1">
      <thead>
        <tr>
          <td>ID</td>
          <td>Name</td>
          <td>City</td>
          <td colspan="3">操作</td>
        </tr>
      </thead>
      <tbody>
    {{ range . }}
      <tr>
        <td>{{ .Id }}</td>
        <td>{{ .Name }}</td>
        <td>{{ .City }}</td>
        <td><a href="/show?id={{ .Id }}">参照</a></td>
        <td><a href="/edit?id={{ .Id }}">編集</a></td>
        <td><a href="/delete?id={{ .Id }}">削除</a></td>
      </tr>
    {{ end }}
       </tbody>
    </table>
  {{ template "footer" }}
{{ end }}
