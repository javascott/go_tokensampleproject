Token Sample Service in Golang

Setup:
1) create a mysqlDB on your local system
2) Run sql/1.sql 
    1) It will create all DB tables
    2) 1 sample date row
    3) create user/grant permissions 
3) Make sure Golang is setup on your PATH, including the following libraries
    1) go get -u github.com/go-sql-driver/mysql
    2) go get google.golang.org/grpc
    3) go get github.com/google/uuid
    
4) To Run the RPC tests
    1) Start Server
        1) navigate a terminal to {project directory}
        2) go run token_server/server.go
    1) Start Server
        1) navigate a terminal to {project directory}
        2) go run token_client/client.go {path_to_test}
        
        
                

