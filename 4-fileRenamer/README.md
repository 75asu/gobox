# File Renamer

A command-line tool for renaming files in a specified directory using custom renaming rules.

### Features

- **File Renaming:** Rename files in a specified directory using custom renaming rules.
- **Rule Definitions:** Define renaming rules using file extensions and new names.

### Usage

To use the file renamer app, follow these steps:

1. **Build the Application:**

   ```bash
   go build -o fileRenamer
   ```

2. **Specify Directory and Rules:**

   Run the `fileRenamer` executable with the following command-line flags:

   - `-dir`: Specify the directory path for file renaming. By default, it uses the current directory.

   - `-rules`: Define the renaming rules as a comma-separated list of file extension mappings. For example, you can specify rules like `.txt:_text,.jpg:_image`.

   ```shell
   ./fileRenamer -dir <directory-path> -rules "<rules>"
   ```

3. **Example:**

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

### Command-line Flags

- `-dir`: Specify the directory path for file renaming. (Default: Current directory)
- `-rules`: Define renaming rules as a comma-separated list of file extension mappings.

### Example Rules

- `.txt:_text`: Renames `.txt` files to include "_text" in the filename.
- `.jpg:_image`: Renames `.jpg` files to include "_image" in the filename.

Feel free to customize the renaming rules and directory path to suit your needs. This tool can be handy for batch file renaming tasks.

