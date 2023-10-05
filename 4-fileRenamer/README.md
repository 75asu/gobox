### A file renamer app that allows you to specify the directory, whether to traverse subdirectories recursively, and the renaming rules using command-line flags. 

Here's how you can use it:

- `-dir`: Specify the directory path for file renaming. By default, it uses the current directory.

- `-rules`: Define the renaming rules as a comma-separated list of file extension mappings. 

  - For example, you can specify rules like .txt:_text,.jpg:_image.


**Demo**

```shell
gitpod /workspace/gobox/4-fileRenamer (main) $ go build -o fileRenamer
gitpod /workspace/gobox/4-fileRenamer (main) $ ls
dummy1_text.txt  dummy2_text.txt  fileRenamer  main.go  pexels_image.jpg
gitpod /workspace/gobox/4-fileRenamer (main) $ ./fileRenamer -dir ./ -rules ".txt:_renamed,.jpg:_renamed"
Renamed dummy1_text.txt to dummy1_text_renamed.txt
Renamed dummy2_text.txt to dummy2_text_renamed.txt
Renamed pexels_image.jpg to pexels_image_renamed.jpg
gitpod /workspace/gobox/4-fileRenamer (main) $ ls
dummy1_text_renamed.txt  dummy2_text_renamed.txt  fileRenamer  main.go  pexels_image_renamed.jpg
```