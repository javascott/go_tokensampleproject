Token Sample Services in Go

Setup:
1) create a mysqlDB on your local system
2) Run sql/1.sql 
    1) It will create all DB tables
    2) 1 sample data row
    3) create user/grant permissions 
3) Make sure Golang is setup on your PATH, including the following libraries
    1) go get -u github.com/go-sql-driver/mysql
    2) go get google.golang.org/grpc
    3) go get github.com/google/uuid
    
4) To Run the RPC system
    1) Start Server
        1) navigate a terminal to {project directory}
        2) go run token_server/server.go
    2) Start Client
        1) navigate a terminal to {project directory}
        2) go run token_client/client.go {path_to_test}
        
5) To Run the REST system
    1) Start Server
        1) navigate a terminal to {project directory}
        2) go run rest_server/rest.go 
    2) Use Postman or similar client
        1) GET localhost:8080/auth 
            1) it will return you a token
        2) GET localhost:8080/{path_to_test}
            1) you will need to set the Header parameter of Authorization: {token from step 1}
           

