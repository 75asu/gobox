# Todo App

A simple command-line todo app for managing your tasks.

### Features

- **Task Management:** Add, remove, and list tasks in a todo list.

### Usage

To use the todo app, follow these steps:

1. **Navigate to the Todo App Directory:**

   ```shell
   cd 1-todo/
   ```

2. **Run the Application:**

   Run the `main.go` file to start the todo app:

   ```shell
   go run main.go
   ```

3. **Interact with the Todo List:**

   The app will display a menu with options to add, remove, and quit. You can perform the following actions:

   - **Add a Task:** Enter "add" and provide a task name.

   - **Remove a Task:** Enter "remove" and provide the task ID to remove.

   - **Quit:** Enter "quit" to exit the app.

   Here's an example:

   ```shell
   gitpod /workspace/gobox/1-todo (main) $ go run main.go 
   TODO List:
   Enter an option (add/remove/quit): 
   add
   Enter task name: 
   learn golang
   TODO List:
   1: learn golang
   Enter an option (add/remove/quit): 
   add
   Enter task name: 
   code every day until the end of the year
   TODO List:
   1: learn golang
   2: code every day until the end of the year
   Enter an option (add/remove/quit): 
   remove
   Enter task ID to remove: 
   2
   TODO List:
   1: learn golang
   Enter an option (add/remove/quit): 
   quit
   ```

### Example

You can use this todo app to manage your tasks and keep track of your to-do list efficiently.

Feel free to add, remove, and list tasks as needed to help you stay organized.

