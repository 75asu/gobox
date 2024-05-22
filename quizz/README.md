### word count tool

- [Here](https://github.com/gophercises/quiz/blob/master/README.md) is the challenge
- Here is how the solution in action looks like

```
gitpod /workspace/gobox/quizz (main) $ ./quizz -limit=5
Problem #1: 5+5 = 10
Correct!
Problem #2: 1+1 = 2
Correct!
Problem #3: 8+3 = 11
Correct!
Problem #4: 1+2 = 
You scored 3 out of 12.

gitpod /workspace/gobox/quizz (main) $ ./quizz -limit=5
Problem #1: 5+5 = 234
Wrong!
Problem #2: 1+1 = 5423
Wrong!
Problem #3: 8+3 = 11
Correct!
Problem #4: 1+2 = 
You scored 1 out of 12.

gitpod /workspace/gobox/quizz (main) $ ./quizz -csv=problems.csv
Problem #1: 5+5 = 10
Correct!
Problem #2: 1+1 = 2
Correct!
Problem #3: 8+3 = 11
Correct!
Problem #4: 1+2 = ^C
```