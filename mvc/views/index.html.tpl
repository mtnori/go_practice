{{ define "index" }}
  {{ template "header" }}
    {{ template "menu"  }}
    <h2>一覧</h2>
    <table border="1">
      <thead>
        <tr>
          <th>ID</th>
          <th>名前</th>
          <th>出身地</th>
          <th colspan="3">操作</th>
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
