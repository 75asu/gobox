- collect port from env file
- add a router
- pass route and port to a server
- use cors in the router
- add helper function to return JSON payload
- add a HTTP handler
- hookup the handler with the http method and path in chi router
- verify the router by sending a request
- add helper function to return standard Error messages
- add a new handler for the helper function
- hook the respondWithError helper function with another handler
- setup database
  - create a new database

    ```shell
    PGPORT=5432
    PGHOST="localhost"
    PGUSER="postgres"
    PGDB="rssagg"
    createdb -h $PGHOST -p $PGPORT -U $PGUSER $PGDB
    ```

  - install sqlc and goose package
  - setup sqlc, sql directory, queries directory, schema directory, try sqlc generate to verify once
  - import database connection in main.go
  - import postgres driver package
  - adjust an handler for create user
  - hook up the handler in a route in main
  - define the handleCreateUser properly
- Create new model User and use this in handlerCreateUser
