**[gophersnippets](https://gophersnippets.com)** is a collection of code snippets with tests and testable examples for the Go programming language. 

Each snippet:
 - is small and shows how to achieve something specific with Go 
 - contains at least one test or testable example
 - runs on [Go playground](https://play.golang.org/) since it now supports tests and testable examples  

{{ $repo := .Repo}}
{{range $category := .Categories}}
## {{$category}}
{{- range $snippet := index $.Snippets $category }}
 - [{{$snippet.Title}}](https://gophersnippets.com/{{$snippet.DirectoryName}})
{{- end -}}
{{end}}

## Comments

https://gophersnippets.com uses Github issues for comments, powered by https://utteranc.es/. You can add comments directly on the corresponding Github issue or use the comment form under each code section. 

## License
<a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/">Creative Commons Attribution-ShareAlike 4.0 International License</a>.
