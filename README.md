## test-ITMX
### Setup and Run:

### Setup

#### install
    
 - [install Go](https://go.dev/doc/install)

- [install sqlite](https://www.sqlite.org/download.html)
#### config at path "confg.confg.yaml"
1. config number port 
    ```bash
        port: { { app-port } }
    ```
2. config sqlite-dbname 
    ```bash
        dbname: { { sqlite-dbname } }
    ```
1. config sqlite-dbpath 
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
    sqlite3 Db-name.db < filename.sql


### test Go
1. run unit test all files and display coverage.
    ```bash
    go test ./... -cover