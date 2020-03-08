package main

import (
	"fmt"
	"net/http"
	TokenDaos "../token_server/lib"
	"strconv"
	"strings"
)

func checkTokenExists(token string) bool {
	return TokenDaos.TokenExists(token)
}

func checkPathCount(path string) int {
	return TokenDaos.PathCount(path)
}

type ReturnCount struct {
	count int
}

func checkPath(w http.ResponseWriter, r *http.Request) {
	if checkTokenExists(r.Header.Get("Authorization")) {
		path := strings.Replace(r.URL.String(), "/", "", -1)
		returnCount := ReturnCount{checkPathCount(path)}
		//converting this to look like a JSON object
		returnJSON := "{\"count\": \"" + strconv.FormatInt(int64(returnCount.count), 10) + "\"}"
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, returnJSON)

	} else {
		w.WriteHeader(401)
		fmt.Fprintf(w, "Unauthorized")
	}
}

type ReturnToken struct {
	token string
}

func insertNewToken(w http.ResponseWriter, r *http.Request) {
	returnToken := ReturnToken{TokenDaos.CreateToken()}
	//converting this to look like a JSON object
	returnJSON := "{\"token\": \"" + returnToken.token + "\"}"
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, returnJSON)
}

func main() {
	fmt.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/auth", insertNewToken)
	http.HandleFunc("/", checkPath)
	http.ListenAndServe(":8080", nil)
}