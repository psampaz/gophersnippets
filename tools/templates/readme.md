**[gophersnippets](https://gophersnippets.com)** is a collection of code snippets with tests and testable examples for the Go programming language. 

Each snippet:
 - is small and shows how to achieve something specific with Go 
 - contains at least one test or testable example
 - runs on [Go playground](https://play.golang.org/) since it now supports tests and testable examples  

## Snippets
{{ $repo := .Repo}}
{{range $category := .Categories}}
### {{$category}}
{{- range $snippet := index $.Snippets $category }}
 - [{{$snippet.Title}}](https://gophersnippets.com/{{$snippet.DirectoryName}})
{{- end -}}
{{end}}

## Comments

https://gophersnippets.com uses Github issues for comments, powered by https://utteranc.es/. You can add comments directly on the corresponding Github issue or use the comment form under each code section. 

## Contributing

You are welcome to contribute testable Go code snippets.

Please read the following before submitting a PR:
- https://blog.golang.org/examples
- https://blog.golang.org/playground
- https://github.com/golang/go/wiki/TableDrivenTests

### Guidelines:

- Include only one code snippet per PR
- The snippet should be small and solve one problem only
- Include comments when necessary
- When you submit a PR the following checks will run automatically:
  - code includes at least 1 test or testable example
  - code runs successfully on play.golang.org
  - code passes go vet
  - code includes a category and a title in the first line of the code (Category: Title)
    ```go
    // Strings: How to reverse a string
    // package main
    ```
    or includes a category and title in the first line of the code, an empty comment line and then a description
    ```go
    // Strings: How to reverse a string
    //
    // Description
    // package main
    ```
  - code should be included in **one file only**, named **main_test.go**. This is necessary in order to test and run snippets locally
  - folder should match the slug version of the title:
    
    if the title is **How to reverse a string**, 
    
    the folder should be **snippets/how-to-reverse-a-string**
    
    and the full path should be **snippets/how-to-reverse-a-string/main_test.go**
  - Do no edit anything inside docs folder. It is generated automatically on succesful merge  
  - Do no edit README.md. It is generated automatically on succesful merge. Edit https://github.com/psampaz/gophersnippets/edit/contributing/tools/templates/readme.md instead.
    
## License
<a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/">Creative Commons Attribution-ShareAlike 4.0 International License</a>.
