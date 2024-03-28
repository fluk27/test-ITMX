## test-ITMX
### Setup and Run:

### Setup

#### install
    
 - [install Go](https://go.dev/doc/install)

- [install sqlite](https://www.sqlite.org/download.html)
#### config at path "config/confg.yaml"
1. config number port 
    ```bash
        port: { { app-port } }
    ```
2. config sqlite  database name 
    ```bash
        dbname: { { sqlite-dbname } }
    ```
1. config sqlite database path 
    ```bash
        dbpath: { { sqlite-dbpath } }
    ```
### Run Go
1. run install all package.

    ```bash
    go mod tidy
    ```

    2. start go server.

    ```bash
    go run cmd/main.go
    ```
### Run script sqlite
- if you want sql execute to database.Please follow these steps.

    1.create file just last name ".sql".

    2.run sqlite command.

    ```bash
    sqlite3 Db-name.db < filename.sql
    ```


### test Go
1. run unit test all files and display coverage.
    ```bash
    go test ./... -cover
