{{ define "show" }}
  {{ template "header" }}
    {{ template "menu" }}
    <h2>ID: {{ .Id }}</h2>
    <p>名前:{{ .Name }}</p>
    <p>出身地:{{ .City }}</p><br /><a href="/edit?id={{ .Id }}">編集</a></p>
  {{ template "footer" }}
{{ end }}
