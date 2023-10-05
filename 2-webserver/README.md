# Static Site Web Server

A simple web server for hosting a static website on port 8080.

### Features

- **Web Server:** Host a static website on port 8080.

### Usage

To use the static site web server, follow these steps:

1. **Run the Application:**

   Run the `main.go` file to start the web server:

   ```shell
   go run main.go
   ```

   The server will start listening on port 8080 by default.

2. **Access the Website:**

   Open a web browser and navigate to `http://localhost:8080` to access the hosted static website.

   ![Static Site](image.png)

### Default Port

The web server listens on port 8080 by default. You can customize the port by modifying the code in `main.go` if needed.

### Example

```shell
gitpod /workspace/gobox/2-webserver (main) $ go run main.go 
Server is listening on :8080
```

The server is now running, and you can access the site at `http://localhost:8080`.

Feel free to replace the `image.png` file with your own static website content to host your site using this simple web server.

