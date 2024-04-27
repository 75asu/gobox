### word count tool

- [Here](https://codingchallenges.fyi/challenges/challenge-wc) is the challenge
- Here is how the tool in action looks like

```
╭─ ~/gobox/wc-tool on  main! 17:37:54
╰─▶▶ go build .

╭─ ~/gobox/wc-tool on  main! 17:38:00
╰─▶▶ ./wc-tool readthis.txt 
byteCount:      845
lineCount:       23
wordCount:      122
charCount:      845

╭─ ~/gobox/wc-tool on  main! 17:38:09
╰─▶▶ ./wc-tool -b readthis.txt
byteCount:      845

╭─ ~/gobox/wc-tool on  main! 17:38:23
╰─▶▶ ./wc-tool -l readthis.txt
lineCount:       23

╭─ ~/gobox/wc-tool on  main! 17:38:28
╰─▶▶ ./wc-tool -w readthis.txt 
wordCount:      122

╭─ ~/gobox/wc-tool on  main! 17:38:33
╰─▶▶ ./wc-tool -c readthis.txt
charCount:      845

