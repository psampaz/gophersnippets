**[gophersnippets](https://gophersnippets.com)** is a collection of code snippets with tests and testable examples for the Go programming language. 

Each snippet:
 - is small and shows how to achieve something specific with Go 
 - contains at least one test or testable example
 - runs on [Go playground](https://play.golang.org/) since it now supports tests and testable examples  



## Conversions
 - [How to convert int to string](https://gophersnippets.com/how-to-convert-int-to-string)
## HTTP
 - [How to print a raw HTTP response](https://gophersnippets.com/how-to-print-a-raw-http-response)
 - [How to read an HTTP response status code](https://gophersnippets.com/how-to-read-an-http-response-status-code)
## Interfaces
 - [How to check if a type satisfies an interface at runtime](https://gophersnippets.com/how-to-check-if-a-type-satisfies-an-interface-at-runtime)
## Maps
 - [How to copy a map](https://gophersnippets.com/how-to-copy-a-map)
 - [How to create a set using a map](https://gophersnippets.com/how-to-create-a-set-using-a-map)
## Numbers
 - [How to calculate the sum of multiple numbers](https://gophersnippets.com/how-to-calculate-the-sum-of-multiple-numbers)
 - [What is the maximum value of numeric types](https://gophersnippets.com/what-is-the-maximum-value-of-numeric-types)
## Other
 - [How to format Go code programmatically](https://gophersnippets.com/how-to-format-go-code-programmatically)
 - [How to print the binary representation of an integer](https://gophersnippets.com/how-to-print-the-binary-representation-of-an-integer)
 - [How to print the same variable multiple times using printf](https://gophersnippets.com/how-to-print-the-same-variable-multiple-times-using-printf)
## Parser
 - [How to parse comments from Go Code](https://gophersnippets.com/how-to-parse-comments-from-go-code)
## Runtime
 - [How to print Go version](https://gophersnippets.com/how-to-print-go-version)
 - [How to print information about the operating system, architecture and pointer size](https://gophersnippets.com/how-to-print-information-about-the-operating-system-architecture-and-pointer-size)
## Slices
 - [How to check if a slice contains a specific element](https://gophersnippets.com/how-to-check-if-a-slice-contains-a-specific-element)
 - [How to delete an element from a slice](https://gophersnippets.com/how-to-delete-an-element-from-a-slice)
 - [How to filter a slice](https://gophersnippets.com/how-to-filter-a-slice)
 - [How to find the maximum element of a slice](https://gophersnippets.com/how-to-find-the-maximum-element-of-a-slice)
 - [How to find the minimum element of a slice](https://gophersnippets.com/how-to-find-the-minimum-element-of-a-slice)
 - [How to remove duplicate elements from a slice](https://gophersnippets.com/how-to-remove-duplicate-elements-from-a-slice)
 - [How to reverse a slice](https://gophersnippets.com/how-to-reverse-a-slice)
 - [How to shuffle a slice](https://gophersnippets.com/how-to-shuffle-a-slice)
## Strings
 - [How to calculate the hamming distance between two strings](https://gophersnippets.com/how-to-calculate-the-hamming-distance-between-two-strings)
 - [How to check if a string ends with a substring](https://gophersnippets.com/how-to-check-if-a-string-ends-with-a-substring)
 - [How to check if a string is lowercase](https://gophersnippets.com/how-to-check-if-a-string-is-lowercase)
 - [How to check if a string is uppercase](https://gophersnippets.com/how-to-check-if-a-string-is-uppercase)
 - [How to check if a string starts with substring](https://gophersnippets.com/how-to-check-if-a-string-starts-with-substring)
 - [How to check if string is valid JSON](https://gophersnippets.com/how-to-check-if-string-is-valid-json)
 - [How to concatenate strings](https://gophersnippets.com/how-to-concatenate-strings)
 - [How to count the number of words in a string](https://gophersnippets.com/how-to-count-the-number-of-words-in-a-string)
 - [How to get the md5 checksum of a string](https://gophersnippets.com/how-to-get-the-md5-checksum-of-a-string)
 - [How to get the sha1 checksum of a string](https://gophersnippets.com/how-to-get-the-sha1-checksum-of-a-string)
 - [How to get the sha256 checksum of a string](https://gophersnippets.com/how-to-get-the-sha256-checksum-of-a-string)
 - [How to pretty print JSON](https://gophersnippets.com/how-to-pretty-print-json)
 - [How to reverse a string](https://gophersnippets.com/how-to-reverse-a-string)
## Testing
 - [How to test a function that panics](https://gophersnippets.com/how-to-test-a-function-that-panics)
## Time
 - [How to measure the execution time of a function](https://gophersnippets.com/how-to-measure-the-execution-time-of-a-function)

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
