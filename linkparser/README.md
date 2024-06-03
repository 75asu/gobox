### Link Parser

- [Here](https://github.com/gophercises/link?tab=readme-ov-file#exercise-4-html-link-parser) is the challenge

- This is how the output looks
```bash
gitpod /workspace/gobox/linkparser (main) $ go run examples/ex1/main.go 
[{Href:/other-page Text:A link to another page some span} {Href:/page-two Text:A link to a second page}]

gitpod /workspace/gobox/linkparser (main) $ go run examples/ex2/main.go 
[{Href:https://www.twitter.com/joncalhoun Text:Check me out on twitter} {Href:https://github.com/gophercises Text:Gophercises is on Github!}]

gitpod /workspace/gobox/linkparser (main) $ go run examples/ex3/main.go 
[{Href:# Text:Login} {Href:/lost Text:Lost? Need help?} {Href:https://twitter.com/marcusolsson Text:@marcusolsson}]

gitpod /workspace/gobox/linkparser (main) $ go run examples/ex4/main.go 
[{Href:/dog-cat Text:dog cat}]
```
