# lets-db

## Database Management System Components

1. Connection Manager:

    This db use grpc as the connection protocol to send query and receive data.

2. Query Parser:

    To parse the query, the current implementation is for simple key-value string-string storage.

3. Query Analyzer: Pending

4. Query Optimizer: Pending

5. IO:

    Functions to read/write data given the key.

6. Engine:

    Manage the core db engine, the current implementation is for in-memory hash table.

7. Result Formatter: Pending

## User Manual: 

### To run all tests, cd to the repo root directory then run:

```
go test ./... -v
```
